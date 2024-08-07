package dtos

type VerificationDTO struct {
	UserId           int `json:"userId" validate:"required"`
	VerificationCode int `json:"verificationCode" validate:"required,min=100000,max=999999"`
}

type CreateVerificationCodeDTO struct {
	Email string `json:"email" validate:"email,required"`
}
