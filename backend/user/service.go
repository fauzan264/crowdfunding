package user

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(input RegisterUserInput) (User, error)
	Login(input LoginInput) (User, error)
	IsEmailAvailable(input CheckEmailInput) (bool, error)
	SaveAvatar(ID uuid.UUID, fileLocation string) (User, error)
	GetUserByID(ID uuid.UUID) (User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func(s *service) RegisterUser(input RegisterUserInput) (User, error) {
	user := User{}
	user.ID = uuid.New()
	user.Name = input.Name
	user.Email = input.Email
	user.Occupation = input.Occupation
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}

	user.Password = string(passwordHash)
	user.AvatarFileName = "images.jpg"
	user.Role = "user"

	userCreated, err := uuid.Parse("00000000-0000-0000-0000-000000000000")
	if err != nil {
		fmt.Printf("Error parsing UUID: %v\n", err)
		return user, err
	}

	user.CreatedBy = userCreated
	user.CreatedAt = time.Now()


	newUser, err := s.repository.Save(user)
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}

func (s *service) Login(input LoginInput) (User, error) {
	email := input.Email
	password := input.Password

	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return user, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *service) IsEmailAvailable(input CheckEmailInput) (bool, error) {
	email := input.Email

	_, err := s.repository.FindByEmail(email)
	if err != nil {
		return true, err
	}

	return false, nil
}

func (s *service) SaveAvatar(ID uuid.UUID, fileLocation string) (User, error) {
	user, err := s.repository.FindByID(ID)
	if err != nil {
		return user, err
	}

	user.AvatarFileName = fileLocation
	// user.UpdatedBy = uuid.MustParse("5a4b0db3-4ed2-42c5-b283-e9e32aadef21")
	user.UpdatedBy = user.ID
	user.UpdatedAt = time.Now()

	updatedUser, err := s.repository.Update(user)
	if err != nil {
		return updatedUser, err
	}

	return updatedUser, nil
}

func (s *service) GetUserByID(ID uuid.UUID) (User, error) {
	user, err := s.repository.FindByID(ID)
	if err != nil {
		return user, err
	}

	return user, nil
}