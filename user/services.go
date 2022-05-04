package user

import (
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(input RegisterUserInput) (User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}
func (s *service) RegisterUser(input RegisterUserInput) (User, error) {
	user := User{}

	user.FullName = input.FullName
	user.Email = input.Email
	user.Occupation = input.Occupation
	password, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}
	user.Password = string(password)
	user.Role = "user"

	status, err := s.repository.Save(user)
	if err != nil {
		return user, err
	}
	return status, nil

}
