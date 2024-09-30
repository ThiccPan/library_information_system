package usecase

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/thiccpan/library_information_system/internal/entity"
	"github.com/thiccpan/library_information_system/internal/model"
	"github.com/thiccpan/library_information_system/internal/repository"
	"gorm.io/gorm"
)

const profileLocation = "resource/user_profile/"

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

	res := model.UserToResponse(user)
	res.RoleId = user.Role_id
	return res, nil
}

func (uu *UserUsecase) GetById(ctx context.Context, id uint) (*model.UserResponse, error) {
	tx := uu.Db.WithContext(ctx).Begin()
	if r := recover(); r != nil {
		slog.Info("err exit", r)
		tx.Rollback()
	}
	user := &entity.User{Id: id}
	if err := uu.Repository.GetById(tx, user); err != nil {
		tx.Rollback()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			error_val := errors.New("user not found in record")
			slog.Error("user not found in record", "error", err.Error())
			return &model.UserResponse{Status: http.StatusBadRequest, Error: error_val}, error_val
		}
		slog.Error("failed to find user", "error", err.Error())
		return &model.UserResponse{Status: http.StatusInternalServerError, Error: err}, err
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
		Status: http.StatusOK,
		Error:  nil,
		Users:  users,
	}, nil
}

func (uu *UserUsecase) UpdateUser(ctx context.Context, request *model.UpdateUserRequest, withProfile bool) (*model.UserResponse, error) {
	tx := uu.Db.WithContext(ctx).Begin()
	if r := recover(); r != nil {
		slog.Info("err exit", r)
		tx.Rollback()
	}

	// add to data store
	user := &entity.User{
		Id:       request.Id,
		Email:    request.Email,
		Password: request.Password,
		Name:     request.Name,
	}
	if request.Profile != nil {
		res := fmt.Sprintf("%d-img.jpeg", request.Id)
		user.ProfileImg = res
	}

	if err := uu.Repository.UpdateById(tx, user); err != nil {
		tx.Rollback()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			error_val := errors.New("user not found in record")
			slog.Error("user not found in record", "error", err.Error())
			return &model.UserResponse{Status: http.StatusBadRequest, Error: error_val}, error_val
		}
		slog.Error("failed to update user", "error", err.Error())
		return &model.UserResponse{Status: http.StatusInternalServerError, Error: err}, err
	}

	if request.Profile != nil {
		_, err := moveFile(fmt.Sprint(profileLocation, user.ProfileImg), request.Profile)
		if err != nil {
			tx.Rollback()
			return &model.UserResponse{Status: http.StatusInternalServerError, Error: err}, err
		}
	}

	response := model.UserToResponse(user)
	response.Status = http.StatusOK

	// commit transaction
	if err := tx.Commit().Error; err != nil {
		slog.Error("failed to create new user", "error", err.Error())
		return nil, err
	}

	return response, nil
}
