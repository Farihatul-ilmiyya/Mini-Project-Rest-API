package repository

import (
	"mini-project/package/models"
	"fmt"

	"gorm.io/gorm"
)

type IUserRepo interface {
	GetAllUser() ([]models.UserCostum, error)
	GetUserByUsername(username string) (models.User, error)
	InsertUser(user models.User) error
	GetUserById(id int) (models.User, error)
}

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) IUserRepo {
	return &UserRepo{db}
}

func (r UserRepo) GetAllUser() ([]models.UserCostum, error) {
	users := []models.UserCostum{}
	err := r.db.Find(&users).Error // SELECT * FROM users;
	if err != nil {
		fmt.Println("error while GetAllUser", err)
	}
	return users, err
}

func (repo UserRepo) GetUserByUsername(username string) (models.User, error) {
	var user models.User
	err := repo.db.Where("username = ?", username).First(&user).Error // SELECT * FROM users WHERE username = ? LIMIT 1;
	if err != nil {
		fmt.Println("error while GetUserByUsernameAndPassword", err)
	}
	return user, err
}

func (repo UserRepo) InsertUser(user models.User) error {
	err := repo.db.Create(&user).Error
	if err != nil {
		fmt.Println("error while InsertUser", err)
	}
	return err
}

func (r UserRepo) GetUserById(id int) (models.User, error) {
	var user models.User
	err := r.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		fmt.Println("error while GetUserById", err)
	}
	return user, err
}