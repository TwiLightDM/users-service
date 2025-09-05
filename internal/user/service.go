package user

import (
	"github.com/google/uuid"
)

type Service interface {
	CreateUser(user User) (User, error)
	ReadAllUsers() ([]User, error)
	ReadUserById(id string) (User, error)
	UpdateUser(id string, user User) (User, error)
	DeleteUser(id string) error
}

type service struct {
	repo Repository
}

func NewUserService(repo Repository) Service {
	return &service{repo: repo}
}

func (s service) CreateUser(user User) (User, error) {
	user.Id = uuid.NewString()

	err := s.repo.CreateUser(user)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func (s service) ReadAllUsers() ([]User, error) {
	users, err := s.repo.ReadAllUsers()
	if err != nil {
		return []User{}, err
	}

	return users, nil
}

func (s service) ReadUserById(id string) (User, error) {
	user, err := s.repo.ReadUserById(id)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func (s service) UpdateUser(id string, user User) (User, error) {
	user, err := s.repo.UpdateUser(id, user)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func (s service) DeleteUser(id string) error {
	err := s.repo.DeleteUser(id)
	if err != nil {
		return err
	}

	return nil
}
