package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
	"log"
	"os"
	"time"
)

type Model struct {
	ID        uint       `json:"id,omitempty" gorm:"primary_key"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" sql:"index"`
}

var db *gorm.DB

func Init() {
	host := viper.GetString("db.host")
	port := viper.GetString("db.port")
	user := viper.GetString("db.user")
	password := viper.GetString("db.password")
	database := viper.GetString("db.database")
	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", user, password, host, port, database)
	//fmt.Println("Connect to", url)
	var err error
	db, err = gorm.Open("mysql", url)
	if err != nil {
		panic(err)
	}
	//defer db.Close()
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	db.LogMode(true)
	//db.SetLogger(gorm.Logger{revel.TRACE})
	db.SetLogger(log.New(os.Stdout, "\n", 0))
}

func GetDB() *gorm.DB {
	return db
}
