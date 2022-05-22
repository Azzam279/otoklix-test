package tests

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"time"

	"github.com/gorilla/mux"

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
	TITLE   string = "Ini judul testing"
	CONTENT string = "Ini konten testing"
)

func TestCreateBlogValid(t *testing.T) {
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

	assert.Equal(t, http.StatusOK, status, "should be equal & return code 200")

	// Decode the HTTP response
	var blog dto.Blogs
	json.NewDecoder(io.Reader(response.Body)).Decode(&blog)

	// Assert HTTP response
	assert.NotNil(t, blog.ID, "should return not null")

	// Clean up database
	cleanUp(blog)
}

func TestCreateBlogInvalid(t *testing.T) {
	// Create a new HTTP POST request
	var jsonReq = []byte(`{"title": "` + TITLE + `", "content": ""}`)
	req, _ := http.NewRequest("POST", "/posts", bytes.NewBuffer(jsonReq))

	// Assign HTTP handle function (controller CreateBlog function)
	handler := http.HandlerFunc(blogController.CreateBlog)

	// Record HTTP response (httptest)
	response := httptest.NewRecorder()

	// Dispatch the HTTP request
	handler.ServeHTTP(response, req)

	// Add assertions on the HTTP status code and the response
	status := response.Code

	assert.Equal(t, http.StatusInternalServerError, status, "should be equal & return code 500")
}

func TestGetBlogs(t *testing.T) {
	// Insert new post
	BlogId := setup()

	// Create a new HTTP GET request
	req, _ := http.NewRequest("GET", "/posts", nil)

	// Assign HTTP handle function (controller GetBlogs function)
	handler := http.HandlerFunc(blogController.GetBlogs)

	// Record HTTP response (httptest)
	response := httptest.NewRecorder()

	// Dispatch the HTTP request
	handler.ServeHTTP(response, req)

	// Add assertions on the HTTP status code and the response
	status := response.Code

	// Assert HTTP response
	assert.Equal(t, http.StatusOK, status, "should be equal & return code 200")

	// Decode the HTTP response
	var blogs []dto.Blogs
	json.NewDecoder(io.Reader(response.Body)).Decode(&blogs)

	// Assert HTTP response
	assert.NotNil(t, blogs[0].ID, "should return not null (ID)")
	assert.NotNil(t, blogs[0].Title, "should return not null (Title)")
	assert.NotNil(t, blogs[0].Content, "should return not null (Content)")

	// Clean up database
	currentBlog := dto.Blogs{
		ID: BlogId,
	}
	cleanUp(currentBlog)
}

func TestGetBlogValid(t *testing.T) {
	// Insert new post
	BlogId := setup()

	// Create a new HTTP GET request
	req, _ := http.NewRequest("GET", "/posts", nil)
	req = mux.SetURLVars(req, map[string]string{
		"id": strconv.Itoa(BlogId),
	})

	// Assign HTTP handle function (controller GetBlog function)
	handler := http.HandlerFunc(blogController.GetBlog)

	// Record HTTP response (httptest)
	response := httptest.NewRecorder()

	// Dispatch the HTTP request
	handler.ServeHTTP(response, req)

	// Add assertions on the HTTP status code and the response
	status := response.Code

	// Assert HTTP response
	assert.Equal(t, http.StatusOK, status, "should be equal & return code 200")

	// Decode the HTTP response
	var blog dto.Blogs
	json.NewDecoder(io.Reader(response.Body)).Decode(&blog)

	// Assert HTTP response
	assert.NotNil(t, blog.ID, "should return not null (ID)")
	assert.NotNil(t, blog.Title, "should return not null (Title)")
	assert.NotNil(t, blog.Content, "should return not null (Content)")

	// Clean up database
	currentBlog := dto.Blogs{
		ID: BlogId,
	}
	cleanUp(currentBlog)
}

func TestGetBlogInvalid(t *testing.T) {
	// Create a new HTTP GET request
	req, _ := http.NewRequest("GET", "/posts", nil)

	// Assign HTTP handle function (controller GetBlog function)
	handler := http.HandlerFunc(blogController.GetBlog)

	// Record HTTP response (httptest)
	response := httptest.NewRecorder()

	// Dispatch the HTTP request
	handler.ServeHTTP(response, req)

	// Add assertions on the HTTP status code and the response
	status := response.Code

	// Assert HTTP response
	assert.Equal(t, http.StatusInternalServerError, status, "should be equal & return code 500")
}

