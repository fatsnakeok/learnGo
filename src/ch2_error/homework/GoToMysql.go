package main

import (
	_ "github.com/go-sql-driver/mysql"
	//"github.com/jmoiron/sqlx"
	"database/sql"
	"github.com/pkg/errors"
)



// 用户结构体
type User struct {
	UserId int
	Username string
	Sex string
	Email string
}

// 数据库指针
var db *sql.DB

// 初始化数据库连接，init()方法系统会在main方法之前执行
func initConn() (db *sql.DB, err error){
	db, err = sql.Open("mysql", "root:rootroot@tcp(127.0.0.1:3306)/mytest")
	if err != nil {
		err = errors.Wrap(err, "open mysql failed,")
		return
	}
	return
}







func queryByParam (db *sql.DB,sqlStr string, id int ) (u User, err error) {
	err = db.QueryRow(sqlStr, id).Scan(&u.UserId,&u.Username)
	switch {
	case err == sql.ErrNoRows:
		err = errors.Wrap(err, "ErrNoRows:" + sqlStr)
	case err != nil:
		err = errors.Wrap(err, "scan error:" + sqlStr)
	}

	return
}

