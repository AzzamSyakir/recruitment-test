package routes

import (
	"database/sql"
	"net/http"

	"clean-golang/application/controller"
	"clean-golang/application/repository"
	"clean-golang/application/service"
	"clean-golang/config"

	"github.com/gorilla/mux"
)

func SetupRoutes(db *sql.DB) *mux.Router {
	userRepo := repository.NewInstance(db)
	userService := service.NewInstance(*userRepo)
	userController := controller.NewInstance(*userService)
	router := mux.NewRouter()

	// Routes for user handling
	router.HandleFunc("/users", userController.CreateUserController).Methods("POST")
	router.HandleFunc("/users", userController.FetchUserController).Methods("GET")
	router.HandleFunc("/users/{id}", userController.GetUserController).Methods("GET")
	router.HandleFunc("/users/{id}", userController.UpdateUserController).Methods("PUT")
	router.HandleFunc("/users/{id}", userController.DeleteUser).Methods("DELETE")

	// Add more routes as needed

	return router
}
func RunServer() {
	db := config.InitDB()
	router := SetupRoutes(db)

	// Mulai server HTTP dengan router yang telah dikonfigurasi
	http.Handle("/", router)
	http.ListenAndServe(":9000", nil)
}
