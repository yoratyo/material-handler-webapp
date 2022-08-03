# material-handler-webapp

## Install assetFS generator:
```
go install github.com/go-bindata/go-bindata/go-bindata@latest
go install github.com/elazarl/go-bindata-assetfs/go-bindata-assetfs@latest
```

## Generate static asset:
```
go-bindata-assetfs -o assets/css/asset.go assets/css/...
go-bindata-assetfs -o assets/fonts/asset.go assets/fonts/...
go-bindata-assetfs -o assets/img/asset.go assets/img/...
go-bindata-assetfs -o assets/js/asset.go assets/js/...
go-bindata-assetfs -o assets/scss/asset.go assets/scss/...
go-bindata-assetfs -o assets/vendor/asset.go assets/vendor/...
```
Or run via Makefile
```
make generate-assets
```
