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

var likeService service.LikeService

func init() {
	db := config.GetDb()
	likeService = service_impl.NewLikeServiceImpl(db)
}

type LikeHandler struct {
}

func (h *LikeHandler) CreateLike(w http.ResponseWriter, r *http.Request) {
	var like model.Like
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&like); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if err := likeService.CreateLike(&like); err != nil {
		http.Error(w, "Error creating like", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *LikeHandler) GetLikeByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	likeID, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid like ID", http.StatusBadRequest)
		return
	}

	like, err := likeService.GetLikeByID(uint(likeID))
	if err != nil {
		http.Error(w, "Like not found", http.StatusNotFound)
		return
	}

	utils.JsonResponse(w, like)
}

func (h *LikeHandler) UpdateLike(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	likeID, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid like ID", http.StatusBadRequest)
		return
	}

	var like model.Like
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&like); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	like.ID = uint(likeID)
	if err := likeService.UpdateLike(&like); err != nil {
		http.Error(w, "Error updating like", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *LikeHandler) DeleteLike(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	likeID, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid like ID", http.StatusBadRequest)
		return
	}

	if err := likeService.DeleteLike(uint(likeID)); err != nil {
		http.Error(w, "Error deleting like", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *LikeHandler) GetAllLikes(w http.ResponseWriter, r *http.Request) {
	likes, err := likeService.GetAllLikes()
	if err != nil {
		http.Error(w, "Error fetching likes", http.StatusInternalServerError)
		return
	}

	utils.JsonResponse(w, likes)
}

func (h *LikeHandler) GetLikesByPostID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	postID, err := strconv.ParseUint(vars["postId"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid post ID", http.StatusBadRequest)
		return
	}

	likes, err := likeService.GetLikesByPostID(uint(postID))
	if err != nil {
		http.Error(w, "Error fetching likes", http.StatusInternalServerError)
		return
	}

	utils.JsonResponse(w, likes)
}

func (h *LikeHandler) GetLikeCountByPostID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	postID, err := strconv.ParseUint(vars["postId"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid post ID", http.StatusBadRequest)
		return
	}

	count, err := likeService.GetLikeCountByPostID(uint(postID))
	if err != nil {
		http.Error(w, "Error fetching like count", http.StatusInternalServerError)
		return
	}

	utils.JsonResponse(w, count)
}
