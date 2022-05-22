package tests

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"

	"otoklix/controller"
	"otoklix/dto"
	"otoklix/repository"
	"otoklix/service"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	blogRepo       repository.BlogRepository = repository.NewBlogRepository()
	blogService    service.BlogService       = service.NewBlogService(blogRepo)
	blogController controller.BlogController = controller.NewBlogController(blogService)
)

const (
	TITLE   string = "Title 1 testing"
	CONTENT string = "Content 1 testing"
)

func TestCreateBlog(t *testing.T) {
	// Create a new HTTP POST request
	var jsonReq = []byte(`{"title": "` + TITLE + `", "content": "` + CONTENT + `"}`)
	req, _ := http.NewRequest("POST", "/posts", bytes.NewBuffer(jsonReq))

	// Assign HTTP handle function (controller CreateBlog function)
	handler := http.HandlerFunc(blogController.CreateBlog)

	// Record HTTP response (httptest)
	response := httptest.NewRecorder()

	// Dispatch the HTTP request
	handler.ServeHTTP(response, req)

	// Add assertions on the HTTP status code and the response
	status := response.Code

	if status != http.StatusOK {
		t.Fatalf("Handler returned a wrong status code: got %v expected %v", status, http.StatusOK)
	}

	// Decode the HTTP response
	var blog dto.Blogs
	json.NewDecoder(io.Reader(response.Body)).Decode(&blog)

	// Assert HTTP response
	assert.NotNil(t, blog.ID)
	assert.Equal(t, TITLE, blog.Title)
	assert.Equal(t, CONTENT, blog.Content)

	// Clean up database
}

func cleanUp() {

}
