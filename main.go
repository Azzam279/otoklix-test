package main

import (
	"otoklix/controller"

	"github.com/gin-gonic/gin"
)

var (
	blogController controller.BlogController = controller.NewBlogController()
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.POST("/posts", blogController.CreateBlog)
	r.GET("/posts", blogController.GetBlogs)
	r.GET("/posts/:id", blogController.GetBlog)
	r.PUT("/posts/:id", blogController.UpdateBlog)
	r.DELETE("/posts/:id", blogController.DeleteBlog)

	r.Run(":8081")
}
