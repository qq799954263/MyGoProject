package main

import (
	"fmt"
	"database/sql"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

/*
		dip  = "127.0.0.1:3306"
		username = "root"
		password = ""
		name     = "Student"
*/

func main(){

	//Query()
	//Query2()
	//insert("黎明","2班","男")
	//update(5)

	//delete()
	Query2()
}
func Query() {
	db, err := sql.Open("mysql", "root:@/Student")
	checkError(err)

	//查询数据 乱序打印
	rows, err := db.Query("SELECT * FROM StudentList")
	checkError(err)

	for rows.Next(){
		columns,_ := rows.Columns()

		scanArgs := make([]interface{},len(columns))
		values := make([]interface{},len(columns))

		for i := range values{
			scanArgs[i] = &values[i]
		}

		err = rows.Scan(scanArgs...)
		record := make(map[string]string)
		for i,col := range values{
			if col != nil {
				record[columns[i]] = string(col.([]byte))
			}
		}
		fmt.Println(record)
	}
	rows.Close()
}

func Query2() {
	fmt.Println("Query2")
	db, err := sql.Open("mysql", "root:@/Student")
	checkError(err)

	// 查询数据 有序打印
	rows, err := db.Query("SELECT id,username,class,sex FROM StudentList")
	checkError(err)

	for rows.Next() {
		var id int
		var username string
		var class string
		var sex string
		//注意这里的Scan括号中的参数顺序，和 SELECT 的字段顺序要保持一致。
		if err := rows.Scan(&id, &username, &class, &sex); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s: id  %d 班级 %s 性别 %s \n", username, id, class, sex)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	rows.Close()
}
// 插入数据
func insert(name,class,sex string) {
	db,err := sql.Open("mysql","root:@/Student")
	checkError(err)

	stmt, err := db.Prepare("insert StudentList set username=?,class=?,sex=?")
	checkError(err)

	res,err :=stmt.Exec(name,class,sex)
	checkError(err)

	id,err := res.LastInsertId()
	checkError(err)

	fmt.Println(id)
}

//修改数据
func update(id int){
	db,err := sql.Open("mysql","root:@/Student")
	checkError(err)

	stmt,err := db.Prepare("UPDATE StudentList set username=? where id=?")
	checkError(err)

	res,err := stmt.Exec("赵四",id)
	checkError(err)

	affect,err := res.RowsAffected()
	checkError(err)

	fmt.Println(affect)
	stmt.Close()
}

func delete(){
	db,err := sql.Open("mysql","root:@/Student")
	checkError(err)

	stmt,err := db.Prepare("delete from StudentList where id=?")
	checkError(err)

	res,err := stmt.Exec(4)
	checkError(err)

	affect,err := res.RowsAffected()
	checkError(err)

	fmt.Println(affect)
	stmt.Close()

}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
