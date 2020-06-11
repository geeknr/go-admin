package route

import (
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go-admin-svr/internal/middleware"
	"go-admin-svr/internal/route/admin"
	"go-admin-svr/internal/route/api"
	"net/http"
)

func Test(c *gin.Context) {
	c.JSON(http.StatusOK, "Hello HMSC!")
}

func InitApi() *gin.Engine {
	m := gin.New()
	m.Use(gin.Recovery())

	if viper.GetBool("gzip") {
		m.Use(gzip.Gzip(gzip.DefaultCompression))
	}
	if viper.GetBool("cors.enable") {
		m.Use(middleware.Cors())
	}
	v1 := m.Group("/v1")
	v1.GET("/version/info", api.GetVersionInfo)
	return m
}

func InitAdmin() *gin.Engine {

	m := gin.New()
	m.Use(gin.Recovery())

	if viper.GetBool("gzip") {
		m.Use(gzip.Gzip(gzip.DefaultCompression))
	}

	if viper.GetBool("cors.enable") {
		m.Use(middleware.Cors())
	}

	v1 := m.Group("/v1")
	v1.GET("/version/info", admin.GetVersionInfo)

	return m
}
