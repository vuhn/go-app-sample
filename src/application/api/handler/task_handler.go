package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/vuhn/go-app-sample/application/dto"
	"github.com/vuhn/go-app-sample/errs"
	"github.com/vuhn/go-app-sample/pkg/idgenerator"
	"github.com/vuhn/go-app-sample/service"
)

type TaskHandler struct {
	taskService service.TaskService
	idGenerator idgenerator.IDGenerator
}

// NewTaskHandler setup rest api handlers for user
func NewTaskHandler(echo *echo.Echo,
	taskService service.TaskService,
	idGenerator idgenerator.IDGenerator,
) {
	handler := &TaskHandler{
		taskService: taskService,
		idGenerator: idGenerator,
	}
	echo.POST("/tasks", handler.CreateTask)
}

// CreateTask is method for create task api endpoint
func (h *TaskHandler) CreateTask(c echo.Context) error {
	task := dto.CreateTaskRequest{}
	if err := c.Bind(&task); err != nil {
		return c.JSON(http.StatusBadRequest, dto.NewErrorResponse(errs.ErrInvalidRequestBody.Error()))
	}

	if err := c.Validate(task); err != nil {
		return err
	}

	task.ID = h.idGenerator.GenerateNewID()
	taskEntity, err := h.taskService.CreateTask(task.ToEntity())
	if err == errs.ErrUserNotFound {
		return c.JSON(http.StatusBadRequest, dto.NewErrorResponse(err.Error()))
	}

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.NewErrorResponse(err.Error()))
	}

	taskDto := dto.NewTaskResponseFromEntity(taskEntity)
	return c.JSON(http.StatusCreated, dto.NewSuccessResponse(taskDto))
}
