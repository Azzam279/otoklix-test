package repository

import (
	"otoklix/datastruct"
)

type BlogRepository interface {
	CreateBlog(blog datastruct.Blogs) *datastruct.Blogs
	GetBlogs() *[]datastruct.Blogs
	GetBlog(id int) *datastruct.Blogs
	UpdateBlog(blog *datastruct.Blogs, result datastruct.Blogs) *datastruct.Blogs
	DeleteBlog(blog *datastruct.Blogs) *datastruct.Blogs
}

type blogRepository struct{}

func NewBlogRepository() BlogRepository {
	return &blogRepository{}
}

func (*blogRepository) CreateBlog(blog datastruct.Blogs) *datastruct.Blogs {
	// Connection to the database
	db := initDb()
	// Close connection database
	defer db.Close()

	// insert data
	db.Create(&blog)

	return &blog
}

func (*blogRepository) GetBlogs() *[]datastruct.Blogs {
	// Connection to the database
	db := initDb()
	// Close connection database
	defer db.Close()

	var blogs []datastruct.Blogs
	// SELECT * FROM blogs;
	db.Find(&blogs)

	return &blogs
}

func (*blogRepository) GetBlog(id int) *datastruct.Blogs {
	// Connection to the database
	db := initDb()
	// Close connection database
	defer db.Close()

	var blog datastruct.Blogs
	// SELECT * FROM blogs WHERE id = 1;
	db.First(&blog, id)

	return &blog
}

func (*blogRepository) UpdateBlog(blog *datastruct.Blogs, result datastruct.Blogs) *datastruct.Blogs {
	// Connection to the database
	db := initDb()
	// Close connection database
	defer db.Close()

	// update data
	db.Model(&blog).Updates(result)

	return blog
}

func (*blogRepository) DeleteBlog(blog *datastruct.Blogs) *datastruct.Blogs {
	// Connection to the database
	db := initDb()
	// Close connection database
	defer db.Close()

	// delete data
	db.Delete(&blog)

	return blog
}
