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

func (uu *UserUsecase) Login() {}

func (uu *UserUsecase) ShowAllUsers() {}
