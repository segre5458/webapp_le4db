package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get("authenticatedUser")
	if user == nil {
		c.Redirect(302, "/login")
		c.Abort()
		return
	}
	c.Next()
}
