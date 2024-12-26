package user

import (
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository interface {
	Save(user User) (User, error)
	FindByEmail(email string) (User, error)
	FindByID(id uuid.UUID) (User, error)
	Update(user User) (User, error)
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
}

func (r *repository) FindByID(id uuid.UUID) (User, error) {
	var user User
	result := r.db.Where("id = ?", id).Find(&user)
	
	if result.RowsAffected == 0 {
		return user, errors.New("User Not Found")
	}

	if result.Error != nil {
		return user, result.Error
	}

	return user, nil
}

func (r *repository) Update(user User) (User, error) {
	err := r.db.Save(&user).Error
	// err := r.db.Where("id = ?", user.ID).Updates(&user).Error
	
	if err != nil {
		return user, err
	}

	return user, nil
}