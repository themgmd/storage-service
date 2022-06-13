package http

import "github.com/gin-gonic/gin"

func FileHandlerInit(api *gin.RouterGroup) {
	api.GET("/file/:id", getFileById)
	api.POST("/file", uploadFile)
}

func getFileById(c *gin.Context) {
	panic("Implement get file")
}

func uploadFile(c *gin.Context) {
	panic("Implement create file")
}
