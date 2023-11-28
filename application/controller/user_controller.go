package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"clean-golang/application/responses"
	"clean-golang/application/service"

	"github.com/gorilla/mux"
)

type UserController struct {
	userService service.UserService
}

func NewInstance(us service.UserService) *UserController {
	return &UserController{userService: us}
}

func (uc *UserController) CreateUserController(w http.ResponseWriter, r *http.Request) {
	var user struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		responses.ErrorResponse(w, "Failed to read user data from the request", http.StatusBadRequest)
		return
	}

	userID, err := uc.userService.CreateUser(user.Name, user.Email, user.Password)
	if err != nil {
		errorMessage := fmt.Sprintf("Failed to create user: %v", err)
		responses.ErrorResponse(w, errorMessage, http.StatusInternalServerError)
		return
	}

	currentTime := time.Now()

	// Membuat objek data pengguna untuk dikirim dalam respons
	userData := struct {
		ID        int64     `json:"id"`
		Name      string    `json:"name"`
		Email     string    `json:"email"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}{
		ID:        userID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
	}

	responses.SuccessResponse(w, "Success", userData, http.StatusCreated)
}
func (uc *UserController) FetchUserController(w http.ResponseWriter, r *http.Request) {
	userData, err := uc.userService.FetchUser()
	if err != nil {
		errorMessage := fmt.Sprintf("Failed to get user: %v", err)
		responses.ErrorResponse(w, errorMessage, http.StatusInternalServerError)
		return
	}
	// Membuat objek data pengguna untuk dikirim dalam respons
	responseData := struct {
		Username  string `json:"username"`
		Email     string `json:"email"`
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at"`
	}{
		Username:  userData["username"],
		Email:     userData["email"],
		CreatedAt: userData["created_at"],
		UpdatedAt: userData["updated_at"],
	}
	// Mengembalikan data pengguna sebagai JSON
	w.Header().Set("Content-Type", "application/json")
	responses.SuccessResponse(w, "Success", responseData, http.StatusOK)
}
func (uc *UserController) GetUserController(w http.ResponseWriter, r *http.Request) {
	// Mendapatkan parameter id
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		responses.ErrorResponse(w, "id harus disertakan", http.StatusBadRequest)
		return
	}
	userData, err := uc.userService.GetUser(id)
	if err != nil {
		errorMessage := fmt.Sprintf("Failed to get user: %v", err)
		responses.ErrorResponse(w, errorMessage, http.StatusInternalServerError)
		return
	}

	// Membuat objek data pengguna untuk dikirim dalam respons
	responseData := struct {
		Username  string `json:"username"`
		Email     string `json:"email"`
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at"`
	}{
		Username:  userData["username"],
		Email:     userData["email"],
		CreatedAt: userData["created_at"],
		UpdatedAt: userData["updated_at"],
	}
	// Mengembalikan data pengguna sebagai JSON
	w.Header().Set("Content-Type", "application/json")
	responses.SuccessResponse(w, "Success", responseData, http.StatusOK)
}

func (uc *UserController) UpdateUserController(w http.ResponseWriter, r *http.Request) {
	// Mendapatkan parameter id
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		responses.ErrorResponse(w, "id harus disertakan", http.StatusBadRequest)
		return
	}

	var user struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		responses.ErrorResponse(w, "Failed to read user data from the request", http.StatusBadRequest)
		return
	}

	// Update user in the service layer
	err = uc.userService.UpdateUser(id, user.Name, user.Email, user.Password)
	if err != nil {
		errorMessage := fmt.Sprintf("Failed to update user: %v", err)
		responses.ErrorResponse(w, errorMessage, http.StatusInternalServerError)
		return
	}

	currentTime := time.Now()

	// Membuat objek data pengguna untuk dikirim dalam respons
	userData := struct {
		ID        int64     `json:"id"`
		Name      string    `json:"name"`
		Email     string    `json:"email"`
		UpdatedAt time.Time `json:"updated_at"`
	}{
		ID:        int64(id),
		Name:      user.Name,
		Email:     user.Email,
		UpdatedAt: currentTime,
	}

	responses.SuccessResponse(w, "Success", userData, http.StatusCreated)
}
func (uc *UserController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	// Mendapatkan parameter id
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		responses.ErrorResponse(w, "id harus disertakan", http.StatusBadRequest)
		return
	}
	// delete user in the service layer
	err = uc.userService.DeleteUser(id)
	if err != nil {
		errorMessage := fmt.Sprintf("Failed to Delete user: %v", err)
		responses.ErrorResponse(w, errorMessage, http.StatusInternalServerError)
		return
	}

	responses.OtherResponses(w, "Success delete user", http.StatusCreated)
}
