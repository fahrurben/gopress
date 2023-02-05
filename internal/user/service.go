package user

import (
	"context"
	"errors"
	"golang.org/x/crypto/bcrypt"
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
	userType := ctx.Value("auth_type").(int)

	if userType != TYPE_ADMIN {
		return nil, errors.New("this request not authorized")
	}

	password := []byte(request.Password)
	// Hashing the password with the default cost of 10
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}

	user := User{
		Name:      request.Name,
		Email:     request.Email,
		Password:  string(hashedPassword),
		Type:      request.Type,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	id, err := s.repository.Save(user)

	return id, err
}
