package dtos

type NativeUserRegistrationDTO struct {
	GivenName  string `json:"givenName" validate:"required,min=2,max=50"`
	FamilyName string `json:"familyName" validate:"required,min=2,max=50"`
	Email      string `json:"email" validate:"email,required"`
	Password   string `json:"password" validate:"required"`
}
