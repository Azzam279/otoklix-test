package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.POST("/posts", createBlog)
	r.GET("/posts", getBlogs)
	r.GET("/posts/:id", getBlog)
	r.PUT("/posts/:id", updateBlog)
	r.DELETE("/posts/:id", deleteBlog)

	r.Run(":8081")
}
