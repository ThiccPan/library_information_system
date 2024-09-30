package controller

import (
	"log/slog"
	"net/http"
	"strconv"

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
		jwt:       jwt,
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

	token, err := uc.jwt.GenerateToken(res.Id, res.Email, res.Name, res.RoleId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"message": "failed to login user",
			"error":   err.Error(),
		})
	}
	res.Token = token

	return c.JSON(http.StatusOK, map[string]any{
		"message": "user login successfully",
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

func (uc *UserController) GetProfileController(c echo.Context) error {
	id := c.Get("user").(*config.JwtCustomClaims).Id

	res, err := uc.usecase.GetById(c.Request().Context(), uint(id))
	if err != nil {
		return c.JSON(res.Status, map[string]any{
			"message": "failed to login user",
			"error":   err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]any{
		"message": "successfully get user profile",
		"data":    res,
	})
}

func (uc *UserController) GetByIdController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "failed to get user",
			"error":   err.Error(),
		})
	}

	res, err := uc.usecase.GetById(c.Request().Context(), uint(id))
	if err != nil {
		return c.JSON(res.Status, map[string]any{
			"message": "failed to login user",
			"error":   err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]any{
		"message": "successfully get user",
		"data":    res,
	})
}

func (uc *UserController) UpdateController(c echo.Context) error {
	request := &model.UpdateUserRequest{}
	id := c.Get("user").(*config.JwtCustomClaims).Id
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "failed to update user",
			"error":   err.Error(),
		})
	}
	request.Id = id
	if err := uc.validator.Struct(request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "failed to update user",
			"error":   err.Error(),
		})
	}

	res, err := uc.usecase.UpdateUser(c.Request().Context(), request, false)
	if err != nil {
		return c.JSON(res.Status, map[string]any{
			"message": "failed to update user",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "successfully update user",
		"data":    res,
	})
}

func (uc *UserController) UpdateProfileController(c echo.Context) error {
	id := c.Get("user").(*config.JwtCustomClaims).Id
	file, err := c.FormFile("image")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "failed to update user",
			"error":   err.Error(),
		})
	}
	request := &model.UpdateUserRequest{Id: id, Profile: file}
	res, err := uc.usecase.UpdateUser(c.Request().Context(), request, true)
	if err != nil {
		return c.JSON(res.Status, map[string]any{
			"message": "failed to update user",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "successfully update user profile picture",
		"data":    res,
	})
}
