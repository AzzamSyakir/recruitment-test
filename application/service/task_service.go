package service

import (
	"errors"
	"fmt"
	"os"
	models "recruitment-test/application/entities"
	repository "recruitment-test/application/repositories"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

type TaskService struct {
	TaskRepository repository.TaskRepository
}

func NewTaskService(taskRepository repository.TaskRepository) *TaskService {
	return &TaskService{
		TaskRepository: taskRepository,
	}
}

func (ts *TaskService) CreateTask(task models.Task, tokenString string) (string, error) {
	type DataUsers struct {
		Username string `json:"username"`
		UserId   string `json:"user_id"`
		jwt.StandardClaims
	}
	// Parse token dan dapatkan claims
	token, err := jwt.ParseWithClaims(tokenString, &DataUsers{}, func(token *jwt.Token) (interface{}, error) {
		// Verifikasi bahwa metode tanda tangan sesuai
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("metode tanda tangan tidak valid: %v", token.Header["alg"])
		}
		// Kembalikan secret key untuk memverifikasi tanda tangan
		secretKeyString := os.Getenv("SECRET_KEY")
		return []byte(secretKeyString), nil
	})
	if err != nil {
		return tokenString, fmt.Errorf("gagal mengurai token: %v", err)
	}

	// Mengekstrak klaim dari token
	claims, ok := token.Claims.(*DataUsers)
	if !ok {
		return tokenString, nil
	}
	task.Owner_id = claims.UserId
	// Validasi input untuk setiap field
	if task.Title == "" {
		return "", errors.New("title is required")
	}

	if task.Description == "" {
		return "", errors.New("description is required")
	}

	if task.DueDate.IsZero() {
		return "", errors.New("due date is required")
	}

	// Pengaturan default untuk status jika tidak diisi
	if task.Status == "" {
		task.Status = models.TaskStatus(models.NotDone)
	}

	// Validasi bisnis lainnya
	if task.DueDate.Before(time.Now()) {
		return "", errors.New("due date must be in the future")
	}
	// Generate a UUID for the tasks
	uuid := uuid.New()
	taskNewID := uuid.String()

	// Panggil repository untuk membuat tugas
	taskID, err := ts.TaskRepository.CreateTask(task, taskNewID)
	if err != nil {
		return "0", err
	}

	return taskID, nil
}

func (ts *TaskService) FetchTasks() ([]models.Task, error) {
	return ts.TaskRepository.FetchTasks()
}

func (ts *TaskService) UpdateTask(id int, updatedTask models.Task) error {
	// Validasi input untuk setiap field
	if updatedTask.Title == "" {
		return errors.New("title is required")
	}

	if updatedTask.Description == "" {
		return errors.New("description is required")
	}

	if updatedTask.DueDate.IsZero() {
		return errors.New("due date is required")
	}

	// Validasi bisnis lainnya
	if updatedTask.DueDate.Before(time.Now()) {
		return errors.New("due date must be in the future")
	}

	// Panggil repository untuk melakukan pembaruan tugas
	err := ts.TaskRepository.UpdateTask(id, updatedTask)
	if err != nil {
		return err
	}

	return nil
}

func (ts *TaskService) DeleteTask(id int) error {
	// Panggil repository untuk menghapus tugas
	err := ts.TaskRepository.DeleteTask(id)
	if err != nil {
		return err
	}

	return nil
}

func (ts *TaskService) GetTask(id int) (*models.Task, error) {
	return ts.TaskRepository.GetTask(id)
}
