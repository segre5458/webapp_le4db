package routes

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/segre5458/webapp_le4db/backend/controllers"
)

func SetupRoutes(router *gin.Engine, context *gin.Context, Db *sql.DB) {
	router.GET("/", func(c *gin.Context) {
        controllers.Index(c, Db)
    })
    router.POST("/add", func(c *gin.Context) {
        controllers.AddRack(c, Db)
    })
    router.GET("/search", func(c *gin.Context) {
        controllers.SearchRack(c, Db)
    })
    router.POST("/delete", func(c *gin.Context) {
        controllers.DeleteRack(c, Db)
    })
	router.GET("/rack/:unitNumber", func(c *gin.Context) {
		controllers.Rack(c, Db)
	})
}
