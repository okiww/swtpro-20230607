package handler

import (
	"github.com/SawitProRecruitment/UserService/repository"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestLoginEndpoint_Success(t *testing.T) {
	// Create a new Echo instance
	e := echo.New()
	ctr := gomock.NewController(t)
	defer ctr.Finish()
	// Create a request with a valid payload
	payload := `{"phone_number": "+6285797178075", "password": "Passw0rd!"}`
	req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(payload))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Create a mock server instance with an existing user
	mockRepo := repository.NewMockRepositoryInterface(ctr)
	// Set up expectations for GetUserByPhoneNumber
	mockRepo.EXPECT().GetUserByPhoneNumber(gomock.Any(), gomock.Any()).Return(&repository.GetUserByPhoneNumberOutput{
		Id:          1,
		FullName:    "John Doe",
		PhoneNumber: "+6285797178075",
		Password:    "$2a$12$d9UJSAHKJeIQmA4BOA9SduFkxe6vcokolrRgJRS3dgfhyO0/NIFi6", // hashed password
	}, nil)
	s := &Server{
		Repository: mockRepo,
	}

	// Call the Login endpoint
	err := s.Login(c)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestLoginEndpoint_Failure(t *testing.T) {
	// Create a new Echo instance
	e := echo.New()
	ctr := gomock.NewController(t)
	defer ctr.Finish()
	// Create a request with a valid payload
	payload := `{"phone_number": "+6285797178075", "password": ""}`
	req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(payload))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Create a mock server instance with an existing user
	mockRepo := repository.NewMockRepositoryInterface(ctr)
	s := &Server{
		Repository: mockRepo,
	}

	// Call the Login endpoint
	err := s.Login(c)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestUpdateEndpoint_Success(t *testing.T) {
	// Create a new Echo instance
	e := echo.New()
	ctr := gomock.NewController(t)
	defer ctr.Finish()
	// Create a request with a valid payload
	payload := `{"full_name": "John Doe", "phone_number": "+6285797178075"}`
	req := httptest.NewRequest(http.MethodPut, "/update", strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set(echo.HeaderAuthorization, "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6NSwiZnVsbG5hbWUiOiJmdWxsTmFtZSIsImV4cCI6MTcwNjUxMTU0MH0.pU5wD8U_GGI9V-5u6RxHSU3ZlWrtF01YZEYTxM5TkTo")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Create a mock server instance with an existing user
	mockRepo := repository.NewMockRepositoryInterface(ctr)
	// Set up expectations for GetUserByPhoneNumber
	mockRepo.EXPECT().GetUserByPhoneNumber(gomock.Any(), gomock.Any()).Return(nil, nil)

	// Set up expectations for UpdateUserById
	mockRepo.EXPECT().UpdateUserById(gomock.Any(), gomock.Any(), 5).Return(&repository.UpdateUserByIdOutput{
		Id:          5,
		FullName:    "John Doe",
		PhoneNumber: "+6285797178075",
	}, nil)
	s := &Server{
		Repository: mockRepo,
	}
	// Call the Update endpoint
	err := s.Update(c)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestUpdateEndpoint_Failure(t *testing.T) {
	// Create a new Echo instance
	e := echo.New()
	ctr := gomock.NewController(t)
	defer ctr.Finish()
	// Create a request with a valid payload
	payload := `{"full_name": "John Doe", "phone_number": "+6285797178075"}`
	req := httptest.NewRequest(http.MethodPut, "/update", strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Create a mock server instance with an existing user
	mockRepo := repository.NewMockRepositoryInterface(ctr)
	s := &Server{
		Repository: mockRepo,
	}
	// Call the Update endpoint
	err := s.Update(c)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, http.StatusUnauthorized, rec.Code)
}

func TestGetProfile_Success(t *testing.T) {
	// Create a new Echo instance
	e := echo.New()
	ctr := gomock.NewController(t)
	defer ctr.Finish()
	// Create a request with a valid payload
	req := httptest.NewRequest(http.MethodGet, "/profile", nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set(echo.HeaderAuthorization, "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6NSwiZnVsbG5hbWUiOiJmdWxsTmFtZSIsImV4cCI6MTcwNjUxMTU0MH0.pU5wD8U_GGI9V-5u6RxHSU3ZlWrtF01YZEYTxM5TkTo")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Create a mock server instance with an existing user
	mockRepo := repository.NewMockRepositoryInterface(ctr)
	// Set up expectations for GetUserById
	mockRepo.EXPECT().GetUserById(gomock.Any(), gomock.Any()).Return(&repository.User{
		FullName:    "John Doe",
		PhoneNumber: "+6285797178075",
	}, nil)
	s := &Server{
		Repository: mockRepo,
	}
	// Call the Update endpoint
	err := s.Profile(c)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestGetProfile_Failure(t *testing.T) {
	// Create a new Echo instance
	e := echo.New()
	ctr := gomock.NewController(t)
	defer ctr.Finish()
	// Create a request with a valid payload
	req := httptest.NewRequest(http.MethodGet, "/profile", nil)
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Create a mock server instance with an existing user
	mockRepo := repository.NewMockRepositoryInterface(ctr)
	s := &Server{
		Repository: mockRepo,
	}
	// Call the Update endpoint
	err := s.Profile(c)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, http.StatusUnauthorized, rec.Code)
}
