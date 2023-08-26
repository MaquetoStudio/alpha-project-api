package services

import (
	"alpha-project-api/models"
	"alpha-project-api/pkg"
	"alpha-project-api/repositories"
)

type UserService struct {
	Repo *repositories.UserRepo
}

func NewUserService() *UserService {
	return &UserService{
		Repo: repositories.NewUserRepo(),
	}
}

func (s *UserService) CreateUser(user models.User) (interface{}, error) {
	var err error
	user.Password, err = pkg.EncodePassword(user.Password)
	if err != nil {
		return nil, err
	}

	created, err := s.Repo.Insert(user)

	return created.ToDTO(), nil
}

func (s *UserService) GetUserByID() (interface{}, error) {

	var result interface{}

	return result, nil
}
