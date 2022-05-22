package service

import (
	"otoklix/dto"
	"otoklix/repository"
	"time"
)

type BlogService interface {
	CreateBlog(blog dto.Blogs) *dto.Blogs
	GetBlogs() *[]dto.Blogs
	GetBlog(id int) *dto.Blogs
	UpdateBlog(blog *dto.Blogs, newBlog dto.Blogs) *dto.Blogs
	DeleteBlog(blog *dto.Blogs) *dto.Blogs
}

type blogService struct{}

var (
	repo repository.BlogRepository
)

func NewBlogService(repository repository.BlogRepository) BlogService {
	repo = repository
	return &blogService{}
}

func (*blogService) CreateBlog(blog dto.Blogs) *dto.Blogs {
	blogInfo := dto.Blogs{
		ID:          blog.ID,
		Title:       blog.Title,
		Content:     blog.Content,
		PublishedAt: blog.PublishedAt,
		CreatedAt:   blog.CreatedAt,
		UpdatedAt:   blog.UpdatedAt,
	}

	return repo.CreateBlog(blogInfo)
}

func (*blogService) GetBlogs() *[]dto.Blogs {
	return repo.GetBlogs()
}

func (*blogService) GetBlog(id int) *dto.Blogs {
	return repo.GetBlog(id)
}

func (*blogService) UpdateBlog(blog *dto.Blogs, newBlog dto.Blogs) *dto.Blogs {
	today, _ := time.Now().Local().MarshalText()
	result := dto.Blogs{
		Title:     newBlog.Title,
		Content:   newBlog.Content,
		UpdatedAt: string(today),
	}

	return repo.UpdateBlog(blog, result)
}

func (*blogService) DeleteBlog(blog *dto.Blogs) *dto.Blogs {
	return repo.DeleteBlog(blog)
}
