package controller

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/thiccpan/library_information_system/internal/model"
	"github.com/thiccpan/library_information_system/internal/usecase"
)

type LoanController struct {
	usecase *usecase.LoanUsecase
}

func NewLoanController(usecase *usecase.LoanUsecase) *LoanController {
	return &LoanController{
		usecase: usecase,
	}
}

func (lc *LoanController) Create(c echo.Context) error {
	res := map[string]any{}
	req := &model.CommandLoanRequest{}
	if err := c.Bind(req); err != nil {
		res["message"] = "failed to add new loan"
		res["error"] = err.Error()
		return c.JSON(http.StatusBadRequest, res)
	}
	loan, info := lc.usecase.Add(c.Request().Context(), req)
	res["message"] = info.Message
	if loan != nil {
		res["data"] = loan
	}
	if info.Err != nil {
		res["error"] = info.Err.Error()
	}
	return c.JSON(info.Code, res)
}

func (lc *LoanController) Get(c echo.Context) error {
	res := map[string]any{}
	req := &model.QueryLoanRequest{}
	loans, info := lc.usecase.Get(c.Request().Context(), req)
	res["message"] = info.Message
	if loans != nil {
		res["data"] = loans
	}
	if info.Err != nil {
		res["error"] = info.Err.Error()
	}
	return c.JSON(info.Code, res)
}

func (lc *LoanController) GetById(c echo.Context) error {
	res := map[string]any{}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		res["message"] = "invalid id"
		res["error"] = err.Error()
		return c.JSON(http.StatusBadRequest, res)
	}
	req := &model.QueryLoanRequest{Id: uint(id)}
	loan, info := lc.usecase.GetById(c.Request().Context(), req)
	if info.Err != nil {
		res["message"] = "failed to get loan"
		res["error"] = info.Err.Error()
		return c.JSON(http.StatusBadRequest, res)
	}
	res["message"] = info.Message
	if loan != nil {
		res["data"] = loan
	}
	if info.Err != nil {
		res["error"] = info.Err.Error()
	}
	return c.JSON(info.Code, res)
}

func (lc *LoanController) Update(c echo.Context) error {
	res := map[string]any{}
	req := &model.CommandLoanRequest{}
	if err := c.Bind(req); err != nil {
		res["message"] = "failed update loan"
		res["error"] = err.Error()
		return c.JSON(http.StatusBadRequest, res)
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		res["message"] = "invalid id"
		res["error"] = err.Error()
		return c.JSON(http.StatusBadRequest, res)
	}
	req.Id = uint(id)
	loan, info := lc.usecase.Update(c.Request().Context(), req)
	res["message"] = info.Message
	if loan != nil {
		res["data"] = loan
	}
	if info.Err != nil {
		res["error"] = info.Err.Error()
	}
	return c.JSON(info.Code, res)
}

func (lc *LoanController) Delete(c echo.Context) error {
	res := map[string]any{}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		res["message"] = "invalid id"
		res["error"] = err.Error()
		return c.JSON(http.StatusBadRequest, res)
	}
	req := &model.CommandLoanRequest{Id: uint(id)}
	loan, info := lc.usecase.Delete(c.Request().Context(), req)
	res["message"] = info.Message
	if loan != nil {
		res["data"] = loan
	}
	if info.Err != nil {
		res["error"] = info.Err.Error()
	}
	return c.JSON(info.Code, res)
}