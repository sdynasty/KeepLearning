package dao

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/xormplus/xorm"
)

var (
	engine *xorm.Engine
	db *sql.DB
)

//func InitXEngine() {
//	db, err := xorm.NewEngine("mysql", "root:123456@/tcp:(localhost:3306)/clever?charset=utf8")
//	if err != nil {
//		panic(err)
//	}
//	engine = db
//}

func InitMysql() {
	d, err := sql.Open("mysql", "root:123456@/tcp:(localhost:3306)/clever")
	if err != nil {
		panic(err)
	}
	db = d
}

//func getDBEngine() *xorm.Engine {
//	return engine
//}

func getDB() *sql.DB {
	return db
}