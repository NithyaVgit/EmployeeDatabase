package validation

import (
	"errors"
	"go/employee/attendance/domain"

	"github.com/go-playground/validator/v10"
)

var (
	validates *validator.Validate
)

func InitializeValidation() {
	validates = validator.New()
}

/*
Employee Method
Arugments :  Employee Struct
Return : Error Struct
*/

func EmployeeStructValidation(request domain.EmployeeClientRequest) error {
	structErr := validates.Struct(request) // Check Struct Validation
	if structErr != nil {
		for _, walletErr := range structErr.(validator.ValidationErrors) {
			return errors.New(walletErr.Error())
		}
	}

	return nil
}
