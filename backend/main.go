package main

import (
	// "fmt"
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"	

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

	// sql := "SELECT * FROM TEST_USER WHERE USER_ID = $1;"

	// pstatement, err := Db.Prepare(sql)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer pstatement.Close()

	// qeury := 1
	// var testUser TestUser

	// err = pstatement.QueryRow(qeury).Scan(&testUser.UserID, &testUser.Password)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(testUser.UserID, testUser.Password)
	router := gin.Default()
    router.LoadHTMLGlob("frontend/*.html")

    router.GET("/", func(ctx *gin.Context){
		sql := "SELECT * FROM TEST_USER"
		var testUser []TestUser
		rows, err := Db.Query(sql)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		for rows.Next() {
			var user TestUser
			err := rows.Scan(&user.UserID, &user.Password)
			if err != nil {
				log.Fatal(err)
			}
			testUser = append(testUser, user)
		}

		ctx.HTML(200, "index.html", gin.H{
			"testUser": testUser,
		})
    })

	router.POST("/add", func(ctx *gin.Context){
		userID := ctx.PostForm("userID")
		password := ctx.PostForm("password")

		sql := "INSERT INTO TEST_USER (user_id, user_password) VALUES ($1, $2)"
		_, err := Db.Exec(sql, userID, password)
		if err != nil {
			log.Fatal(err)
		}

		ctx.Redirect(302, "/")
	})

    router.Run()
}
