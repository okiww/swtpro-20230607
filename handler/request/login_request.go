package request

import (
	"errors"
	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"time"
)

// Secret key for signing the token. Keep it secure!
var secretKey = []byte("sawitpro_key")

func ValidateLoginPayload(payload *generated.LoginPayload) error {
	if payload.PhoneNumber == "" {
		return errors.New("phone_number cannot be empty")
	}

	if payload.Password == "" {
		return errors.New("password cannot be empty")
	}

	return nil
}

// UserClaims is an example of custom claims you might want to include in the token.
type UserClaims struct {
	Id       int    `json:"id"`
	Fullname string `json:"fullname"`
	jwt.StandardClaims
}

func GenerateJWTToken(id int, fullName string) (string, error) {
	claims := UserClaims{
		Id:       id,
		Fullname: "fullName",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
		},
	}

	// Create the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ValidateAccessToken validates the access token and returns user claims
func ValidateAccessToken(tokenString string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*UserClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

// CheckPasswordHash compares a hashed password with its possible plaintext equivalent
func CheckPasswordHash(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
