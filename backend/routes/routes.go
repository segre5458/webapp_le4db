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

    router.POST("/rack/:unitNumber/server/add", func(c *gin.Context) {
        controllers.AddServer(c, Db)
    })
    router.GET("/rack/:unitNumber/server/search", func(c *gin.Context) {
        controllers.SearchServer(c, Db)
    })
    router.POST("/rack/:unitNumber/server/delete", func(c *gin.Context) {
        controllers.DeleteServer(c, Db)
    })

    router.POST("/rack/:unitNumber/networkDevice/add", func(c *gin.Context) {
        controllers.AddNetworkDevice(c, Db)
    })
    router.GET("/rack/:unitNumber/networkDevice/search", func(c *gin.Context) {
        controllers.SearchNetworkDevice(c, Db)
    })
    router.POST("/rack/:unitNumber/networkDevice/delete", func(c *gin.Context) {
        controllers.DeleteNetworkDevice(c, Db)
    })

}
