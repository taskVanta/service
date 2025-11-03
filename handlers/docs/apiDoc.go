package handlers

import "github.com/gin-gonic/gin"

func ServeAPIDocs(c *gin.Context) {
	c.File("./static/index.html")
}
