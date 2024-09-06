package util

import (
	"github.com/go-playground/validator/v10"
)

func Validate[T any](data T) map[string]string {
	validate := validator.New()
	err := validate.Struct(data)
	validationMap := make(map[string]string)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		for _, ve := range validationErrors {
			validationMap[ve.Field()] = ve.Tag()
		}
		return validationMap
	}
	return validationMap

}
