package main

import (
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func createBlog(c *gin.Context) {
	// Connection to the database
	db := initDb()
	// Close connection database
	defer db.Close()

	var blog Blogs
	c.Bind(&blog)

	if blog.Title != "" && blog.Content != "" {
		// convert time to string
		today, _ := time.Now().Local().MarshalText()
		blog.PublishedAt = string(today)
		blog.CreatedAt = string(today)
		blog.UpdatedAt = string(today)
		// insert data
		db.Create(&blog)
		// display success
		c.JSON(201, blog)
	} else {
		// display error
		c.JSON(422, gin.H{"error": "Fields are empty"})
	}
}

func getBlogs(c *gin.Context) {
	// Connection to the database
	db := initDb()
	// Close connection database
	defer db.Close()

	// SELECT * FROM blogs;
	var blogs []Blogs
	db.Find(&blogs)

	// Display JSON result
	c.JSON(200, blogs)
}

func getBlog(c *gin.Context) {
	// Connection to the database
	db := initDb()
	// Close connection database
	defer db.Close()

	// Get id blog
	id := c.Params.ByName("id")
	var blog Blogs
	// SELECT * FROM blogs WHERE id = 1;
	db.First(&blog, id)

	if blog.ID != 0 {
		// Display JSON result
		c.JSON(200, blog)
	} else {
		// Display JSON error
		c.JSON(404, gin.H{"error": "Post not found"})
	}
}

func updateBlog(c *gin.Context) {
	// Connection to the database
	db := initDb()
	// Close connection database
	defer db.Close()

	// Get id blog
	id := c.Params.ByName("id")
	var blog Blogs
	// SELECT * FROM blogs WHERE id = 1;
	db.First(&blog, id)

	if blog.Title != "" && blog.Content != "" {
		if blog.ID != 0 {
			var newBlog Blogs
			c.Bind(&newBlog)

			today, _ := time.Now().Local().MarshalText()
			result := Blogs{
				Title:     newBlog.Title,
				Content:   newBlog.Content,
				UpdatedAt: string(today),
			}
			// update data
			db.Model(&blog).Updates(result)

			// output message
			output := Blogs{
				ID:          blog.ID,
				Title:       newBlog.Title,
				Content:     newBlog.Content,
				PublishedAt: blog.PublishedAt,
				CreatedAt:   blog.CreatedAt,
				UpdatedAt:   string(today),
			}
			// Display modified data in JSON message "success"
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

func deleteBlog(c *gin.Context) {
	// Connection to the database
	db := initDb()
	// Close connection database
	defer db.Close()

	// Get id blog
	id := c.Params.ByName("id")
	var blog Blogs
	// SELECT * FROM blogs WHERE id = 1;
	db.First(&blog, id)

	if blog.ID != 0 {
		// delete data
		db.Delete(&blog)
		// output message
		output := Blogs{
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
