package connecting

import (
	"database/sql"
	"fmt"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "1111"
	dbname   = "MyData"
)

var Db *sql.DB

func InitDB() *sql.DB {
	var err error
	conn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	Db, err = sql.Open("postgres", conn)
	check(err)
	return Db
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
