package dtos

type NewPasswordDTO struct {
	Password string `json:"password" validate:"required"`
	UserId   int    `json:"userId" validate:"required"`
}
