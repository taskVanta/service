package docs

import (
	"net/http"
	"service/static"

	"github.com/gin-gonic/gin"
)

//go:embed ../static/index.html

func ServeAPIDocs(c *gin.Context) {
	data, err := static.StaticFiles.ReadFile("static/index.html") // Must match embedded path
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to load API docs")
		return
	}
	c.Data(http.StatusOK, "text/html; charset=utf-8", data)
}
