package usecase

import (
	"context"
	"errors"
	"log/slog"
	"net/http"

	"github.com/thiccpan/library_information_system/internal/entity"
	"github.com/thiccpan/library_information_system/internal/model"
	"github.com/thiccpan/library_information_system/internal/repository"
	"gorm.io/gorm"
)

type UserUsecase struct {
	Db         *gorm.DB
	Repository *repository.UserRepository
}

func NewUserUsecase(Db *gorm.DB, Repository *repository.UserRepository) *UserUsecase {
	return &UserUsecase{
		Db:         Db,
		Repository: Repository,
	}
}

func (uu *UserUsecase) Register(ctx context.Context, request *model.RegisterUserRequest) (*model.UserResponse, error) {
	tx := uu.Db.WithContext(ctx).Begin()
	defer func() {
		if r := recover(); r != nil {
			slog.Info("err exit", r)
			tx.Rollback()
		}
	}()

	// add to data store
	user := &entity.User{
		Email:    request.Email,
		Name:     request.Name,
		Password: request.Password,
		Role_id:  entity.USER.Id,
	}

	if err := uu.Repository.Add(tx, user); err != nil {
		tx.Rollback()
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			error_val := errors.New("user email has already been registered")
			slog.Error("duplicate user data", "error", err.Error())
			return &model.UserResponse{Status: http.StatusBadRequest, Error: error_val}, error_val
		}
		slog.Error("failed to create new user", "error", err.Error())
		return &model.UserResponse{Status: http.StatusInternalServerError, Error: err}, err
	}

	// commit transaction
	if err := tx.Commit().Error; err != nil {
		slog.Error("failed to create new user", "error", err.Error())
		return nil, err
	}

	// send result
	return model.UserToResponse(user), nil
}

func (uu *UserUsecase) Login(ctx context.Context, request *model.LoginUserRequest) (*model.UserResponse, error) {
	tx := uu.Db.WithContext(ctx).Begin()
	if r := recover(); r != nil {
		slog.Info("err exit", r)
		tx.Rollback()
	}

	user := &entity.User{
		Email:    request.Email,
		Password: request.Password,
	}

	if err := uu.Repository.GetByEmail(tx, user); err != nil {
		tx.Rollback()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			error_val := errors.New("user not found in record")
			slog.Error("user not found in record", "error", err.Error())
			return &model.UserResponse{Status: http.StatusBadRequest, Error: error_val}, error_val
		}
		slog.Error("failed to find user", "error", err.Error())
		return &model.UserResponse{Status: http.StatusInternalServerError, Error: err}, err
	}

	if user.Password != request.Password {
		error_val := errors.New("invalid credentials")
		slog.Error("invalid credentials")
		return &model.UserResponse{Status: http.StatusUnauthorized, Error: error_val}, error_val
	}

	// commit transaction
	if err := tx.Commit().Error; err != nil {
		slog.Error("failed to create new user", "error", err.Error())
		return nil, err
	}

	return model.UserToResponse(user), nil
}

func (uu *UserUsecase) ShowAllUsers(ctx context.Context) (*model.UsersResponse, error) {
	users, err := uu.Repository.GetAll(uu.Db)
	if err != nil {
		slog.Error("failed to get users", "error", err.Error())
		return &model.UsersResponse{Status: http.StatusInternalServerError, Error: err}, err
	}
	return &model.UsersResponse{
		Status: http.StatusInternalServerError,
		Error:  nil,
		Users: users,
	}, nil
}
