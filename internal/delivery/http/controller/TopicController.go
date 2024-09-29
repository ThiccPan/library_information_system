package controller

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/thiccpan/library_information_system/internal/model"
	"github.com/thiccpan/library_information_system/internal/usecase"
)

type TopicController struct {
	usecase *usecase.TopicUsecase
}

func NewTopicController(usecase *usecase.TopicUsecase) *TopicController {
	return &TopicController{usecase: usecase}
}

func (bc *TopicController) Create(c echo.Context) error {
	req := &model.CreateTopicRequest{}
	res := map[string]any{}
	if err := c.Bind(req); err != nil {
		res["error"] = err.Error()
		return c.JSON(http.StatusBadRequest, res)
	}
	topic, info := bc.usecase.Add(c.Request().Context(), req)
	res["message"] = info.Message
	if topic != nil {
		res["data"] = topic
	}
	if info.Err != nil {
		res["error"] = info.Err.Error()
	}
	return c.JSON(info.Code, res)
}

func (bc *TopicController) Get(c echo.Context) error {
	res := map[string]any{}
	topics, info := bc.usecase.Get(c.Request().Context())
	res["message"] = info.Message
	if info.Err != nil {
		res["error"] = info.Err.Error()
	}
	if topics != nil {
		res["data"] = topics
	}
	return c.JSON(info.Code, res)
}

func (bc *TopicController) GetById(c echo.Context) error {
	res := map[string]any{}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		res["message"] = "failed to get topic"
		res["error"] = err.Error()
		return c.JSON(http.StatusBadRequest, res)
	}

	topic, info := bc.usecase.GetById(c.Request().Context(), uint(id))
	res["message"] = info.Message
	if topic != nil {
		res["data"] = topic
	}
	if info.Err != nil {
		res["error"] = info.Err.Error()
	}
	return c.JSON(info.Code, res)
}

func (bc *TopicController) Update(c echo.Context) error {
	res := map[string]any{}
	req := &model.UpdateTopicRequest{}
	if err := c.Bind(req); err != nil {
		res["error"] = err.Error()
		res["message"] = "failed to update topic"
		return c.JSON(http.StatusBadRequest, res)
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		res["error"] = err.Error()
		res["message"] = "failed to update topic"
		return c.JSON(http.StatusBadRequest, res)
	}
	req.Id = uint(id)

	topic, info := bc.usecase.Update(c.Request().Context(), req)
	if topic != nil {
		res["data"] = topic
	}
	if info.Err != nil {
		res["error"] = info.Err.Error()
	}
	return c.JSON(info.Code, res)
}

func (bc *TopicController) Delete(c echo.Context) error {
	res := map[string]any{}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		res["message"] = "failed to delete topic"
		res["error"] = err.Error()
		return c.JSON(http.StatusBadRequest, res)
	}

	topic, info := bc.usecase.Delete(c.Request().Context(), uint(id))
	res["message"] = info.Message
	if topic != nil {
		res["data"] = topic
	}
	if info.Err != nil {
		res["error"] = info.Err.Error()
	}
	return c.JSON(info.Code, res)
}
