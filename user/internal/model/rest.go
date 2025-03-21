package model

import (
	"github.com/google/uuid"
)

type RestAPIErrorResponse struct {
	Error string `json:"error"`
}

type RestAPIRegisterRequest struct {
	Name         string `json:"name"`
	PhoneOrEmail string `json:"phoneOrEmail"`
	Password     string `json:"password"`
}

type RestAPIRegisterResponse struct {
	Id uuid.UUID `json:"id"`
}

type RestAPILoginRequest struct {
	PhoneOrEmail string `json:"phoneOrEmail"`
	Password     string `json:"password"`
}

type RestAPILoginResponse struct {
	JWT string `json:"jwt"`
}

type RestAPIGetProfileResponse struct {
	Id    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Phone string    `json:"phone,omitempty"`
	Email string    `json:"email,omitempty"`
}
