package controller

import (
	"log/slog"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/thiccpan/library_information_system/internal/config"
	"github.com/thiccpan/library_information_system/internal/model"
	"github.com/thiccpan/library_information_system/internal/usecase"
)

type UserController struct {
	usecase   *usecase.UserUsecase
	validator *validator.Validate
	jwt       *config.AuthJWT
}

func NewUserController(u *usecase.UserUsecase, validator *validator.Validate, jwt *config.AuthJWT) *UserController {
	return &UserController{
		usecase:   u,
		validator: validator,
		jwt: jwt,
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

func (uc *UserController) LoginController(c echo.Context) error {
	req := &model.LoginUserRequest{}
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "failed to login user",
			"error":   err.Error(),
		})
	}

	if err := uc.validator.Struct(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "failed to login user",
			"error":   err.Error(),
		})
	}

	res, err := uc.usecase.Login(c.Request().Context(), req)
	if err != nil {
		return c.JSON(res.Status, map[string]any{
			"message": "failed to login user",
			"error":   err.Error(),
		})
	}

	token, err := uc.jwt.GenerateToken(res.Id, res.Email, res.Name, 1)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"message": "failed to login user",
			"error":   err.Error(),
		})
	}
	res.Token = token

	return c.JSON(http.StatusOK, map[string]any{
		"message": "user registered successfully",
		"data":    res,
	})
}

func (uc *UserController) GetAllController(c echo.Context) error {
	res, err := uc.usecase.ShowAllUsers(c.Request().Context())
	if err != nil {
		return c.JSON(res.Status, map[string]any{
			"message": "failed to login user",
			"error":   err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]any{
		"message": "successfully fetch users",
		"data":    res,
	})
}