package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"

	"github.com/segre5458/webapp_le4db/backend/controllers"
	"github.com/segre5458/webapp_le4db/backend/middleware"
)

func SetupRoutes(router *gin.Engine, context *gin.Context, Db *sql.DB) {
    router.GET("/login", func(c *gin.Context) {
        c.HTML(200, "index.html", gin.H{})
    })
    router.POST("/login", func(c *gin.Context) {
        controllers.Login(c, Db)
    })

    authGroup := router.Group("/")
    authGroup.Use(middleware.AuthMiddleware)

	authGroup.GET("/", func(c *gin.Context) {
        controllers.Index(c, Db)
    })

    authGroup.POST("/add", func(c *gin.Context) {
        controllers.AddRack(c, Db)
    })
    authGroup.GET("/search", func(c *gin.Context) {
        controllers.SearchRack(c, Db)
    })
    authGroup.POST("/delete", func(c *gin.Context) {
        controllers.DeleteRack(c, Db)
    })
    authGroup.GET("/edit/:unitNumber", func(c *gin.Context) {
        controllers.EditRack(c, Db)
    })
    authGroup.POST("/update/:unitNumber", func(c *gin.Context) {
        controllers.UpdateRack(c, Db)
    })


	authGroup.GET("/rack/:unitNumber", func(c *gin.Context) {
		controllers.Rack(c, Db)
	})

    authGroup.POST("/rack/:unitNumber/server/add", func(c *gin.Context) {
        controllers.AddServer(c, Db)
    })
    authGroup.GET("/rack/:unitNumber/server/search", func(c *gin.Context) {
        controllers.SearchServer(c, Db)
    })
    authGroup.POST("/rack/:unitNumber/server/delete", func(c *gin.Context) {
        controllers.DeleteServer(c, Db)
    })
    authGroup.GET("/rack/:unitNumber/server/edit/:serverName", func(c *gin.Context) {
        controllers.EditServer(c, Db)
    })
    authGroup.POST("/rack/:unitNumber/server/update/:serverName", func(c *gin.Context) {
        controllers.UpdateServer(c, Db)
    })

    authGroup.POST("/rack/:unitNumber/networkDevice/add", func(c *gin.Context) {
        controllers.AddNetworkDevice(c, Db)
    })
    authGroup.GET("/rack/:unitNumber/networkDevice/search", func(c *gin.Context) {
        controllers.SearchNetworkDevice(c, Db)
    })
    authGroup.POST("/rack/:unitNumber/networkDevice/delete", func(c *gin.Context) {
        controllers.DeleteNetworkDevice(c, Db)
    })
    authGroup.GET("/rack/:unitNumber/networkDevice/edit/:deviceName", func(c *gin.Context) {
        controllers.EditNetworkDevice(c, Db)
    })
    authGroup.POST("/rack/:unitNumber/networkDevice/update/:deviceName", func(c *gin.Context) {
        controllers.UpdateNetworkDevice(c, Db)
    })
}
