package user

import (
	"gorm.io/gorm"
)

type Repository interface {
	CreateUser(user User) error
	ReadAllUsers() ([]User, error)
	ReadUserById(id string) (User, error)
	UpdateUser(id string, user User) (User, error)
	DeleteUser(id string) error
}

type repository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r repository) CreateUser(user User) error {
	return r.db.Create(&user).Error
}

func (r repository) ReadAllUsers() ([]User, error) {
	var users []User
	err := r.db.Find(&users).Error
	return users, err
}

func (r repository) ReadUserById(id string) (User, error) {
	var user User
	err := r.db.Where("id = ?", id).First(&user).Error
	return user, err
}

func (r repository) UpdateUser(id string, user User) (User, error) {
	var updatedUser User

	err := r.db.Model(&User{}).
		Where("id = ?", id).
		Updates(user).
		Scan(&updatedUser).Error

	if err != nil {
		return User{}, err
	}

	return updatedUser, nil
}

func (r repository) DeleteUser(id string) error {
	return r.db.Delete(&User{}, "id = ?", id).Error
}
