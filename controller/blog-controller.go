package controller

import (
	"otoklix/dto"
	"otoklix/service"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

type BlogController interface {
	CreateBlog(c *gin.Context)
	GetBlogs(c *gin.Context)
	GetBlog(c *gin.Context)
	UpdateBlog(c *gin.Context)
	DeleteBlog(c *gin.Context)
}

type blogController struct{}

var (
	blogService service.BlogService = service.NewBlogService()
)

func NewBlogController() BlogController {
	return &blogController{}
}

func (*blogController) CreateBlog(c *gin.Context) {
	var blog dto.Blogs
	c.Bind(&blog)

	if blog.Title != "" && blog.Content != "" {
		// convert time to string
		today, _ := time.Now().Local().MarshalText()
		// store request to dto struct
		newBlog := dto.Blogs{
			Title:       blog.Title,
			Content:     blog.Content,
			PublishedAt: string(today),
			CreatedAt:   string(today),
			UpdatedAt:   string(today),
		}

		res := blogService.CreateBlog(newBlog)
		// display success
		c.JSON(201, res)
	} else {
		// display error
		c.JSON(422, gin.H{"error": "Fields are empty"})
	}
}

func (*blogController) GetBlogs(c *gin.Context) {
	res := blogService.GetBlogs()

	// Display JSON result
	c.JSON(200, res)
}

func (*blogController) GetBlog(c *gin.Context) {
	// Get id blog
	id, _ := strconv.Atoi(c.Params.ByName("id"))

	blog := blogService.GetBlog(id)

	if blog.ID != 0 {
		// Display JSON result
		c.JSON(200, blog)
	} else {
		// Display JSON error
		c.JSON(404, gin.H{"error": "Post not found"})
	}
}

func (*blogController) UpdateBlog(c *gin.Context) {
	// Get id blog
	id, _ := strconv.Atoi(c.Params.ByName("id"))

	blog := blogService.GetBlog(id)

	if blog.Title != "" && blog.Content != "" {
		if blog.ID != 0 {
			var newBlog dto.Blogs
			c.Bind(&newBlog)

			res := blogService.UpdateBlog(blog, newBlog)

			// output message
			output := dto.Blogs{
				ID:          blog.ID,
				Title:       newBlog.Title,
				Content:     newBlog.Content,
				PublishedAt: blog.PublishedAt,
				CreatedAt:   blog.CreatedAt,
				UpdatedAt:   res.UpdatedAt,
			}
			// Display modified data in JSON
			c.JSON(200, output)
		} else {
			// Display JSON error
			c.JSON(404, gin.H{"error": "Post not found"})
		}
	} else {
		// Display JSON error
		c.JSON(422, gin.H{"error": "Fields are empty"})
	}
}

func (*blogController) DeleteBlog(c *gin.Context) {
	// Get id blog
	id, _ := strconv.Atoi(c.Params.ByName("id"))

	blog := blogService.GetBlog(id)

	if blog.ID != 0 {
		blogService.DeleteBlog(blog)
		// output message
		output := dto.Blogs{
			ID:          blog.ID,
			Title:       blog.Title,
			Content:     blog.Content,
			PublishedAt: blog.PublishedAt,
			CreatedAt:   blog.CreatedAt,
			UpdatedAt:   blog.UpdatedAt,
		}
		// Display JSON result
		c.JSON(200, output)
	} else {
		// Display JSON error
		c.JSON(404, gin.H{"error": "Post not found"})
	}
}
