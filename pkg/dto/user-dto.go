package dto

import "gorm.io/gorm"

type UserDto struct {
	gorm.Model
	Name    string     `json:"name" validate:"required"`
	Phone   string     `json:"phone" validate:"required"`
	Address AddressDto `json:"address" validate:"required"`
}
