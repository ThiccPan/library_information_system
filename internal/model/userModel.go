package model

import "github.com/thiccpan/library_information_system/internal/entity"

type JWT struct {
}

type UserResponse struct {
	Status     int    `json:"status_code"`
	Error      error  `json:"error"`
	Id         any    `json:"id"`
	Email      string `json:"email"`
	Name       string `json:"name"`
	ProfileImg string `json:"profile_img"`
	Token      string `json:"token"`
}

func UserToResponse(user *entity.User) *UserResponse {
	return &UserResponse{
		Id:         user.Id,
		Email:      user.Email,
		Name:       user.Name,
		ProfileImg: user.ProfileImg,
	}
}

type RegisterUserRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Name     string `json:"name" validate:"required,max=100"`
	Password string `json:"password" validate:"required,max=100"`
}

type UpdateUserRequest struct {
	Id       string `json:"-" validate:"required,max=100"`
	Email    string `json:"email" validate:"email,omitempty"`
	Password string `json:"password,omitempty" validate:"max=100"`
	Name     string `json:"name,omitempty" validate:"max=100"`
}

type LoginUserRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,max=100"`
}

type LogoutUserRequest struct {
	ID string `json:"id" validate:"required,max=100"`
}

type GetUserRequest struct {
	ID string `json:"id" validate:"required,max=100"`
}
