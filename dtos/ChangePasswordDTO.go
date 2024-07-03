package dtos

type ChangePasswordDTO struct {
	Email string `json:"email" validate:"email,required"`
}

type NewPasswordDTO struct {
	VerificationCode int    `json:"verificationCode" validate:"required,min=100000,max=999999"`
	Password         string `json:"password" validate:"required"`
	UserId           int    `json:"userId" validate:"required"`
}
