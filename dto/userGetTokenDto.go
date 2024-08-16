package dto

import "github.com/google/uuid"

type UserTokenDto struct {
	Email          string    `json:"email"`
	LoginRequestId uuid.UUID `json:"loginRequestId"`
	LoginCode      string    `json:"loginCode"`
}