func TestUpdateBlogValid(t *testing.T) {
	// Insert new post
	BlogId := setup()

	// Create a new HTTP PUT request
	var jsonReq = []byte(`{"title": "` + TITLE + `", "content": "` + CONTENT + `"}`)
	req, _ := http.NewRequest("PUT", "/posts", bytes.NewBuffer(jsonReq))
	req = mux.SetURLVars(req, map[string]string{
		"id": strconv.Itoa(BlogId),
	})

	// Assign HTTP handle function (controller UpdateBlog function)
	handler := http.HandlerFunc(blogController.UpdateBlog)

	// Record HTTP response (httptest)
	response := httptest.NewRecorder()

	// Dispatch the HTTP request
	handler.ServeHTTP(response, req)

	// Add assertions on the HTTP status code and the response
	status := response.Code

	// Assert HTTP response
	assert.Equal(t, http.StatusOK, status, "should be equal & return code 200")

	// Decode the HTTP response
	var blog dto.Blogs
	json.NewDecoder(io.Reader(response.Body)).Decode(&blog)

	// Assert HTTP response
	assert.NotNil(t, blog.ID, "should return not null (ID)")
	assert.NotNil(t, blog.Title, "should return not null (Title)")
	assert.NotNil(t, blog.Content, "should return not null (Content)")

	// Clean up database
	currentBlog := dto.Blogs{
		ID: BlogId,
	}
	cleanUp(currentBlog)
}

func TestUpdateBlogInvalid(t *testing.T) {
	// Insert new post
	BlogId := setup()

	// Create a new HTTP PUT request
	var jsonReq = []byte(`{"title": "` + TITLE + `", "content": ""}`)
	req, _ := http.NewRequest("PUT", "/posts", bytes.NewBuffer(jsonReq))
	req = mux.SetURLVars(req, map[string]string{
		"id": strconv.Itoa(BlogId),
	})

	// Assign HTTP handle function (controller UpdateBlog function)
	handler := http.HandlerFunc(blogController.UpdateBlog)

	// Record HTTP response (httptest)
	response := httptest.NewRecorder()

	// Dispatch the HTTP request
	handler.ServeHTTP(response, req)

	// Add assertions on the HTTP status code and the response
	status := response.Code

	// Clean up database
	currentBlog := dto.Blogs{
		ID: BlogId,
	}
	cleanUp(currentBlog)

	// Assert HTTP response
	assert.Equal(t, http.StatusBadRequest, status, "should be equal & return code 400")
}

func TestDeleteBlogValid(t *testing.T) {
	// Insert new post
	BlogId := setup()

	// Create a new HTTP DELETE request
	req, _ := http.NewRequest("DELETE", "/posts", nil)
	req = mux.SetURLVars(req, map[string]string{
		"id": strconv.Itoa(BlogId),
	})

	// Assign HTTP handle function (controller DeleteBlog function)
	handler := http.HandlerFunc(blogController.DeleteBlog)

	// Record HTTP response (httptest)
	response := httptest.NewRecorder()

	// Dispatch the HTTP request
	handler.ServeHTTP(response, req)

	// Add assertions on the HTTP status code and the response
	status := response.Code

	// Assert HTTP response
	assert.Equal(t, http.StatusOK, status, "should be equal & return code 200")
}

func TestDeleteBlogInvalid(t *testing.T) {
	// Create a new HTTP DELETE request
	req, _ := http.NewRequest("DELETE", "/posts", nil)

	// Assign HTTP handle function (controller DeleteBlog function)
	handler := http.HandlerFunc(blogController.DeleteBlog)

	// Record HTTP response (httptest)
	response := httptest.NewRecorder()

	// Dispatch the HTTP request
	handler.ServeHTTP(response, req)

	// Add assertions on the HTTP status code and the response
	status := response.Code

	// Assert HTTP response
	assert.Equal(t, http.StatusInternalServerError, status, "should be equal & return code 500")
}

func setup() int {
	today, _ := time.Now().Local().MarshalText()
	blog := dto.Blogs{
		Title:       TITLE,
		Content:     CONTENT,
		PublishedAt: string(today),
		CreatedAt:   string(today),
		UpdatedAt:   string(today),
	}
	d := blogRepo.CreateBlog(blog)
	return d.ID
}

func cleanUp(blog dto.Blogs) {
	newBlog := dto.Blogs{
		ID: blog.ID,
	}
	blogRepo.DeleteBlog(&newBlog)
}
