package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-contrib/cors"

	"github.com/segre5458/webapp_le4db/backend/routes"
	
)

func main() {
	var Db *sql.DB

	Db, err := sql.Open("postgres", "host=postgres user=postgres password=postgres dbname=postgres sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer Db.Close()

	router := gin.Default()
    router.LoadHTMLGlob("frontend/*.html")

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	router.Use(cors.New(config))

	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))

	context := &gin.Context{}
	routes.SetupRoutes(router,context, Db)

    router.Run(":8080")
}
