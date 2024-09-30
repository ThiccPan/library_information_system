package controller

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/thiccpan/library_information_system/internal/model"
	"github.com/thiccpan/library_information_system/internal/usecase"
)

type BookController struct {
	usecase *usecase.BookUsecase
}

func NewBookController(usecase *usecase.BookUsecase) *BookController {
	return &BookController{
		usecase: usecase,
	}
}

func (bc *BookController) Get(c echo.Context) error {
	res := map[string]any{}
	req := &model.QueryBookRequest{}
	books, info := bc.usecase.Get(c.Request().Context(), req)
	if info.Err != nil {
		res["message"] = "failed to get books"
		res["error"] = info.Err.Error()
		return c.JSON(http.StatusBadRequest, res)
	}
	res["message"] = info.Message
	if books != nil {
		res["data"] = books
	}
	if info.Err != nil {
		res["error"] = info.Err.Error()
	}
	return c.JSON(info.Code, res)
}

func (bc *BookController) GetById(c echo.Context) error {
	res := map[string]any{}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		res["message"] = "failed to get books"
		res["error"] = err.Error()
		return c.JSON(http.StatusBadRequest, res)
	}
	req := &model.QueryBookRequest{Id: uint(id)}
	book, info := bc.usecase.GetById(c.Request().Context(), req)
	if info.Err != nil {
		res["message"] = "failed to get books"
		res["error"] = info.Err.Error()
		return c.JSON(http.StatusBadRequest, res)
	}
	res["message"] = info.Message
	if book != nil {
		res["data"] = book
	}
	if info.Err != nil {
		res["error"] = info.Err.Error()
	}
	return c.JSON(info.Code, res)
}

func (bc *BookController) Create(c echo.Context) error {
	res := map[string]any{}
	req := &model.CreateBookRequest{}
	if err := c.Bind(req); err != nil {
		res["message"] = "failed to add new book"
		res["error"] = err.Error()
		return c.JSON(http.StatusBadRequest, res)
	}
	book, info := bc.usecase.Add(c.Request().Context(), req)
	if info.Err != nil {
		res["message"] = "failed to add new book"
		res["error"] = info.Err.Error()
		return c.JSON(http.StatusBadRequest, res)
	}
	res["message"] = info.Message
	if book != nil {
		res["data"] = book
	}
	if info.Err != nil {
		res["error"] = info.Err.Error()
	}
	return c.JSON(info.Code, res)
}

func (bc *BookController) Update(c echo.Context) error {
	res := map[string]any{}
	req := &model.UpdateBookRequest{}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		res["message"] = "failed to update book"
		res["error"] = err.Error()
		return c.JSON(http.StatusBadRequest, res)
	}
	if err := c.Bind(req); err != nil {
		res["message"] = "failed to update book"
		res["error"] = err.Error()
		return c.JSON(http.StatusBadRequest, res)
	}
	req.Id = uint(id)
	file, _ := c.FormFile("image")
	if file != nil {
		req.CoverImg = file
	}

	book, info := bc.usecase.Update(c.Request().Context(), req)
	if info.Err != nil {
		res["message"] = "failed to add new book"
		res["error"] = info.Err.Error()
		return c.JSON(http.StatusBadRequest, res)
	}
	res["message"] = info.Message
	if book != nil {
		res["data"] = book
	}
	if info.Err != nil {
		res["error"] = info.Err.Error()
	}
	return c.JSON(info.Code, res)
}

func (bc *BookController) Delete(c echo.Context) error {
	res := map[string]any{}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		res["message"] = "failed to delete book"
		res["error"] = err.Error()
		return c.JSON(http.StatusBadRequest, res)
	}
	req := &model.QueryBookRequest{Id: uint(id)}
	book, info := bc.usecase.Delete(c.Request().Context(), req)
	if info.Err != nil {
		res["message"] = "failed to delete book"
		res["error"] = info.Err.Error()
		return c.JSON(http.StatusBadRequest, res)
	}
	res["message"] = info.Message
	if book != nil {
		res["data"] = book
	}
	if info.Err != nil {
		res["error"] = info.Err.Error()
	}
	return c.JSON(info.Code, res)
}
