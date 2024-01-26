package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"social-media/pkg/config"
	"social-media/pkg/model"
	"social-media/pkg/service"
	service_impl "social-media/pkg/service-impl"
	"social-media/pkg/utils"
	"strconv"
)

type CommentHandler struct {
}

func init() {
	db := config.GetDb()
	commentService = service_impl.NewCommentServiceImpl(db)
}

func NewCommentHandler() *CommentHandler {
	return &CommentHandler{}
}

var commentService service.CommentService

func (h *CommentHandler) CreateComment(w http.ResponseWriter, r *http.Request) {
	var comment model.Comment
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&comment); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if err := commentService.CreateComment(&comment); err != nil {
		http.Error(w, "Error creating comment", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *CommentHandler) GetCommentByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	commentID, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid comment ID", http.StatusBadRequest)
		return
	}

	comment, err := commentService.GetCommentByID(uint(commentID))
	if err != nil {
		http.Error(w, "Comment not found", http.StatusNotFound)
		return
	}

	utils.JsonResponse(w, comment)
}

func (h *CommentHandler) GetAllComments(w http.ResponseWriter, r *http.Request) {
	comments, err := commentService.GetAllComments()
	if err != nil {
		http.Error(w, "Error fetching comments", http.StatusInternalServerError)
		return
	}

	utils.JsonResponse(w, comments)
}

func (h *CommentHandler) UpdateComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	commentID, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid comment ID", http.StatusBadRequest)
		return
	}

	var comment model.Comment
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&comment); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	comment.ID = uint(commentID)
	if err := commentService.UpdateComment(&comment); err != nil {
		http.Error(w, "Error updating comment", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *CommentHandler) DeleteComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	commentID, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid comment ID", http.StatusBadRequest)
		return
	}

	if err := commentService.DeleteComment(uint(commentID)); err != nil {
		http.Error(w, "Error deleting comment", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *CommentHandler) GetCommentsByPost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	postID, err := strconv.ParseUint(vars["postId"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid post ID", http.StatusBadRequest)
		return
	}

	comments, err := commentService.GetCommentsByPost(uint(postID))
	if err != nil {
		http.Error(w, "Error fetching comments", http.StatusInternalServerError)
		return
	}

	utils.JsonResponse(w, comments)
}
