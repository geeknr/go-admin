package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetVersionInfo(c *gin.Context) {
	c.JSON(http.StatusOK, "Hello, ADMIN! v1.0")
}
