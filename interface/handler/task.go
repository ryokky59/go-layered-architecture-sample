package handler

import (
	"net/http"
	"strconv"

	"sample/usecase"
	"sample/domain/model"
	"github.com/labstack/echo"
)

type TaskHandler struct {
	taskUsecase usecase.TaskUsecase
}

func NewTaskHandler(taskUsecase usecase.TaskUsecase) TaskHandler {
	return TaskHandler{taskUsecase: taskUsecase}
}

func (th *TaskHandler) Post() echo.HandlerFunc {
	return func(c echo.Context) error {
        var task model.Task
        if err := c.Bind(&task); err != nil {
			return c.JSON(http.StatusBadRequest, task)
		}

		createdTask, err := th.taskUsecase.Create(&task)
		if err != nil {
			return c.JSON(http.StatusBadRequest, createdTask)
		}

        return c.JSON(http.StatusCreated, createdTask)
    }
}

func (th *TaskHandler) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi((c.Param("id")))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		foundTask, err := th.taskUsecase.FindByID(id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, foundTask)
		}

        return c.JSON(http.StatusOK, foundTask)
    }
}

func (th *TaskHandler) Put() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi((c.Param("id")))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		var task model.Task
        if err := c.Bind(&task); err != nil {
			return c.JSON(http.StatusBadRequest, task)
		}

		updatedTask, err := th.taskUsecase.Update(id, &task)
		if err != nil {
			return c.JSON(http.StatusBadRequest, updatedTask)
		}

        return c.JSON(http.StatusOK, updatedTask)
    }
}

func (th *TaskHandler) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi((c.Param("id")))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		err = th.taskUsecase.Delete(id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

        return c.NoContent(http.StatusNoContent)
    }
}
