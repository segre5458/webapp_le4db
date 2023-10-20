package main

import (
	"fmt"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type TestUser struct {
	UserID   int
	Password string
}

func main() {
	var Db *sql.DB

	Db, err := sql.Open("postgres", "host=postgres user=postgres password=postgres dbname=postgres sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer Db.Close()

	sql := "SELECT * FROM TEST_USER WHERE USER_ID = $1;"

	pstatement, err := Db.Prepare(sql)
	if err != nil {
		log.Fatal(err)
	}
	defer pstatement.Close()

	qeury := 1
	var testUser TestUser

	err = pstatement.QueryRow(qeury).Scan(&testUser.UserID, &testUser.Password)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(testUser.UserID, testUser.Password)
}
