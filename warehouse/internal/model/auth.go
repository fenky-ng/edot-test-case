package model

import (
	"github.com/google/uuid"
)

type GetUserProfileInput struct {
	JWT string
}

type GetUserProfileOutput struct {
	HttpCode int
	Error    string
	Id       uuid.UUID
}

type HttpGetUserProfileResponse struct {
	Error string    `json:"error"`
	Id    uuid.UUID `json:"id"`
}
