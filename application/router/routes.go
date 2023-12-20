package routes

import (
	"database/sql"
	"net/http"
	"recruitment-test/application/controller"
	"recruitment-test/application/middleware"
	"recruitment-test/config"

	"github.com/gorilla/mux"
)

func SetupRoutes(db *sql.DB) *mux.Router {
	userController := controller.UserInitialize(db)
	taskController := controller.TaskInitialize(db)
	// Create a new router
	router := mux.NewRouter()

	// Protected routes
	protectedRoutes := router.PathPrefix("/").Subrouter()
	protectedRoutes.Use(middleware.AuthMiddleware)

	// Authentication routes
	router.HandleFunc("/users", userController.CreateUserController).Methods("POST")
	router.HandleFunc("/users/login", userController.LoginUser).Methods("POST")
	router.HandleFunc("/users/logout", userController.LogoutUser).Methods("POST")

	// User routes
	protectedRoutes.HandleFunc("/users", userController.FetchUserController).Methods("GET")
	protectedRoutes.HandleFunc("/users/{id}", userController.GetUserController).Methods("GET")
	protectedRoutes.HandleFunc("/users/{id}", userController.UpdateUserController).Methods("PUT")
	protectedRoutes.HandleFunc("/users/{id}", userController.DeleteUser).Methods("DELETE")

	// Task routes
	protectedRoutes.HandleFunc("/tasks", taskController.CreateTaskController).Methods("POST")
	protectedRoutes.HandleFunc("/tasks", taskController.GetTaskController).Methods("GET")
	protectedRoutes.HandleFunc("/tasks/{id}", taskController.GetTaskController).Methods("GET")
	protectedRoutes.HandleFunc("/tasks/{id}", taskController.UpdateTaskController).Methods("PUT")
	protectedRoutes.HandleFunc("/tasks/{id}", taskController.DeleteTaskController).Methods("DELETE")

	return router
}

func RunServer() {
	db := config.InitDB()
	router := SetupRoutes(db)

	// Mulai server HTTP dengan router yang telah dikonfigurasi
	http.Handle("/", router)
	http.ListenAndServe(":9000", nil)
}
