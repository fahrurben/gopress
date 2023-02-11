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

func (s *ServiceImpl) Update(ctx context.Context, id int, request UpdateUserRequest) (bool, error) {
	userType := ctx.Value("auth_type").(int)
	if userType != TYPE_ADMIN {
		return false, errors.New("this request not authorized")
	}

	user, err := s.repository.FindById(id)
	if err != nil {
		return false, err
	}
	user.Name = request.Name
	_, err = s.repository.Update(*user)
	if err != nil {
		return false, err
	}

	return true, err
}

func (s ServiceImpl) Delete(ctx context.Context, id int) error {
	userType := ctx.Value("auth_type").(int)
	if userType != TYPE_ADMIN {
		return errors.New("this request not authorized")
	}

	user, err := s.repository.FindById(id)
	if err != nil {
		return err
	}
	if user == nil {
		return errors.New("entity not found")
	}

	return s.repository.DeleteById(id)
}
