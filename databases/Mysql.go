package databases

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id         int64
	Username   string
	Password   string
	Status     bool
	CreateTime time.Time
}

func InitDatabase() *sql.DB {
	// 数据库连接

	dsn := "root:root@tcp(localhost:3306)/webgin"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("数据库连接失败", err)
	}
	return db
}

func CreateTable(db *sql.DB) error {
	sql := `
	CREATE TABLE IF NOT EXISTS user(
		id int(11) NOT NULL AUTO_INCREMENT,
		username varchar(100) NOT NULL,
		password varchar(100) NOT NULL,
		status tinyint(1) NOT NULL,
		create_time datetime NOT NULL,
		PRIMARY KEY (id)
	) ENGINE=InnoDB DEFAULT CHARSET=utf8;`
	_, Err := ModifyDB(sql, db)
	return Err
}

func ModifyDB(sql string, db *sql.DB, args ...interface{}) (int64, error) {
	result, err := db.Exec(sql, args...)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	count, err := result.RowsAffected()
	if err != nil {
		log.Println(err)
		return 0, err
	}
	return count, nil
}

func Insert(user User) (int64, error) {
	return ModifyDB("insert into user(username, password, status, create_time) values(?,?,?,?)", InitDatabase(), user.Username, user.Password, user.Status, user.CreateTime)
}

func Update(user User, username string) (int64, error) {
	return ModifyDB("update user set password=? where username=?", InitDatabase(), user.Password, username)
}

func Query(id int64) {
	db := InitDatabase()
	stmtOut, err := db.Prepare("select * from user where id=?")
	if err != nil {
		panic(err.Error())
	}
	defer stmtOut.Close()
	rows := stmtOut.QueryRow(id)
	user := new(User)
	rows.Scan(&user.Id, &user.Username, &user.Password, &user.Status, &user.CreateTime)
	fmt.Sprintf("user(id:%d, username:%s, password:%s, status:%t, time:%s)", user.Id, user.Username, user.Password, user.Status, user.CreateTime)
}

func QueryUsername(username string) (*User, error) {
	db := InitDatabase()
	stmtOut, err := db.Prepare("select * from user where username=?")
	if err != nil {
		panic(err.Error())
	}
	defer stmtOut.Close()
	rows := stmtOut.QueryRow(username)
	user := new(User)
	err = rows.Scan(&user.Id, &user.Username, &user.Password, &user.Status, &user.CreateTime)
	return user, nil
}

func QueryAll(username string) (*User, error) {
	db := InitDatabase()
	stmtOut, err := db.Prepare("select * from user where username=?")
	if err != nil {
		panic(err.Error())
	}
	defer stmtOut.Close()
	rows := stmtOut.QueryRow(username)
	if err != nil {
		return nil, err
	}
	user := new(User)
	err = rows.Scan(&user.Id, &user.Username, &user.Password, &user.Status, &user.CreateTime)
	return user, nil
}
