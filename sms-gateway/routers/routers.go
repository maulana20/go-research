package routers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"sms-gateway/routers/api"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Safous Service Gateway")
	})
	r.POST("/v1/sms", api.SmsGateway)
	r.NoRoute(api.Fallback)
	return r
}