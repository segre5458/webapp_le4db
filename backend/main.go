package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/gin-gonic/gin"
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

	context := &gin.Context{}
	routes.SetupRoutes(router,context, Db)

    router.Run()
}
