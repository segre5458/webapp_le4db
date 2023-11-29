package main

import (
	"database/sql"
	"log"

	middleware "command-line-arguments/home/segre/KU/le4db/webapp/backend/middleware/auth.go"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"

	"github.com/segre5458/webapp_le4db/backend/routes"
	"github.com/segre5458/webapp_le4db/backend/middleware"
)

func main() {
	var Db *sql.DB

	Db, err := sql.Open("postgres", "host=postgres user=postgres password=postgres dbname=postgres sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer Db.Close()

	router := gin.Default()
    // router.LoadHTMLGlob("frontend/*.html")

	router.Use(cors.New(middleware.CORS()))

	// router.Use(gin.Static("/frontend-react/build"))

	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))

	context := &gin.Context{}
	routes.SetupRoutes(router,context, Db)

    router.Run(":8080")
}
