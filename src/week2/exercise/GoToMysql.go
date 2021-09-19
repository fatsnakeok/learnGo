package exercise

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

// 用户结构体
type User struct {
	UserId   int    `db:"user_id"`
	Username string `db:"username"`
	Sex      string `db:"sex"`
	Email    string `db:"email"`
}

// 数据库指针
var db *sqlx.DB

// 初始化数据库连接，init()方法系统会在main方法之前执行
func init() {
	database, err := sqlx.Open("mysql", "root:rootroot@tcp(127.0.0.1:3306)/mytest")
	if err != nil {
		fmt.Println("open mysql failed,", err)
	}
	db = database
}

func main() {

	//u := User{1,"user01","man","user01@163.com"}
	//err := save (u )
	//fmt.Println(err)

	userList, err := findById(15)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(userList)

}

func save(u User) error {

	sql := "insert into user(id,username,sex,email)values(?,?,?)"

	//执行sql语句
	r, err := db.Exec(sql, u.UserId, u.Username, u.Sex, u.Email)
	fmt.Println(r.LastInsertId())
	return errors.Wrap(err, "exec "+sql+" failed")
}

func findById(id int) ([]User, error) {
	var user []User
	sql := "select user_id, username,sex,email from user where user_id=? "
	err := db.Select(&user, sql, id)
	if err != nil {
		return user, errors.Wrap(err, "exec "+sql+" failed")
	}
	return user, nil
}

func updateUserName(userName string, userId int) (int64, error) {
	//执行SQL语句
	sql := "update user set username =? where user_id = ?"
	res, err := db.Exec(sql, userName, userId)

	if err != nil {
		return 0, errors.Wrap(err, "exec "+sql+" failed")
	}

	//查询影响的行数，判断修改插入成功
	row, err := res.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "exec "+sql+" rows failed")
	}

	return row, nil
}

func deleteById(userId int) (int64, error) {
	sql := "delete from user where user_id=?"

	res, err := db.Exec(sql, userId)
	if err != nil {
		return 0, errors.Wrap(err, "exec "+sql+" failed")
	}

	row, err := res.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "exec "+sql+" rows failed")
	}
	return row, nil
}

func testTransaction() {
	//开启事务
	conn, err := db.Begin()
	if err != nil {
		fmt.Println("begin failed :", err)
		return
	}

	//执行插入语句
	r, err := conn.Exec("insert into user(username, sex, email)values(?, ?, ?)", "user01", "man", "usre01@163.com")

	if err != nil {
		fmt.Println("exec failed, ", err)
		conn.Rollback() //出现异常，进行回滚操作
		return
	}

	id, err := r.LastInsertId()
	if err != nil {
		fmt.Println("exec failed, ", err)
		conn.Rollback()
		return
	}
	fmt.Println("insert succ:", id)

	r, err = conn.Exec("insert into user(username, sex, email)values(?, ?, ?)", "user02", "man", "user02@163.com")
	if err != nil {
		fmt.Println("exec failed, ", err)
		conn.Rollback()
		return
	}
	id, err = r.LastInsertId()
	if err != nil {
		fmt.Println("exec failed, ", err)
		conn.Rollback()
		return
	}
	fmt.Println("insert succ:", id)

	conn.Commit()
}
