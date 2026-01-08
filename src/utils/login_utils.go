package utils

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

var passwordRegex = `^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[!@#$%^&*(),.?":{}|<>]).{6,}$`

func PasswordValidator(fl validator.FieldLevel) bool {
	password := fl.Field().String()

	re := regexp.MustCompile(passwordRegex)

	return re.MatchString(password)
}
