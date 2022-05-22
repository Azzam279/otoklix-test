package controller

import (
	"encoding/json"
	"net/http"
	"otoklix/dto"
	"otoklix/service"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

type BlogController interface {
	CreateBlog(response http.ResponseWriter, request *http.Request)
	GetBlogs(response http.ResponseWriter, request *http.Request)
	GetBlog(response http.ResponseWriter, request *http.Request)
	UpdateBlog(response http.ResponseWriter, request *http.Request)
	DeleteBlog(response http.ResponseWriter, request *http.Request)
}

type blogController struct{}

var (
	blogService service.BlogService
)

func NewBlogController(service service.BlogService) BlogController {
	blogService = service
	return &blogController{}
}

func (*blogController) CreateBlog(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	var blog dto.Blogs
	err := json.NewDecoder(request.Body).Decode(&blog)
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}

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
		response.WriteHeader(http.StatusOK)
		json.NewEncoder(response).Encode(res)
	} else {
		// display error
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"error": "Fields are empty"}`))
	}
}

func (*blogController) GetBlogs(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	res := blogService.GetBlogs()

	// Display JSON result
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(res)
}

func (*blogController) GetBlog(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	// Get id blog
	params := mux.Vars(request)
	id, _ := strconv.Atoi(params["id"])

	blog := blogService.GetBlog(id)

	if blog.ID != 0 {
		// Display JSON result
		response.WriteHeader(http.StatusOK)
		json.NewEncoder(response).Encode(blog)
	} else {
		// Display JSON error
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"error": "Post not found"}`))
	}
}

func (*blogController) UpdateBlog(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	// Get id blog
	params := mux.Vars(request)
	id, _ := strconv.Atoi(params["id"])
	// Get blog by id
	blog := blogService.GetBlog(id)

	if blog.Title != "" && blog.Content != "" {
		if blog.ID != 0 {
			var newBlog dto.Blogs
			err := json.NewDecoder(request.Body).Decode(&newBlog)
			if err != nil {
				http.Error(response, err.Error(), http.StatusBadRequest)
				return
			}

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
			response.WriteHeader(http.StatusOK)
			json.NewEncoder(response).Encode(output)
		} else {
			// Display JSON error
			response.WriteHeader(http.StatusInternalServerError)
			response.Write([]byte(`{"error": "Post not found"}`))
		}
	} else {
		// Display JSON error
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte(`{"error": "Fields are empty"}`))
	}
}

func (*blogController) DeleteBlog(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	// Get id blog
	params := mux.Vars(request)
	id, _ := strconv.Atoi(params["id"])

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
		response.WriteHeader(http.StatusOK)
		json.NewEncoder(response).Encode(output)
	} else {
		// Display JSON error
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"error": "Post not found"}`))
	}
}
