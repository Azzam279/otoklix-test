package datastruct

type Blogs struct {
	ID          int    `form:"id" json:"id"`
	Title       string `form:"title" json:"title"`
	Content     string `form:"content" json:"content"`
	PublishedAt string `form:"published_at" json:"published_at"`
	CreatedAt   string `form:"created_at" json:"created_at"`
	UpdatedAt   string `form:"updated_at" json:"updated_at"`
}
