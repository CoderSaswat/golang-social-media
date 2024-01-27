package dto

import "gorm.io/gorm"

type UserDto struct {
	gorm.Model
	Name    string     `json:"name"`
	Phone   string     `json:"phone"`
	Address AddressDto `json:"address"`
}
