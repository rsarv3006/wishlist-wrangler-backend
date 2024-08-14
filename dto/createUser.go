package dto

type CreateUserDto struct {
	Email       string `json:"email"`
	DisplayName string `json:"displayName"`
}
