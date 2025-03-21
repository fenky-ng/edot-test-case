package model

import (
	"fmt"

	"github.com/google/go-cmp/cmp"
	"github.com/google/uuid"
	"go.uber.org/mock/gomock"
)

type GetProfileInput struct {
	Id uuid.UUID
}

type GetProfileOutput struct {
	Id    uuid.UUID
	Name  string
	Phone string
	Email string
}

type GetUserInput struct {
	Id    uuid.UUID
	Phone string
	Email string
}

type GetUserOutput struct {
	Id             uuid.UUID
	Name           string
	Phone          string
	Email          string
	HashedPassword string
	DeletedAt      int64
}

type InsertUserInput struct {
	Id             uuid.UUID
	Name           string
	Phone          string
	Email          string
	HashedPassword string
	CreatedAt      int64
	CreatedBy      string
}

func (expectedInput InsertUserInput) Matcher() gomock.Matcher {
	return gomock.Cond(func(x any) bool {
		actualInput := x.(InsertUserInput)

		// set zero value for ignored attributes
		expectedInput.Id, actualInput.Id = uuid.Nil, uuid.Nil
		expectedInput.CreatedAt, actualInput.CreatedAt = 0, 0

		diff := cmp.Diff(expectedInput, actualInput)
		if diff != "" {
			fmt.Printf("[InsertUserInputMatcher] DEBUG input mismatch (-want +got):\n%s\n", diff)
		}

		return diff == ""
	})
}

type InsertUserOutput struct {
	Id uuid.UUID
}
