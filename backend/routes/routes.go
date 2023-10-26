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
        controllers.AddUser(c, Db)
    })
    router.GET("/search", func(c *gin.Context) {
        controllers.SearchUser(c, Db)
    })
    router.POST("/delete", func(c *gin.Context) {
        controllers.DeleteUser(c, Db)
    })
}
