package dto

type AddressDto struct {
	Village string `json:"village"`
	Pin     string `json:"pin"`
	State   string `json:"state"`
	Country string `json:"country"`
	UserId  uint   `json:"userId"`
}
