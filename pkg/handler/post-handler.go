package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"social-media/pkg/config"
	"social-media/pkg/model"
	"social-media/pkg/service"
	"social-media/pkg/service-impl"
	"social-media/pkg/utils"
	"strconv"
)

var postService service.PostService

func init() {
	db := config.GetDb()
	postService = service_impl.NewPostServiceImpl(db)
}

type PostHandler struct {
}

func (h *PostHandler) CreatePost(w http.ResponseWriter, r *http.Request) {
	var post model.Post
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&post); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if err := postService.CreatePost(&post); err != nil {
		http.Error(w, "Error creating post", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *PostHandler) GetPostByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	postID, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid post ID", http.StatusBadRequest)
		return
	}

	post, err := postService.GetPostByID(uint(postID))
	if err != nil {
		http.Error(w, "Post not found", http.StatusNotFound)
		return
	}

	utils.JsonResponse(w, post)
}

func (h *PostHandler) GetAllPosts(w http.ResponseWriter, r *http.Request) {
	posts, err := postService.GetAllPosts()
	if err != nil {
		http.Error(w, "Error fetching posts", http.StatusInternalServerError)
		return
	}

	utils.JsonResponse(w, posts)
}

func (h *PostHandler) UpdatePost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	postID, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid post ID", http.StatusBadRequest)
		return
	}

	var post model.Post
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&post); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	post.ID = uint(postID)
	if err := postService.UpdatePost(&post); err != nil {
		http.Error(w, "Error updating post", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *PostHandler) DeletePost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	postID, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid post ID", http.StatusBadRequest)
		return
	}

	if err := postService.DeletePost(uint(postID)); err != nil {
		http.Error(w, "Error deleting post", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *PostHandler) GetAllPostsByUserID(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	userId, err := strconv.ParseUint(vars["userId"], 10, 64)
	if err != nil {
		http.Error(writer, "Invalid post ID", http.StatusBadRequest)
		return
	}
	posts, err := postService.GetAllPostsByUserID(uint(userId))
	if err != nil {
		http.Error(writer, "error getting post", http.StatusInternalServerError)
		return
	}
	utils.JsonResponse(writer, posts)

}
