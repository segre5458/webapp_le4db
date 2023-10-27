package controllers

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"

	"github.com/segre5458/webapp_le4db/backend/models"
)

func Login(ctx *gin.Context, DB *sql.DB) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	
	var loginRequest models.LoginRequest
	loginRequest.Username = username
	loginRequest.Password = password

	sql := "SELECT * FROM DCUSER WHERE name = $1 AND password = $2"
	row := DB.QueryRow(sql, loginRequest.Username, loginRequest.Password)

	var dcuser models.DCUser
	err := row.Scan(&dcuser.Uid, &dcuser.Name, &dcuser.Password, &dcuser.Email)
	if err != nil {
		log.Println(err)
		ctx.HTML(400, "login.html", gin.H{"error": "Invalid username or password"})
		return
	}

	session := sessions.Default(ctx)
	session.Set("authenticatedUser", loginRequest.Username)
	session.Save()

	ctx.Redirect(302, "/")
}
