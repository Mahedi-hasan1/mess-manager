package util

import (
	"github.com/go-playground/validator/v10"
	"regexp"

)

func GetCustomValidationMessage(err error) string {
	
	validationErrors := err.(validator.ValidationErrors)

	errorMessages := map[string]map[string]string{
		"Password": {
			"required": "password is required",
			"min":      "password is less than 4 characters",
		},
		"Username": {
			"required": "username is required",
			"min":      "username must be at least 3 characters",
			"max":      "username cannot exceed 50 characters",
		},
		
	}   
	for _, fieldError := range validationErrors {
		field := fieldError.Field()
		tag := fieldError.Tag()

		if fieldMessages, exists := errorMessages[field]; exists {
			if message, exists := fieldMessages[tag]; exists {
				return message
			}
		}
	}

	return err.Error() // fallback to default message
}

func IsValidPhoneNumber(phone string) bool {
	// accepts exactly 11 or 13 digits
	matched, _ := regexp.MatchString(`^([0-9]{11}|[0-9]{13})$`, phone)
	return matched
}