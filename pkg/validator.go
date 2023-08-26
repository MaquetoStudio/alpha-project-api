package pkg

import "github.com/go-playground/validator/v10"

func ValidateStruct(obj interface{}) (bool, map[string]string) {
	validate := validator.New()
	err := validate.Struct(obj)
	if err != nil {
		errors := make(map[string]string)
		for _, err := range err.(validator.ValidationErrors) {
			errors[err.Field()] = err.Error()
		}
		return false, errors
	}
	return true, nil
}
