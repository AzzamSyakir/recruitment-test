package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"recruitment-test/application/entities"
	"recruitment-test/application/repositories"
	"recruitment-test/application/responses"
	"recruitment-test/application/service"

	"github.com/gorilla/mux"
)

type TaskController struct {
	TaskService service.TaskService
}

func NewTaskController(taskService service.TaskService) *TaskController {
	return &TaskController{
		TaskService: taskService,
	}
}
func TaskInitialize(db *sql.DB) *TaskController {
	// Inisialisasi Repository dan Service
	taskRepository := repositories.NewTaskRepository(db)
	taskService := service.NewTaskService(*taskRepository)

	// Inisialisasi Controller
	taskController := NewTaskController(*taskService)

	return taskController
}

// CreateTaskController adalah fungsi untuk menangani pembuatan tugas
func (tc *TaskController) CreateTaskController(w http.ResponseWriter, r *http.Request) {
	var task entities.Task
	tokenString := r.Header.Get("Authorization")

	// Membersihkan token dari string "Bearer "
	tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		responses.ErrorResponse(w, "Failed to read task data from the request", http.StatusBadRequest)
		return
	}

	taskID, err := tc.TaskService.CreateTask(task, tokenString)
	if err != nil {
		errorMessage := fmt.Sprintf("Failed to create task: %v", err)
		responses.ErrorResponse(w, errorMessage, http.StatusInternalServerError)
		return
	}

	// Membuat objek data tugas untuk dikirim dalam respons
	taskData := struct {
		ID          string    `json:"id"`
		Title       string    `json:"title"`
		Description string    `json:"description"`
		Status      string    `json:"status"`
		DueDate     time.Time `json:"due_date"`
	}{
		ID:          taskID,
		Title:       task.Title,
		Description: task.Description,
		Status:      string(task.Status),
		DueDate:     task.DueDate,
	}

	responses.SuccessResponse(w, "Success", taskData, http.StatusCreated)
}

// FetchTaskController adalah fungsi untuk menangani pengambilan daftar tugas
func (tc *TaskController) FetchTaskController(w http.ResponseWriter, r *http.Request) {
	tasksData, err := tc.TaskService.FetchTasks()
	if err != nil {
		errorMessage := fmt.Sprintf("Failed to get tasks: %v", err)
		responses.ErrorResponse(w, errorMessage, http.StatusInternalServerError)
		return
	}

	// Membuat objek data tugas untuk dikirim dalam respons
	var responseData []entities.Task

	for _, task := range tasksData {
		taskData := entities.Task{
			ID:          task.ID,
			Title:       task.Title,
			Description: task.Description,
			Status:      task.Status,
			DueDate:     task.DueDate,
		}

		responseData = append(responseData, taskData)
	}

	// Mengembalikan data tugas sebagai JSON
	w.Header().Set("Content-Type", "application/json")
	responses.SuccessResponse(w, "Success", responseData, http.StatusOK)
}

// GetTaskController adalah fungsi untuk menangani pengambilan tugas berdasarkan ID
func (tc *TaskController) GetTaskController(w http.ResponseWriter, r *http.Request) {
}

// UpdateTaskController adalah fungsi untuk menangani pembaruan tugas
func (tc *TaskController) UpdateTaskController(w http.ResponseWriter, r *http.Request) {
	var task entities.Task

	// Decode data tugas dari permintaan
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		responses.ErrorResponse(w, "Failed to read task data from the request", http.StatusBadRequest)
		return
	}

	// Mendapatkan parameter ID dari URL
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		responses.ErrorResponse(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	// Update tugas di service layer
	err = tc.TaskService.UpdateTask(id, task)
	if err != nil {
		errorMessage := fmt.Sprintf("Failed to update task: %v", err)
		responses.ErrorResponse(w, errorMessage, http.StatusInternalServerError)
		return
	}

	// Membuat objek data tugas untuk dikirim dalam respons
	taskData := struct {
		ID          int       `json:"id"`
		Title       string    `json:"title"`
		Description string    `json:"description"`
		Status      string    `json:"status"`
		DueDate     time.Time `json:"due_date"`
	}{
		ID:          id,
		Title:       task.Title,
		Description: task.Description,
		Status:      string(task.Status),
		DueDate:     task.DueDate,
	}

	// Mengembalikan data tugas yang telah diupdate sebagai JSON
	w.Header().Set("Content-Type", "application/json")
	responses.SuccessResponse(w, "Success", taskData, http.StatusOK)
}

// DeleteTaskController adalah fungsi untuk menangani penghapusan tugas
func (tc *TaskController) DeleteTaskController(w http.ResponseWriter, r *http.Request) {
	// ... implementasi endpoint DeleteTaskController
}

// LoginUserController adalah fungsi untuk menangani operasi login pengguna
func (tc *TaskController) LoginUserController(w http.ResponseWriter, r *http.Request) {
	// ... implementasi endpoint LoginUserController
}
