package routes

import (
	"github.com/gorilla/mux"
	"github.com/jmdefelippe/go-gorm-restapi/pkg/controllers"
)

func PublicRoutes(r *mux.Router) {
	r.HandleFunc("/", controllers.HomeHandler)

	// tasks routes
	r.HandleFunc("/tasks", controllers.GetTasksHandler).Methods("GET")
	r.HandleFunc("/tasks/{id}", controllers.GetTaskHandler).Methods("GET")
	r.HandleFunc("/tasks", controllers.CreateTaskHandler).Methods("POST")
	r.HandleFunc("/tasks/{id}", controllers.DeleteTaskHandler).Methods("DELETE")

	// users routes
	r.HandleFunc("/users", controllers.GetUsersHandler).Methods("GET")
	r.HandleFunc("/users/{id}", controllers.GetUserHandler).Methods("GET")
	r.HandleFunc("/users", controllers.PostUserHandler).Methods("POST")
	r.HandleFunc("/users/{id}", controllers.DeleteUserHandler).Methods("DELETE")
}
