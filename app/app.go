package app

import (
	"html/template"
	"os"
	"strings"
	"time"

	assetfs "github.com/elazarl/go-bindata-assetfs"
	"github.com/gin-gonic/gin"
	assetCSS "github.com/yoratyo/material-handler-webapp/assets/css"
	assetFonts "github.com/yoratyo/material-handler-webapp/assets/fonts"
	assetIMG "github.com/yoratyo/material-handler-webapp/assets/img"
	assetJS "github.com/yoratyo/material-handler-webapp/assets/js"
	assetSCSS "github.com/yoratyo/material-handler-webapp/assets/scss"
	assetVendor "github.com/yoratyo/material-handler-webapp/assets/vendor"
	"github.com/yoratyo/material-handler-webapp/handlers"
	assetHTML "github.com/yoratyo/material-handler-webapp/templates"

	"github.com/jasonlvhit/gocron"
)

const (
	PRODUCTIONENV = "production"
)

type App struct {
	Engine  *gin.Engine
	Handler handlers.Handler
}

func (a *App) Start() error {
	t, err := loadTemplate()
	if err != nil {
		panic(err)
	}
	a.Engine.SetHTMLTemplate(t)

	engine := NewRouter(a)

	fsCSS := assetfs.AssetFS{Asset: assetCSS.Asset, AssetDir: assetCSS.AssetDir, AssetInfo: assetCSS.AssetInfo, Prefix: "assets/css", Fallback: "index.html"}
	fsFonts := assetfs.AssetFS{Asset: assetFonts.Asset, AssetDir: assetFonts.AssetDir, AssetInfo: assetFonts.AssetInfo, Prefix: "assets/fonts", Fallback: "index.html"}
	fsJS := assetfs.AssetFS{Asset: assetJS.Asset, AssetDir: assetJS.AssetDir, AssetInfo: assetJS.AssetInfo, Prefix: "assets/js", Fallback: "index.html"}
	fsIMG := assetfs.AssetFS{Asset: assetIMG.Asset, AssetDir: assetIMG.AssetDir, AssetInfo: assetIMG.AssetInfo, Prefix: "assets/img", Fallback: "index.html"}
	fsSCSS := assetfs.AssetFS{Asset: assetSCSS.Asset, AssetDir: assetSCSS.AssetDir, AssetInfo: assetSCSS.AssetInfo, Prefix: "assets/scss", Fallback: "index.html"}
	fsVendor := assetfs.AssetFS{Asset: assetVendor.Asset, AssetDir: assetVendor.AssetDir, AssetInfo: assetVendor.AssetInfo, Prefix: "assets/vendor", Fallback: "index.html"}

	engine.StaticFS("/page/css", &fsCSS)
	engine.StaticFS("/page/fonts", &fsFonts)
	engine.StaticFS("/page/img", &fsIMG)
	engine.StaticFS("/page/js", &fsJS)
	engine.StaticFS("/page/scss", &fsSCSS)
	engine.StaticFS("/page/vendor", &fsVendor)

	go func() {
		gocron.Every(5).Seconds().Do(func() {
			a.Handler.Api.SchedulerGatewayCheckNFC(time.Now())
		})
		<-gocron.Start()
	}()

	return engine.Run(":" + os.Getenv("APP_PORT"))
}

func loadTemplate() (*template.Template, error) {
	t := template.New("")
	listFiles, err := assetHTML.AssetDir("templates")
	if err != nil {
		return nil, err
	}

	for _, name := range listFiles {
		if !strings.HasSuffix(name, ".html") {
			continue
		}
		html, _ := assetHTML.Asset("templates/" + name)
		t, err = t.New(name).Parse(string(html))
		if err != nil {
			return nil, err
		}
	}
	return t, nil
}

func NewApp(handler handlers.Handler) *App {
	if os.Getenv("ENV") == PRODUCTIONENV {
		gin.SetMode(gin.ReleaseMode)
	}

	return &App{
		Engine:  gin.Default(),
		Handler: handler,
	}
}
