package model

import (
	"github.com/google/uuid"
)

type JWTClaims struct {
	Sub uuid.UUID `json:"sub"`
	Exp int64     `json:"exp"`
	Iat int64     `json:"iat"`
}

type RegisterInput struct {
	Name     string
	Phone    string
	Email    string
	Password string
}

type RegisterOutput struct {
	Id uuid.UUID
}

type LoginInput struct {
	Phone    string
	Email    string
	Password string
}

type LoginOutput struct {
	JWT string
}
