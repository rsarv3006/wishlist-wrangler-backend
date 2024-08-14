package dto

import "github.com/google/uuid"

type UpdateUserDto struct {
	UserId      uuid.UUID `json:"userId"`
	DisplayName string    `json:"displayName"`
}
