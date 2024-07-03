package dtos

type LoginDTO struct {
	Email    string `json:"email" validate:"email,required"`
	Password string `json:"password" validate:"required"`
}
