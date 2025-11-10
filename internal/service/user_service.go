package service

import (
	"context"
	"time"
	"userdata-api/internal/models"
	"userdata-api/internal/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (userService *UserService) CreateUser(ctx context.Context, name string, dob time.Time) (int32, error) {
	return userService.repo.CreateUser(ctx, name, dob)
}

func (userService *UserService) UpdateUser(ctx context.Context, id int32, name string, dob time.Time) (*models.User, error) {
	user, err := userService.repo.UpdateUser(ctx, id, name, dob)
	if err != nil {
		return nil, err
	}
	// Calcuate age
	return &models.User{
		Userid: user.ID,
		Name:   user.Name,
		Dob:    user.Dob.Format("2006-01-02"),
		Age:    calculateAge(user.Dob),
	}, err
}

func (userService *UserService) DeleteUser(ctx context.Context, id int32) error {
	return userService.repo.DeleteUser(ctx, id)
}

func (userService *UserService) ListUsers(ctx context.Context, limit int32, offset int32) ([]models.User, error) {
	users, err := userService.repo.ListUsers(ctx, limit, offset)
	if err != nil {
		return nil, err
	}

	// calculating age for all users
	var usersWithDob []models.User

	for _, user := range users {
		modelUser := models.User{
			Userid: user.ID,
			Name:   user.Name,
			Dob:    user.Dob.Format("2006-01-02"),
			Age:    calculateAge(user.Dob),
		}
		usersWithDob = append(usersWithDob, modelUser)
	}

	return usersWithDob, nil
}

func (userService *UserService) GetUserByID(ctx context.Context, id int32) (*models.User, error) {
	user, err := userService.repo.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}

	// Calcuate age
	return &models.User{
		Userid: user.ID,
		Name:   user.Name,
		Dob:    user.Dob.Format("2006-01-02"),
		Age:    calculateAge(user.Dob),
	}, err
}

func calculateAge(dob time.Time) int {
	now := time.Now()
	age := now.Year() - dob.Year()

	// If birthday hasn’t occurred yet this year, subtract one
	if now.YearDay() < dob.YearDay() {
		age--
	}
	return age
}
