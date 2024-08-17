package dtos

type UserResponseDTO struct {
	Email          string                                 `json:"email"`
	Id             int                                    `json:"id"`
	GivenName      string                                 `json:"givenName"`
	FamilyName     string                                 `json:"familyName"`
	AccessToken    string                                 `json:"accessToken"`
	ProfilePicture UserProfilePictureResponseAndUpdateDTO `json:"profilePicture"`
	Location       LocationDTO                            `json:"location"`
	ContactNumbers []ContactNumberDTO                     `json:"contactNumbers"`
}
