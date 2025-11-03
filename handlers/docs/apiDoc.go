package handlers

import "github.com/gin-gonic/gin"

func ServeAPIDocs(c *gin.Context) {
	c.Header("Referrer-Policy", "strict-origin-when-cross-origin")
	c.File("./static/index.html")
}
