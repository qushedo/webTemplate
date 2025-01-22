package service

import (
	"context"
	"webTemplate/internal/domain/common/errorz"
	"webTemplate/internal/domain/dto"
	"webTemplate/internal/domain/entity"
)

type userStorage interface {
	Create(ctx context.Context, user entity.User) (*entity.User, error)
	GetByID(ctx context.Context, id string) (*entity.User, error)
	GetAll(ctx context.Context, limit, offset int) ([]entity.User, error)
	Update(ctx context.Context, user *entity.User) (*entity.User, error)
	Delete(ctx context.Context, id string) error
	GetByEmail(ctx context.Context, email string) (*entity.User, error)
}

type userService struct {
	storage userStorage
}

func NewUserService(storage userStorage) *userService {
	return &userService{storage: storage}
}

func (s *userService) Create(ctx context.Context, registerReq dto.UserRegister, code string) (*entity.User, error) {
	if _, err := s.storage.GetByEmail(ctx, registerReq.Email); err == nil {
		return nil, errorz.EmailAlreadyTaken
	}

	user := entity.User{
		Email:            registerReq.Email,
		Username:         registerReq.Username,
		VerificationCode: code,
	}
	user.SetPassword(registerReq.Password)
	return s.storage.Create(ctx, user)
}

func (s *userService) GetByEmail(ctx context.Context, email string) (*entity.User, error) {
	return s.storage.GetByEmail(ctx, email)
}

func (s *userService) GetByID(ctx context.Context, id string) (*entity.User, error) {
	return s.storage.GetByID(ctx, id)
}

func (s *userService) Update(ctx context.Context, user *entity.User) (*entity.User, error) {
	return s.storage.Update(ctx, user)
}
