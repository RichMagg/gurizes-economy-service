package services

import (
	"time"

	"github.com/RichMagg/gurizes-economy-service/cmd/internal/models"
	"github.com/RichMagg/gurizes-economy-service/cmd/internal/repositories"
	"github.com/google/uuid"
)

type UserService struct {
	repository repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return UserService{repository: repo}
}

func (us *UserService) GetUsers() ([]models.User, error) {
	return us.repository.GetUsers()
}

func (us *UserService) CreateUser() (*models.User, error) {
	user := &models.User{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		LastSeen:  time.Now(),
	}

	err := us.repository.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
