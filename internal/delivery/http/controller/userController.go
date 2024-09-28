package controller

import (
	"log/slog"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/thiccpan/library_information_system/internal/model"
	"github.com/thiccpan/library_information_system/internal/usecase"
)

type UserController struct {
	usecase   *usecase.UserUsecase
	validator *validator.Validate
}

func NewUserController(u *usecase.UserUsecase, validator *validator.Validate) *UserController {
	return &UserController{
		usecase:   u,
		validator: validator,
	}
}

func (uc *UserController) RegisterController(c echo.Context) error {
	req := new(model.RegisterUserRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "failed to register user",
			"error":   err.Error(),
		})
	}

	if err := uc.validator.Struct(req); err != nil {
		slog.Error(err.Error())
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "failed to register user",
			"error":   err.Error(),
		})
	}

	res, err := uc.usecase.Register(c.Request().Context(), req)
	if err != nil {
		return c.JSON(res.Status, map[string]any{
			"message": "failed to register user",
			"error":   res.Error.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]any{
		"message": "user registered successfully",
		"data":    res,
	})
}
