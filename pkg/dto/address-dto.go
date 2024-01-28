package dto

type AddressDto struct {
	Village string `json:"village" validate:"required"`
	Pin     string `json:"pin" validate:"required"`
	State   string `json:"state" validate:"required"`
	Country string `json:"country" validate:"required"`
	UserId  uint   `json:"userId"`
}
