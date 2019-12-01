package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

//Make a connect to sql server
//TODO: need a independ datastruct to represent a connect
//TODO: sparate test code
func conn() {
	LoadConfig("conf.json")
	db, err := sql.Open("mysql", config.sql.user+":"+config.sql.password+"@/"+"express")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println(err)
		return
	}

	stmIns, err := db.Prepare("select username from user")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer stmIns.Close()

	f := "fff"
	fmt.Println(f)
	err = stmIns.QueryRow().Scan(&f)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(f)

	stmbyid, err := db.Prepare("select username from user where uid = ?")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer stmbyid.Close()

	ff := "fff"
	fmt.Println(ff)
	err = stmbyid.QueryRow(1).Scan(&ff)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(ff)

	err = stmbyid.QueryRow(2).Scan(&ff)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(ff)
}
