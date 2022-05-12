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

type UserHandlerTestSuite struct {
	suite.Suite
	server      *echo.Echo
	idGenerator *mockPkg.IDGenerator
	userService *mocks.UserService
}

func (s *UserHandlerTestSuite) SetupTest() {
	e := echo.New()
	e.Validator = validator.NewAPIValidator()

	idGenerator := new(mockPkg.IDGenerator)
	mockUserService := new(mocks.UserService)

	s.server = e
	s.idGenerator = idGenerator
	s.userService = mockUserService
}

func (s *UserHandlerTestSuite) TestCreateUser_ShouldReturnSuccess() {
	userID := uuid.NewString()
	user := &dto.UserRequest{
		ID:        userID,
		Email:     "test1@gmail.com",
		Fullname:  "Test1",
		Password:  "123456",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	body, err := json.Marshal(user)
	s.NoError(err)
	bodyJSON := string(body)

	req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(string(bodyJSON)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	userEntity := user.ToEntity()
	s.userService.
		On("CreateUser", mock.AnythingOfType("*entity.User")).
		Return(userEntity, nil)

	s.idGenerator.On("GenerateNewID").Return(userID)

	NewUserHandler(s.server, s.userService, s.idGenerator)
	s.server.ServeHTTP(rec, req)

	userResp := dto.NewUserResponseFromEntity(userEntity)
	resp := dto.NewSuccessResponse(userResp)
	respBody, err := json.Marshal(resp)
	s.NoError(err)

	// echo framework adds a new line at end of JSON string
	respBodyJSON := string(respBody) + "\n"
	s.Equal(http.StatusCreated, rec.Code)
	s.EqualValues(rec.Body.String(), respBodyJSON)
}

func (s *UserHandlerTestSuite) TestCreateUser_ShouldReturnInvalidValidation() {
	userID := uuid.NewString()
	user := &dto.UserRequest{
		ID:        userID,
		Email:     "invalid_email",
		Fullname:  "Test1",
		Password:  "123456",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	body, err := json.Marshal(user)
	s.NoError(err)
	bodyJSON := string(body)

	req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(string(bodyJSON)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	NewUserHandler(s.server, s.userService, s.idGenerator)
	s.server.ServeHTTP(rec, req)

	validationErrs := []string{
		"invalid_email",
	}
	resp := dto.NewErrorResponse(validationErrs)
	respBody, err := json.Marshal(resp)
	s.NoError(err)

	// echo framework adds a new line at end of JSON string
	respBodyJSON := string(respBody) + "\n"

	s.Equal(http.StatusBadRequest, rec.Code)
	s.EqualValues(rec.Body.String(), respBodyJSON)
}

func (s *UserHandlerTestSuite) TestCreateUser_ShouldReturnEmailExistedError() {
	userID := uuid.NewString()
	user := &dto.UserRequest{
		ID:        userID,
		Email:     "test1@gmail.com",
		Fullname:  "Test1",
		Password:  "123456",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	body, err := json.Marshal(user)
	s.NoError(err)
	bodyJSON := string(body)

	req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(string(bodyJSON)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	s.userService.
		On("CreateUser", mock.AnythingOfType("*entity.User")).
		Return(nil, errs.ErrEmailExisted)

	s.idGenerator.On("GenerateNewID").Return(userID)

	NewUserHandler(s.server, s.userService, s.idGenerator)
	s.server.ServeHTTP(rec, req)

	resp := dto.NewErrorResponse(errs.ErrEmailExisted.Error())
	respBody, err := json.Marshal(resp)
	s.NoError(err)

	// echo framework adds a new line at end of JSON string
	respBodyJSON := string(respBody) + "\n"
	s.Equal(http.StatusBadRequest, rec.Code)
	s.EqualValues(rec.Body.String(), respBodyJSON)
}

func (s *UserHandlerTestSuite) TestCreateUser_ShouldReturnServerError() {
	userID := uuid.NewString()
	user := &dto.UserRequest{
		ID:        userID,
		Email:     "test1@gmail.com",
		Fullname:  "Test1",
		Password:  "123456",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	body, err := json.Marshal(user)
	s.NoError(err)
	bodyJSON := string(body)

	req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(string(bodyJSON)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	s.userService.
		On("CreateUser", mock.AnythingOfType("*entity.User")).
		Return(nil, errs.ErrInternalServer)

	s.idGenerator.On("GenerateNewID").Return(userID)

	NewUserHandler(s.server, s.userService, s.idGenerator)
	s.server.ServeHTTP(rec, req)

	resp := dto.NewErrorResponse(errs.ErrInternalServer.Error())
	respBody, err := json.Marshal(resp)
	s.NoError(err)

	// echo framework adds a new line at end of JSON string
	respBodyJSON := string(respBody) + "\n"
	s.Equal(http.StatusInternalServerError, rec.Code)
	s.EqualValues(rec.Body.String(), respBodyJSON)
}

func (s *UserHandlerTestSuite) TestCreateUser_ShouldReturnBadRequest() {
	bodyJSON := "invalid json data"

	req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(string(bodyJSON)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	NewUserHandler(s.server, s.userService, s.idGenerator)
	s.server.ServeHTTP(rec, req)

	resp := dto.NewErrorResponse(errs.ErrInvalidRequestBody.Error())
	respBody, err := json.Marshal(resp)
	s.NoError(err)

	// echo framework adds a new line at end of JSON string
	respBodyJSON := string(respBody) + "\n"
	s.Equal(http.StatusBadRequest, rec.Code)
	s.EqualValues(rec.Body.String(), respBodyJSON)
}

func (s *UserHandlerTestSuite) TestLogin_ShouldReturnSuccess() {
	user := &dto.UserLoginRequest{
		Email:    "test1@gmail.com",
		Password: "123456",
	}

	body, err := json.Marshal(user)
	s.NoError(err)
	bodyJSON := string(body)

	req := httptest.NewRequest(http.MethodPost, "/users/login", strings.NewReader(string(bodyJSON)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	generatedToken := "generated_token"
	s.userService.
		On("Login", user.Email, user.Password).
		Return(generatedToken, nil)

	NewUserHandler(s.server, s.userService, s.idGenerator)
	s.server.ServeHTTP(rec, req)

	data := map[string]string{
		"token": generatedToken,
	}
	resp := dto.NewSuccessResponse(data)
	respBody, err := json.Marshal(resp)
	s.NoError(err)

	// echo framework adds a new line at end of JSON string
	respBodyJSON := string(respBody) + "\n"
	s.Equal(http.StatusCreated, rec.Code)
	s.EqualValues(rec.Body.String(), respBodyJSON)
}

func (s *UserHandlerTestSuite) TestLogin_ShouldReturnServerError() {
	user := &dto.UserLoginRequest{
		Email:    "test1@gmail.com",
		Password: "123456",
	}

	body, err := json.Marshal(user)
	s.NoError(err)
	bodyJSON := string(body)

	req := httptest.NewRequest(http.MethodPost, "/users/login", strings.NewReader(string(bodyJSON)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	s.userService.
		On("Login", user.Email, user.Password).
		Return("", errs.ErrInternalServer)

	NewUserHandler(s.server, s.userService, s.idGenerator)
	s.server.ServeHTTP(rec, req)

	resp := dto.NewErrorResponse(errs.ErrInternalServer.Error())
	respBody, err := json.Marshal(resp)
	s.NoError(err)

	// echo framework adds a new line at end of JSON string
	respBodyJSON := string(respBody) + "\n"
	s.Equal(http.StatusInternalServerError, rec.Code)
	s.EqualValues(respBodyJSON, rec.Body.String())
}

func (s *UserHandlerTestSuite) TestLogin_ShouldReturnBadRequest() {
	bodyJSON := "invalid json data"

	req := httptest.NewRequest(http.MethodPost, "/users/login", strings.NewReader(string(bodyJSON)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	NewUserHandler(s.server, s.userService, s.idGenerator)
	s.server.ServeHTTP(rec, req)

	resp := dto.NewErrorResponse(errs.ErrInvalidRequestBody.Error())
	respBody, err := json.Marshal(resp)
	s.NoError(err)

	// echo framework adds a new line at end of JSON string
	respBodyJSON := string(respBody) + "\n"
	s.Equal(http.StatusBadRequest, rec.Code)
	s.EqualValues(rec.Body.String(), respBodyJSON)
}

func (s *UserHandlerTestSuite) TestLogin_ShouldReturnPasswordInvalid() {
	user := &dto.UserLoginRequest{
		Email:    "test1@gmail.com",
		Password: "123457",
	}

	body, err := json.Marshal(user)
	s.NoError(err)
	bodyJSON := string(body)

	req := httptest.NewRequest(http.MethodPost, "/users/login", strings.NewReader(string(bodyJSON)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	s.userService.
		On("Login", user.Email, user.Password).
		Return("", errs.ErrPasswordInvalid)

	NewUserHandler(s.server, s.userService, s.idGenerator)
	s.server.ServeHTTP(rec, req)

	resp := dto.NewErrorResponse(errs.ErrPasswordInvalid.Error())
	respBody, err := json.Marshal(resp)
	s.NoError(err)

	// echo framework adds a new line at end of JSON string
	respBodyJSON := string(respBody) + "\n"
	s.Equal(http.StatusBadRequest, rec.Code)
	s.EqualValues(respBodyJSON, rec.Body.String())
}

func (s *UserHandlerTestSuite) TestLogin_ShouldReturnEmailInvalidValidation() {
	user := &dto.UserLoginRequest{
		Email:    "invalid_email",
		Password: "123457",
	}

	body, err := json.Marshal(user)
	s.NoError(err)
	bodyJSON := string(body)

	req := httptest.NewRequest(http.MethodPost, "/users/login", strings.NewReader(string(bodyJSON)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	NewUserHandler(s.server, s.userService, s.idGenerator)
	s.server.ServeHTTP(rec, req)

	validationErrs := []string{
		"invalid_email",
	}
	resp := dto.NewErrorResponse(validationErrs)
	respBody, err := json.Marshal(resp)
	s.NoError(err)

	// echo framework adds a new line at end of JSON string
	respBodyJSON := string(respBody) + "\n"
	s.Equal(http.StatusBadRequest, rec.Code)
	s.EqualValues(respBodyJSON, rec.Body.String())
}

func TestCreateUser(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(UserHandlerTestSuite))
}
