package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"github.com/vuhn/go-app-sample/application/api/validator"
	"github.com/vuhn/go-app-sample/application/dto"
	"github.com/vuhn/go-app-sample/errs"
	mockPkg "github.com/vuhn/go-app-sample/pkg/mocks"
	"github.com/vuhn/go-app-sample/service/mocks"
)

type TaskHandlerTestSuite struct {
	suite.Suite
	server      *echo.Echo
	idGenerator *mockPkg.IDGenerator
	taskService *mocks.TaskService
}

func (s *TaskHandlerTestSuite) SetupTest() {
	e := echo.New()
	e.Validator = validator.NewAPIValidator()

	idGenerator := new(mockPkg.IDGenerator)
	mockUserService := new(mocks.TaskService)

	s.server = e
	s.idGenerator = idGenerator
	s.taskService = mockUserService
}

func (s *TaskHandlerTestSuite) TestCreateTask_ShouldReturnSuccess() {
	taskID := uuid.NewString()
	task := &dto.CreateTaskRequest{
		Title:       "Task 1",
		Description: "Task 1 Description",
		Assignee:    "1",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	body, err := json.Marshal(task)
	s.NoError(err)
	bodyJSON := string(body)

	req := httptest.NewRequest(http.MethodPost, "/tasks", strings.NewReader(string(bodyJSON)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	taskEntity := task.ToEntity()
	s.taskService.
		On("CreateTask", mock.AnythingOfType("*entity.Task")).
		Return(taskEntity, nil)

	s.idGenerator.On("GenerateNewID").Return(taskID)

	NewTaskHandler(s.server, s.taskService, s.idGenerator)
	s.server.ServeHTTP(rec, req)

	taskResp := dto.NewTaskResponseFromEntity(taskEntity)
	resp := dto.NewSuccessResponse(taskResp)
	respBody, err := json.Marshal(resp)
	s.NoError(err)

	// echo framework adds a new line at end of JSON string
	respBodyJSON := string(respBody) + "\n"
	s.Equal(http.StatusCreated, rec.Code)
	s.EqualValues(rec.Body.String(), respBodyJSON)
}

func (s *TaskHandlerTestSuite) TestCreateTask_ShouldReturnErrorInvalidTitle() {
	task := &dto.CreateTaskRequest{
		Title:       "",
		Description: "Task 1 Description",
		Assignee:    "1",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	body, err := json.Marshal(task)
	s.NoError(err)
	bodyJSON := string(body)

	req := httptest.NewRequest(http.MethodPost, "/tasks", strings.NewReader(bodyJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	NewTaskHandler(s.server, s.taskService, s.idGenerator)
	s.server.ServeHTTP(rec, req)

	validationErrs := []string{
		"invalid_title",
	}
	resp := dto.NewErrorResponse(validationErrs)
	respBody, err := json.Marshal(resp)
	s.NoError(err)

	// echo framework adds a new line at end of JSON string
	respBodyJSON := string(respBody) + "\n"
	s.Equal(http.StatusBadRequest, rec.Code)
	s.EqualValues(respBodyJSON, rec.Body.String())
}

func (s *TaskHandlerTestSuite) TestCreateTask_ShouldReturnBadRequestWithInvalidRequestBody() {
	bodyJSON := "invalid json data"

	req := httptest.NewRequest(http.MethodPost, "/tasks", strings.NewReader(bodyJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	NewTaskHandler(s.server, s.taskService, s.idGenerator)
	s.server.ServeHTTP(rec, req)

	resp := dto.NewErrorResponse(errs.ErrInvalidRequestBody.Error())
	respBody, err := json.Marshal(resp)
	s.NoError(err)

	// echo framework adds a new line at end of JSON string
	respBodyJSON := string(respBody) + "\n"
	s.Equal(http.StatusBadRequest, rec.Code)
	s.EqualValues(rec.Body.String(), respBodyJSON)
}

func (s *TaskHandlerTestSuite) TestCreateTask_ShouldReturnErrorUserNotFound() {
	taskID := uuid.NewString()
	task := &dto.CreateTaskRequest{
		Title:       "Task 1",
		Description: "Task 1 Description",
		Assignee:    "100",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	body, err := json.Marshal(task)
	s.NoError(err)
	bodyJSON := string(body)

	req := httptest.NewRequest(http.MethodPost, "/tasks", strings.NewReader(bodyJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	s.taskService.
		On("CreateTask", mock.AnythingOfType("*entity.Task")).
		Return(nil, errs.ErrUserNotFound)

	s.idGenerator.On("GenerateNewID").Return(taskID)

	NewTaskHandler(s.server, s.taskService, s.idGenerator)
	s.server.ServeHTTP(rec, req)

	resp := dto.NewErrorResponse(errs.ErrUserNotFound.Error())
	respBody, err := json.Marshal(resp)
	s.NoError(err)

	// echo framework adds a new line at end of JSON string
	respBodyJSON := string(respBody) + "\n"
	s.Equal(http.StatusBadRequest, rec.Code)
	s.EqualValues(rec.Body.String(), respBodyJSON)
}

func (s *TaskHandlerTestSuite) TestCreateTask_ShouldReturnServerErrorWhenCreateUser() {
	taskID := uuid.NewString()
	task := &dto.CreateTaskRequest{
		Title:       "Task 1",
		Description: "Task 1 Description",
		Assignee:    "100",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	body, err := json.Marshal(task)
	s.NoError(err)
	bodyJSON := string(body)

	req := httptest.NewRequest(http.MethodPost, "/tasks", strings.NewReader(bodyJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	s.taskService.
		On("CreateTask", mock.AnythingOfType("*entity.Task")).
		Return(nil, errs.ErrInternalServer)

	s.idGenerator.On("GenerateNewID").Return(taskID)

	NewTaskHandler(s.server, s.taskService, s.idGenerator)
	s.server.ServeHTTP(rec, req)

	resp := dto.NewErrorResponse(errs.ErrInternalServer.Error())
	respBody, err := json.Marshal(resp)
	s.NoError(err)

	// echo framework adds a new line at end of JSON string
	respBodyJSON := string(respBody) + "\n"
	s.Equal(http.StatusInternalServerError, rec.Code)
	s.EqualValues(rec.Body.String(), respBodyJSON)
}

func TestTaskHandler(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(TaskHandlerTestSuite))
}
