package user

import (
	"context"
	"time"
)

type Repository interface {
	Save(user User) (*int64, error)
	Update(user User) (bool, error)
	FindById(id int) (*User, error)
	FindAll() []User
	DeleteById(id int) error
}

type ServiceImpl struct {
	repository Repository
}

func NewService(repository Repository) *ServiceImpl {
	return &ServiceImpl{repository: repository}
}

func (s *ServiceImpl) Save(ctx context.Context, request CreateUserRequest) (*int64, error) {
	/*
		authUser := ctx.Value("auth_user").(User)

		if authUser.Type != TYPE_ADMIN {
			return nil, errors.New("this request not authorized")
		}
	*/
	user := User{
		Name:      request.Name,
		Email:     request.Email,
		Password:  request.Password,
		Type:      request.Type,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	id, err := s.repository.Save(user)

	return id, err
}
