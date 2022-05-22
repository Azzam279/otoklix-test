package service

import (
	"otoklix/datastruct"
	"otoklix/dto"
	"otoklix/repository"
	"time"
)

type BlogService interface {
	CreateBlog(blog dto.Blogs) *datastruct.Blogs
	GetBlogs() *[]datastruct.Blogs
	GetBlog(id int) *datastruct.Blogs
	UpdateBlog(blog *datastruct.Blogs, newBlog dto.Blogs) *datastruct.Blogs
	DeleteBlog(blog *datastruct.Blogs) *datastruct.Blogs
}

type blogService struct{}

var (
	repo repository.BlogRepository
)

func NewBlogService(repository repository.BlogRepository) BlogService {
	repo = repository
	return &blogService{}
}

func (*blogService) CreateBlog(blog dto.Blogs) *datastruct.Blogs {
	blogInfo := datastruct.Blogs{
		ID:          blog.ID,
		Title:       blog.Title,
		Content:     blog.Content,
		PublishedAt: blog.PublishedAt,
		CreatedAt:   blog.CreatedAt,
		UpdatedAt:   blog.UpdatedAt,
	}

	return repo.CreateBlog(blogInfo)
}

func (*blogService) GetBlogs() *[]datastruct.Blogs {
	return repo.GetBlogs()
}

func (*blogService) GetBlog(id int) *datastruct.Blogs {
	return repo.GetBlog(id)
}

func (*blogService) UpdateBlog(blog *datastruct.Blogs, newBlog dto.Blogs) *datastruct.Blogs {
	today, _ := time.Now().Local().MarshalText()
	result := datastruct.Blogs{
		Title:     newBlog.Title,
		Content:   newBlog.Content,
		UpdatedAt: string(today),
	}

	return repo.UpdateBlog(blog, result)
}

func (*blogService) DeleteBlog(blog *datastruct.Blogs) *datastruct.Blogs {
	return repo.DeleteBlog(blog)
}
