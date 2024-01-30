package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"social-media/pkg/config"
	"social-media/pkg/model"
)

type RedisHandler struct {
}

func (h RedisHandler) PostData(writer http.ResponseWriter, request *http.Request) {
	var user model.User
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&user)
	if err != nil {
		return
	}
	userJSON, err := json.Marshal(user)
	if err != nil {
		fmt.Println("Error marshalling order to JSON:", err)
		return
	}
	key := "users"
	err = config.GetRedisClient().HSet(key, "user-2", userJSON).Err()
	if err != nil {
		fmt.Println("Error storing order details in Redis:", err)
		return
	}

}

//func storeUserInRedis(user User) error {
//	// Convert the user object to JSON
//	userJSON, err := json.Marshal(user)
//	if err != nil {
//		return err
//	}
//
//	// Store the JSON in Redis with a key (assuming user.ID is unique)
//	key := fmt.Sprintf("user:%d", user.ID)
//	err = config.GetRedisClient().Set(rdb.Context(), key, userJSON, 0).Err()
//	if err != nil {
//		return err
//	}
//
//	return nil
//}
