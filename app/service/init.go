package service

import (
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"go-template/app/model"
	"log"
	"os"
	"strconv"
	_ "strconv"
)

var DbEngine *xorm.Engine

func init() {
	driverName := os.Getenv("DB_DRIVER")
	DsName := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_CHARSET"),
	)
	err := errors.New("")
	DbEngine, err = xorm.NewEngine(driverName, DsName)
	if err != nil && err.Error() != "" {
		log.Fatal(err.Error())
	}
	DbEngine.ShowSQL(true)
	openConns, err := strconv.ParseInt(os.Getenv("DB_MAX_OPEN_CONNS"), 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	DbEngine.SetMaxOpenConns(int(openConns))
	err = DbEngine.Sync2(new(model.Novel))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("init data base ok")
}
