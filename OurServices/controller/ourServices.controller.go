package OurServices

import "github.com/gin-gonic/gin"

func Routers() {
	router := gin.Default()
	router.GET("/Services/OurCurrentServices")
	router.POST("/AddService")
	router.GET("/GetById")
	router.Run("localhost:8080")
}
