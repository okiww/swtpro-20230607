package request

import (
	"errors"
	"fmt"
	"github.com/SawitProRecruitment/UserService/generated"
	"golang.org/x/crypto/bcrypt"
	"regexp"
)

const (
	PHONE_NUMBER_RULES_MIN_LENGTH = 10
	PHONE_NUMBER_RULES_MAX_LENGTH = 14
	PHONE_NUMBER_RULES_PATTERN    = `^\+62[0-9]+$`

	PASSWORD_RULES_MIN_LENGTH = 6
	PASSWORD_RULES_MAX_LENGTH = 64

	FULLNAME_RULES_MIN_LENGTH = 3
	FULLNAME_RULES_MAX_LENGTH = 60
)

func ValidateRegisterPayload(payload *generated.RegisterPayload) error {
	if payload.PhoneNumber == "" {
		return errors.New("phone_number cannot be empty")
	}

	phoneNumberLen := len(payload.PhoneNumber)
	if phoneNumberLen < PHONE_NUMBER_RULES_MIN_LENGTH || phoneNumberLen > PHONE_NUMBER_RULES_MAX_LENGTH {
		return errors.New(fmt.Sprintf("phone_number should have minimum %d or maximum %d character", PHONE_NUMBER_RULES_MIN_LENGTH, PHONE_NUMBER_RULES_MAX_LENGTH-1))
	}

	if !validatePhoneNumber(payload.PhoneNumber) {
		return errors.New("phone_number must start with the Indonesia country code “+62”")
	}

	if payload.Password == "" {
		return errors.New("password cannot be empty")
	}

	passwordLen := len(payload.Password)
	if passwordLen < PASSWORD_RULES_MIN_LENGTH || passwordLen > PASSWORD_RULES_MAX_LENGTH {
		return errors.New(fmt.Sprintf("password should have minimum %d or maximum %d character", PASSWORD_RULES_MIN_LENGTH, PASSWORD_RULES_MAX_LENGTH))
	}

	if !validatePassword(payload.Password) {
		return errors.New("password must be containing at least 1 capital characters AND 1 number AND 1 special (non alpha-numeric) characters")
	}

	if payload.FullName == "" {
		return errors.New("full_name cannot be empty")
	}

	fullnameLen := len(payload.FullName)
	if fullnameLen < FULLNAME_RULES_MIN_LENGTH || fullnameLen > FULLNAME_RULES_MAX_LENGTH {
		return errors.New(fmt.Sprintf("full_name should have minimum %d or maximum %d character", FULLNAME_RULES_MIN_LENGTH, FULLNAME_RULES_MAX_LENGTH))
	}

	if !validatePassword(payload.Password) {
		return errors.New("password must be containing at least 1 capital characters AND 1 number AND 1 special (non alpha-numeric) characters")
	}

	return nil
}

func validatePassword(input string) bool {
	// Define the regular expressions for each condition
	containsUppercase := regexp.MustCompile(`[A-Z]`).MatchString
	containsDigit := regexp.MustCompile(`[0-9]`).MatchString
	containsSpecialChar := regexp.MustCompile(`[^A-Za-z0-9]`).MatchString
	validLength := regexp.MustCompile(`.{6,64}`).MatchString

	// Check all conditions
	return containsUppercase(input) && containsDigit(input) && containsSpecialChar(input) && validLength(input)
}

func validatePhoneNumber(input string) bool {
	// Define the regular expression pattern
	pattern := PHONE_NUMBER_RULES_PATTERN

	// Compile the regular expression
	regex := regexp.MustCompile(pattern)

	// Test the input string against the regular expression
	return regex.MatchString(input)
}

func HashPassword(password string) (string, error) {
	// Generate a salt with a cost factor of 12 (adjust this based on your security requirements)
	salt, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return "", err
	}
	return string(salt), nil
}
