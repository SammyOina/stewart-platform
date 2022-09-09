package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	r := gin.Default()
	authConfig := cors.DefaultConfig()
	authConfig.AllowAllOrigins = true
	r.Use(cors.New(authConfig))
	return r
}
