package main

import (
	"database/sql"
	"fmt"

	_ "github.com/godror/godror" // v0.35.1
)

func main() {
	// get oracle connection
	db, err := sql.Open("godror", `user="tmcp" password="padboratmcp" connectString="10.57.64.131:1521/PADB"`)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()
	// scan result
	var result string
	err = db.QueryRow("select 'hello world' from dual").Scan(&result)
	if err != nil {
		fmt.Println(err)
		return
	}
	// print result
	fmt.Println(result)
}
