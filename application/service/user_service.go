package service

import (
	"clean-golang/application/repository"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	UserRepository repository.UserRepository
}

func NewInstance(ur repository.UserRepository) *UserService {
	return &UserService{UserRepository: ur}
}

func (service *UserService) CreateUser(name, email, password string) (int64, error) {
	// Hashing password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}

	// Waktu saat ini
	currentTime := time.Now()

	// Simpan pengguna ke database dengan menggunakan data yang telah Anda validasi
	userID, err := service.UserRepository.CreateUser(name, email, string(hashedPassword), currentTime, currentTime)
	if err != nil {
		return 0, err
	}

	return userID, nil
}

func (us *UserService) FetchUser() (map[string]string, error) {
	return us.UserRepository.FetchUser()
}
func (service *UserService) UpdateUser(id int, name, email, password string) error {
	// Business logic/validation goes here
	// Hashing password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Waktu saat ini
	currentTime := time.Now()

	// Call repository to update user in the database
	err = service.UserRepository.UpdateUser(id, name, email, string(hashedPassword), currentTime)
	if err != nil {
		return err
	}

	return nil
}
func (service *UserService) DeleteUser(id int) error {
	// Business logic/validation goes here

	// Call repository to delete user in the database
	err := service.UserRepository.DeleteUser(id)
	if err != nil {
		return err
	}

	return nil
}
func (us *UserService) GetUser(id int) (map[string]string, error) {
	return us.UserRepository.GetUser(id)
}
