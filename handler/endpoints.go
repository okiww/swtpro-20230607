package handler

import (
	"fmt"
	"github.com/SawitProRecruitment/UserService/handler/request"
	"github.com/SawitProRecruitment/UserService/repository"
	"net/http"

	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/labstack/echo/v4"
)

// This is just a test endpoint to get you started. Please delete this endpoint.
// (GET /hello)
func (s *Server) Hello(ctx echo.Context, params generated.HelloParams) error {

	var resp generated.HelloResponse
	resp.Message = fmt.Sprintf("Hello User %d", params.Id)
	return ctx.JSON(http.StatusOK, resp)
}

func (s *Server) Register(ctx echo.Context) error {
	var payload generated.RegisterPayload
	if err := ctx.Bind(&payload); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	err := request.ValidateRegisterPayload(&payload)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": fmt.Sprintf("%s", err.Error())})
	}

	existingUser, err := s.Repository.GetUserByPhoneNumber(ctx.Request().Context(), payload.PhoneNumber)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
	}

	if existingUser != nil {
		return ctx.JSON(http.StatusConflict, map[string]string{"error": "phone number already registered"})
	}

	// Hash and salt the password
	hashedPassword, err := request.HashPassword(payload.Password)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
	}

	id, err := s.Repository.CreateUserInput(ctx.Request().Context(), &repository.CreateUserInput{
		FullName:    payload.FullName,
		Password:    hashedPassword,
		PhoneNumber: payload.PhoneNumber,
	})

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
	}

	return ctx.JSON(http.StatusOK, generated.RegisterResponse{
		Message: "success register",
		Id:      id,
	})
}
