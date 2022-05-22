package main

import (
	"otoklix/controller"
	"otoklix/http"
	"otoklix/repository"
	"otoklix/service"
)

var (
	blogRepository repository.BlogRepository = repository.NewBlogRepository()
	blogService    service.BlogService       = service.NewBlogService(blogRepository)
	blogController controller.BlogController = controller.NewBlogController(blogService)
	httpRouter     http.Router               = http.NewMaxRouter()
)

func main() {
	httpRouter.GET("/posts", blogController.GetBlogs)
	httpRouter.GET("/posts/{id}", blogController.GetBlog)
	httpRouter.POST("/posts", blogController.CreateBlog)
	httpRouter.PUT("/posts/{id}", blogController.UpdateBlog)
	httpRouter.DELETE("/posts/{id}", blogController.DeleteBlog)

	httpRouter.SERVE(":8081")
}
