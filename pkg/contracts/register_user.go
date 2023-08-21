package contracts

import (
	"github.com/go-playground/validator/v10"
)

type RegisterUserRequest struct {
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Email     string `json:"email_id" validate:"required"`
	Mobile    string `json:"mobile_number"`
	Password  string `json:"password" validate:"required"`
	ClientID  string `json:"client_id" validate:"required"`
}

func ValidateUserInput(user RegisterUserRequest) error {
	var validate *validator.Validate
	validate = validator.New()
	return validate.Struct(user)
}
