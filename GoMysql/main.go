package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql" // 匿名导入
)

type Data struct {
	Id         int64
	Key        string
	Value      string
	CreateTime time.Time
	UpdateTime time.Time
}

func (data Data) output() string {
	return fmt.Sprintf("id:%d,key:%s,value:%s,create_time:%s,update_time:%s", data.Id, data.Key, data.Value, data.CreateTime.Format("2006-01-02 15:04:05"), data.UpdateTime.Format("2006-01-02 15:04:05"))
}

func insert(key string, value string, db *sql.DB) *Data {

	// db.prepare用于预编译sql语句，返回一个sql.stmt对象，该对象可以多次执行sql语句，而无需每次重新编译
	//Prepare 使用参数化查询，将参数与 SQL 语句分开，可以有效防止 SQL 注入。
	//Prepare 会返回 *sql.Stmt 对象，你需要在用完后关闭它（使用 defer stmt.Close()），以释放数据库连接和资源。

	stmtOut, err := db.Prepare("insert into `data` (`key`,`value`) values (?,?)")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer stmtOut.Close()

	result, err := stmtOut.Exec(key, value)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	//result.LastInsertId()用于获取上一次插入操作中生成的自增主键 ID
	id, err := result.LastInsertId()
	return query(id, db)
}

func del(id int64, db *sql.DB) bool {
	stmtOut, err := db.Prepare("DELETE FROM `data` WHERE id = ?")
	if err != nil {
		fmt.Println("Prepare statement error:", err)
		return false
	}
	defer stmtOut.Close()

	result, err := stmtOut.Exec(id)
	if err != nil {
		fmt.Println("Exec error:", err)
		return false
	}

	// 检查受影响的行数
	rownum, err := result.RowsAffected()
	if err != nil {
		fmt.Println("RowsAffected error:", err)
		return false
	}

	// 返回是否有行被删除
	return rownum > 0
}

func query(id int64, db *sql.DB) *Data {
	stmtOut, err := db.Prepare("select * from `data` where id=?")
	if err != nil {
		panic(err.Error())

	}
	defer stmtOut.Close()

	rows := stmtOut.QueryRow(id)
	data := new(Data)
	rows.Scan(&data.Id, &data.Key, &data.Value, &data.CreateTime, &data.UpdateTime)
	fmt.Println(data.output())
	return data
}

func update(id int64, key string, value string, db *sql.DB) *Data {
	stmtOut, err := db.Prepare("update `data` set `key`=?,`value`=? where id=?")
	if err != nil {
		panic(err.Error())
	}
	defer stmtOut.Close()
	stmtOut.Exec(key, value, id)
	return query(id, db)

}

func main() {
	// 数据库配置
	// 格式：用户名:密码@tcp(ip:端口)/数据库名
	dsn := "root:root@tcp(localhost:3306)/golang"
	db, err := sql.Open("mysql", dsn)
	db.Ping()
	if err != nil {
		fmt.Println(err)
		return
	}
	//设置数据库连接池的最大连接数
	db.SetMaxIdleConns(10)
	query(1, db)
	insert("hello", "world", db)
	update(1, "hello", "golang", db)
	del(1, db)
}
