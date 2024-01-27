package request

import (
	"errors"
	"fmt"
	"github.com/SawitProRecruitment/UserService/generated"
)

func ValidateUpdatePayload(payload *generated.UpdatePayload) error {
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

	if payload.FullName == "" {
		return errors.New("full_name cannot be empty")
	}

	fullnameLen := len(payload.FullName)
	if fullnameLen < FULLNAME_RULES_MIN_LENGTH || fullnameLen > FULLNAME_RULES_MAX_LENGTH {
		return errors.New(fmt.Sprintf("full_name should have minimum %d or maximum %d character", FULLNAME_RULES_MIN_LENGTH, FULLNAME_RULES_MAX_LENGTH))
	}

	return nil
}
