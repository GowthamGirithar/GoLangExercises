package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql-master" //import for the mysql driver //package is only for one purpose we use _

)

/**
CREATE TABLE `userinfo` (
       `uid` INT(10) NOT NULL AUTO_INCREMENT,
       `username` VARCHAR(64) NULL DEFAULT NULL,
       `departname` VARCHAR(64) NULL DEFAULT NULL,
       `created` DATE NULL DEFAULT NULL,
       PRIMARY KEY (`uid`)
   );
 */

//database connection

type DBConfig struct {
	HostName string
	Port     int
	UserName string
	Password string
	Schema   string
}


var DbConfigSetting = DBConfig{}

func init() {
	println("init method ")
	//initiallize the db connection parameters
	DbConfigSetting.HostName = "127.0.0.1"
	DbConfigSetting.Password = "root@123"
	DbConfigSetting.UserName = "admin"
	DbConfigSetting.Port = 3306
	DbConfigSetting.Schema = "test"
}

func main() {
	//open the connection
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", DbConfigSetting.UserName, DbConfigSetting.Password,
		DbConfigSetting.HostName, DbConfigSetting.Port, DbConfigSetting.Schema))
	defer db.Close()

	// insert the record
	stmt, err := db.Prepare("INSERT INTO userinfo (username,departname,created) VALUES (?,?,?)")
	checkErr(err)
	defer stmt.Close()
	res, err := stmt.Exec("astaxie", "研发部门", "2012-12-09")
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	// select the record and print
	selDB, err := db.Query("SELECT * FROM userinfo WHERE uid=?", id)
	checkErr(err)
	for selDB.Next() {
		var uid int
		var username, department string
		var created string   // date is also string
		err = selDB.Scan(&uid, &username, &department , &created) // we should have all the fields in * case
		if err != nil {
			panic(err.Error())
		}
		println("the data inserted are " , uid , username , department)
	}


}

func checkErr(e error) {
	if e != nil{
		println(e.Error())
	}

}
