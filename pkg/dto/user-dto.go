package dto

import (
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type UserDto struct {
	gorm.Model
	Name    string     `json:"name" validate:"required"`
	Phone   string     `json:"phone" validate:"required"`
	Address AddressDto `json:"address" validate:"required"`
}

var validate = validator.New()

func ValidateUserDto(user UserDto) error {
	if err := validate.Struct(user); err != nil {
		return err
	}

	//if err := validateAddress(user.Address); err != nil {
	//	return err
	//}
	return nil
}
