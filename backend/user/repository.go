package user

import (
	"errors"

	"gorm.io/gorm"
)

type Repository interface {
	Save(user User) (User, error)
	FindByEmail(email string) (User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(user User) (User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *repository) FindByEmail(email string) (User, error) {
	var user User
	result := r.db.Where("email = ?", email).Find(&user)
	
	if result.RowsAffected == 0 {
		return user, errors.New("User Not Found")
	}
	
	if result.Error != nil {
		return user, result.Error
	}

	return user, nil
	// if err != nil {
	// 	return user, err
	// }
	// return user, nil
}