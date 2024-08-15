package dto

import "github.com/google/uuid"

type CreateLoginRequest struct {
	UserId uuid.UUID `json:"userId"`
}
