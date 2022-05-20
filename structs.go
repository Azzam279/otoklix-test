package main

type Blogs struct {
	ID          int    `gorm:"AUTO_INCREMENT" form:"id" json:"id"`
	Title       string `gorm:"not null" form:"title" json:"title"`
	Content     string `gorm:"not null" form:"content" json:"content"`
	PublishedAt string `form:"published_at" json:"published_at"`
	CreatedAt   string `form:"created_at" json:"created_at"`
	UpdatedAt   string `form:"updated_at" json:"updated_at"`
}
