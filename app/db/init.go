package db

import (
	"errors"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"log"
	"strconv"
	_ "strconv"
)

var DbEngine *xorm.Engine

func init() {
	var _ = mysql.ErrInvalidConn
	DsName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf-8&parseTime=True&loc=Local",
		"root",
		"yPHJEE!uxP9v",
		"mysql",
		"3306",
		"template",
	)
	log.Print(DsName)
	err := errors.New("")
	DbEngine, err = xorm.NewEngine("mysql", DsName)
	if err != nil && err.Error() != "" {
		log.Fatalf("Failed to create database engine: %s", err.Error())
	}
	DbEngine.ShowSQL(true)
	openConns, err := strconv.ParseInt("2", 10, 64)
	if err != nil {
		log.Fatalf("Failed to parse DB_MAX_OPEN_CONNS: %s", err.Error())
	}
	DbEngine.SetMaxOpenConns(int(openConns))
	fmt.Println("Initialized database successfully")
}
