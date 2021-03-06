package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5431
	user     = "asheesh.sharma-mba"
	password = "Arush1234$"
	dbname   = "mydb"
)

func main() {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	defer db.Close()

	// insert
	// hardcoded
	insertStmt := `insert into "students"("name", "role") values('John', 1)`
	_, e := db.Exec(insertStmt)
	CheckError(e)

	// dynamic
	insertDynStmt := `insert into "students"("name", "role") values($1, $2)`
	_, e = db.Exec(insertDynStmt, "Jane", 2)
	CheckError(e)
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
