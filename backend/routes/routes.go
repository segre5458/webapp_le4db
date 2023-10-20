package routes

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/segre5458/webapp_le4db/backend/controllers"
)

func SetupRoutes(router *gin.Engine, Db *sql.DB) {
	router.GET("/", controllers.Index, Db)
	router.POST("/add", controllers.AddUser, Db)
	router.GET("/search", controllers.SearchUser, Db)
	router.POST("/delete", controllers.DeleteUser, Db)
}
