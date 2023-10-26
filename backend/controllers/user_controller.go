package controllers

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"

	"github.com/segre5458/webapp_le4db/backend/models"
)

func Index(ctx *gin.Context, Db *sql.DB) {
	sql := "SELECT * FROM TEST_USER"
	var testUser []models.TestUser
	rows, err := Db.Query(sql)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var user models.TestUser
		err := rows.Scan(&user.UserID, &user.Password)
		if err != nil {
			log.Fatal(err)
		}
		testUser = append(testUser, user)
	}

	ctx.HTML(200, "index.html", gin.H{
		"testUser": testUser,
	})
}

func AddUser(ctx *gin.Context, Db *sql.DB) {
	userID := ctx.PostForm("userID")
	password := ctx.PostForm("password")

	sql := "INSERT INTO TEST_USER (user_id, user_password) VALUES ($1, $2)"
	_, err := Db.Exec(sql, userID, password)
	if err != nil {
		log.Fatal(err)
	}

	ctx.Redirect(302, "/")
}

func SearchUser(ctx *gin.Context, Db *sql.DB) {
	userID := ctx.Query("userID")

	sql := "SELECT * FROM TEST_USER WHERE USER_ID = $1"
	var testUser []models.TestUser
	rows, err := Db.Query(sql, userID)
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var user models.TestUser
		err := rows.Scan(&user.UserID, &user.Password)
		if err != nil {
			if err.Error() == "sql: no rows in result set" {
				ctx.HTML(200, "index.html", gin.H{
					"message": "No User",
				})
			} else {
				log.Fatal(err)
			}
		}
		testUser = append(testUser, user)
	}
	ctx.HTML(200, "index.html", gin.H{
		"testUser": testUser,
	})
}

func DeleteUser(ctx *gin.Context, Db *sql.DB) {
	userID := ctx.PostForm("userID")

	sql := "DELETE FROM TEST_USER WHERE USER_ID = $1"
	_, err := Db.Exec(sql, userID)
	if err != nil {
		log.Fatal(err)
	}

	ctx.Redirect(302, "/")
}
