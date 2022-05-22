package repository

import (
	"otoklix/dto"
)

type BlogRepository interface {
	CreateBlog(blog dto.Blogs) *dto.Blogs
	GetBlogs() *[]dto.Blogs
	GetBlog(id int) *dto.Blogs
	UpdateBlog(blog *dto.Blogs, result dto.Blogs) *dto.Blogs
	DeleteBlog(blog *dto.Blogs) *dto.Blogs
}

type blogRepository struct{}

func NewBlogRepository() BlogRepository {
	return &blogRepository{}
}

func (*blogRepository) CreateBlog(blog dto.Blogs) *dto.Blogs {
	// Connection to the database
	db := initDb()

	// insert data
	db.Create(&blog)

	return &blog
}

func (*blogRepository) GetBlogs() *[]dto.Blogs {
	// Connection to the database
	db := initDb()

	var blogs []dto.Blogs
	// SELECT * FROM blogs;
	db.Find(&blogs)

	return &blogs
}

func (*blogRepository) GetBlog(id int) *dto.Blogs {
	// Connection to the database
	db := initDb()

	var blog dto.Blogs
	// SELECT * FROM blogs WHERE id = 1;
	db.First(&blog, id)

	return &blog
}

func (*blogRepository) UpdateBlog(blog *dto.Blogs, result dto.Blogs) *dto.Blogs {
	// Connection to the database
	db := initDb()

	// update data
	db.Model(&blog).Updates(result)

	return blog
}

func (*blogRepository) DeleteBlog(blog *dto.Blogs) *dto.Blogs {
	// Connection to the database
	db := initDb()

	// delete data
	db.Delete(&blog)

	return blog
}
