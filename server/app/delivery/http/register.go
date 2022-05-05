package http

import "github.com/gin-gonic/gin"

func RegisterHTTPEndpoints(router *gin.RouterGroup) {
	h := NewHandler()

	router.GET("/hello", h.HelloWorld)
}
