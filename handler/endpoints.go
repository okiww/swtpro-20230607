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

func (s *Server) Login(ctx echo.Context) error {
	var payload generated.LoginPayload

	if err := ctx.Bind(&payload); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	err := request.ValidateLoginPayload(&payload)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": fmt.Sprintf("%s", err.Error())})
	}

	existingUser, err := s.Repository.GetUserByPhoneNumber(ctx.Request().Context(), payload.PhoneNumber)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
	}

	if existingUser == nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid username or password"})
	}

	// Check if the provided password matches the hashed password in the database
	if !request.CheckPasswordHash(payload.Password, existingUser.Password) {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Incorrect Password"})
	}

	accessToken, err := request.GenerateJWTToken(existingUser.Id, existingUser.FullName)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
	}
	return ctx.JSON(http.StatusOK, generated.LoginResponse{
		AccessToken: accessToken,
	})
}

func (s *Server) Update(ctx echo.Context) error {
	authHeader := ctx.Request().Header.Get("Authorization")
	// Check if "Authorization" header is not set
	if authHeader == "" {
		return ctx.JSON(http.StatusUnauthorized, map[string]string{"error": "Authorization header not set"})
	}
	// Remove "Bearer " prefix
	tokenString := authHeader[7:]

	// Validate the access token
	claims, err := request.ValidateAccessToken(tokenString)
	if err != nil {
		return ctx.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid Token"})
	}

	var payload generated.UpdatePayload
	if err := ctx.Bind(&payload); err != nil {
		fmt.Println(err.Error())
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	err = request.ValidateUpdatePayload(&payload)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": fmt.Sprintf("%s", err.Error())})
	}
	existingUser, err := s.Repository.GetUserByPhoneNumber(ctx.Request().Context(), payload.PhoneNumber)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
	}

	if existingUser != nil && existingUser.Id != claims.Id {
		return ctx.JSON(http.StatusConflict, map[string]string{"error": "phone number already registered"})
	}

	data, err := s.Repository.UpdateUserById(ctx.Request().Context(), &repository.UpdateUserByIdInput{
		FullName:    payload.FullName,
		PhoneNumber: payload.PhoneNumber,
	}, claims.Id)

	var resp generated.UpdateResponse
	resp.Message = "success update"
	resp.Data.FullName = data.FullName
	resp.Data.PhoneNumber = data.PhoneNumber

	return ctx.JSON(http.StatusOK, resp)
}

func (s *Server) Profile(ctx echo.Context) error {
	authHeader := ctx.Request().Header.Get("Authorization")
	// Check if "Authorization" header is not set
	if authHeader == "" {
		return ctx.JSON(http.StatusUnauthorized, map[string]string{"error": "Authorization header not set"})
	}
	// Remove "Bearer " prefix
	tokenString := authHeader[7:]
	// Validate the access token
	claims, err := request.ValidateAccessToken(tokenString)
	if err != nil {
		return ctx.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid Token"})
	}

	existingUser, err := s.Repository.GetUserById(ctx.Request().Context(), claims.Id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
	}

	if existingUser == nil {
		return ctx.JSON(http.StatusNotFound, map[string]string{"error": "User Not Found"})
	}

	return ctx.JSON(http.StatusOK, generated.ProfileResponse{
		FullName:    existingUser.FullName,
		PhoneNumber: existingUser.PhoneNumber,
	})
}
