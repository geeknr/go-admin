package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetVersionInfo(c *gin.Context) {
	c.JSON(http.StatusOK, "Hello, API! v1.0")
}
