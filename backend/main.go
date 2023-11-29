package main

import (
	"database/sql"
	"log"
	"os"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"

	"github.com/segre5458/webapp_le4db/backend/routes"
)

func main() {
	var Db *sql.DB
	host := os.Getenv("POSTGRES_HOST")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")

	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", host, user, password, dbName)

	Db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer Db.Close()

	router := gin.Default()
    router.LoadHTMLGlob("frontend/*.html")

	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))

	context := &gin.Context{}
	routes.SetupRoutes(router,context, Db)

    router.Run(":8080")
}
