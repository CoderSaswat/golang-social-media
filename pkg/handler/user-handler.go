package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/mashingan/smapping"
	"net/http"
	"social-media/pkg/config"
	"social-media/pkg/dto"
	"social-media/pkg/model"
	_ "social-media/pkg/model"
	"social-media/pkg/service"
	service_impl "social-media/pkg/service-impl"
	"social-media/pkg/utils"
	"strconv"
)

var userService service.UserService
var productService service.ProductService

func init() {
	db := config.GetDb()
	userService = service_impl.NewUserServiceImpl(db)
	productService = service_impl.NewProductServiceImpl(db)
}

type UserHandler struct {
}

func (h *UserHandler) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = userService.CreateUser(&user)
	//productService.GetProductById()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *UserHandler) GetAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	users, err := userService.GetAllUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(users)
}

func (h *UserHandler) GetUserByIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]
	id, _ := strconv.Atoi(userID)

	// Convert userID to uint or handle error if conversion fails
	// ...

	user, err := userService.GetUserByID(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(user)
}
func (h *UserHandler) GetUserByIDHandlerV2(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]
	id, _ := strconv.Atoi(userID)

	// Convert userID to uint or handle error if conversion fails
	// ...

	user, err := userService.GetUserByID(uint(id))
	userDto := dto.UserDto{}
	err = smapping.FillStruct(&userDto, smapping.MapFields(user))
	if err != nil {
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(userDto)
}

func (h *UserHandler) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]
	id, _ := strconv.Atoi(userID)

	// Convert userID to uint or handle error if conversion fails
	// ...

	var updatedUser model.User
	err := json.NewDecoder(r.Body).Decode(&updatedUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = userService.UpdateUser(uint(id), &updatedUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *UserHandler) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]
	id, _ := strconv.Atoi(userID)

	// Convert userID to uint or handle error if conversion fails
	// ...

	err := userService.DeleteUser(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *UserHandler) FollowOrUnfollwHandler(writer http.ResponseWriter, request *http.Request) {
	var followInput dto.FollowInput
	err := json.NewDecoder(request.Body).Decode(&followInput)
	if err != nil {
		utils.ErrorResponse(writer, nil, http.StatusBadRequest, err)
		return
	}
	err = userService.FollowOrUnfollow(followInput)
	if err != nil {
		utils.ErrorResponse(writer, nil, http.StatusBadRequest, err)
		return
	}
	utils.SuccessResponse(writer, nil, http.StatusOK, "follow or unfollow action success")
}

func (h *UserHandler) GetFollowingsAndFollowersInfo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]
	id, _ := strconv.Atoi(userID)
	followingOutput, err := userService.GetFollowingsAndFollowersInfo(uint(id))
	if err != nil {
		utils.ErrorResponse(w, nil, 400, err)
		return
	}
	utils.SuccessResponse(w, followingOutput, http.StatusOK, "followers and followings info successfully retrieved")
}

func (h *UserHandler) CreateUserHandlerV2(w http.ResponseWriter, r *http.Request) {
	var userDto dto.UserDto
	err := json.NewDecoder(r.Body).Decode(&userDto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = userService.CreateUserV2(&userDto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
