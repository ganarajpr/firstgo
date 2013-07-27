
package main
import (
	"database/sql"
	"fmt"
	//"encoding/json"
	//"strings"
)

import _ "github.com/go-sql-driver/mysql"


type FuncList struct{
	Name string
	Definition string
	Inputs []string
	Outputs []string
}

func main() {
	db, err := sql.Open("mysql", "root@/fad")
	if err != nil {
		panic(err.Error())  // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()
	rows, err := db.Query("SELECT * FROM FuncList")
	if err != nil {
		panic(err.Error())
	}

	for rows.Next() {
		var name string
		var id int
		var definition string
		rows.Scan(&id,&name,&definition)
		fmt.Println(id,name,definition)
	}
}
