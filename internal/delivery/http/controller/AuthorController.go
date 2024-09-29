package controller

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/thiccpan/library_information_system/internal/model"
	"github.com/thiccpan/library_information_system/internal/usecase"
)

type AuthorController struct {
	usecase *usecase.AuthorUsecase
}

func NewAuthorController(usecase *usecase.AuthorUsecase) *AuthorController {
	return &AuthorController{usecase: usecase}
}

func (bc *AuthorController) Create(c echo.Context) error {
	req := &model.CreateAuthorRequest{}
	res := map[string]any{}
	if err := c.Bind(req); err != nil {
		res["error"] = err.Error()
		return c.JSON(http.StatusBadRequest, res)
	}
	author, info := bc.usecase.Add(c.Request().Context(), req)
	res["message"] = info.Message
	if author != nil {
		res["data"] = author
	}
	if info.Err != nil {
		res["error"] = info.Err.Error()
	}
	return c.JSON(info.Code, res)
}

func (bc *AuthorController) Get(c echo.Context) error {
	res := map[string]any{}
	authors, info := bc.usecase.Get(c.Request().Context())
	res["message"] = info.Message
	if info.Err != nil {
		res["error"] = info.Err.Error()
	}
	if authors != nil {
		res["data"] = authors
	}
	return c.JSON(info.Code, res)
}

func (bc *AuthorController) GetById(c echo.Context) error {
	res := map[string]any{}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		res["message"] = "failed to get author"
		res["error"] = err.Error()
		return c.JSON(http.StatusBadRequest, res)
	}

	author, info := bc.usecase.GetById(c.Request().Context(), uint(id))
	res["message"] = info.Message
	if author != nil {
		res["data"] = author
	}
	if info.Err != nil {
		res["error"] = info.Err.Error()
	}
	return c.JSON(info.Code, res)
}

func (bc *AuthorController) Update(c echo.Context) error {
	res := map[string]any{}
	req := &model.UpdateAuthorRequest{}
	if err := c.Bind(req); err != nil {
		res["error"] = err.Error()
		res["message"] = "failed to update author"
		return c.JSON(http.StatusBadRequest, res)
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		res["error"] = err.Error()
		res["message"] = "failed to update author"
		return c.JSON(http.StatusBadRequest, res)
	}
	req.Id = uint(id)

	author, info := bc.usecase.Update(c.Request().Context(), req)
	if author != nil {
		res["data"] = author
	}
	if info.Err != nil {
		res["error"] = info.Err.Error()
	}
	return c.JSON(info.Code, res)
}

func (bc *AuthorController) Delete(c echo.Context) error {
	res := map[string]any{}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		res["message"] = "failed to delete author"
		res["error"] = err.Error()
		return c.JSON(http.StatusBadRequest, res)
	}

	author, info := bc.usecase.Delete(c.Request().Context(), uint(id))
	res["message"] = info.Message
	if author != nil {
		res["data"] = author
	}
	if info.Err != nil {
		res["error"] = info.Err.Error()
	}
	return c.JSON(info.Code, res)
}
