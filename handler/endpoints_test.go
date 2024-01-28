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
