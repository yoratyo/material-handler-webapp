package shared

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDB() (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"))
	log.Println("Trying to connect DB..")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		//Logger: logger.Default.LogMode(logger.Info),
	})

	log.Println("Connected to DB.")
	return db, err
}

func GetORMTransaction(c *gin.Context, resource Resource) (*gorm.DB, error) {
	var ok bool

	orm := resource.DB

	trxInt, exist := c.Get(MySQLTransaction)
	if exist {
		if orm, ok = trxInt.(*gorm.DB); !ok {
			return nil, errors.New("invalid transaction")
		}
	}

	return orm, nil
}
