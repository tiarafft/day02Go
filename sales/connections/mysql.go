package connections

import (
	"database/sql"
	"fmt"
	// export mysql
	_ "github.com/go-sql-driver/mysql"
)

// GetConnection : fucntion ot create connection
func GetConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:lenovo10@tcp(localhost:3306)/gosales?parseTime=true")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("db is connected")
	}
	// defer db.Close()
	// check for connection
	err = db.Ping()
	fmt.Println(err)
	if err != nil {
		fmt.Println("db is not connected")
		fmt.Println(err.Error())
	}
	return db
}
