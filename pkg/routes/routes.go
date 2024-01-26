package routes

import "github.com/gorilla/mux"

func InitializeRoutes() *mux.Router {
	router := mux.NewRouter().UseEncodedPath()
	SetupUserRoutes(router)
	SetupPostRoutes(router)
	SetupCommentsRoutes(router)
	SetupLikeRoutes(router)
	SetupNotificationRoutes(router)
	return router
}
