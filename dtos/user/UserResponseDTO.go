package dtos

type UserResponseDTO struct {
	Id                  int                                    `json:"id"`
	FavouriteProperties []int                                  `json:"favouriteProperties"`
	Email               string                                 `json:"email"`
	GivenName           string                                 `json:"givenName"`
	FamilyName          string                                 `json:"familyName"`
	Role                string                                 `json:"role"`
	AccessToken         string                                 `json:"accessToken"`
	ProfilePicture      UserProfilePictureResponseAndUpdateDTO `json:"profilePicture"`
	Location            LocationDTO                            `json:"location"`
	ContactNumbers      []ContactNumberDTO                     `json:"contactNumbers"`
}
