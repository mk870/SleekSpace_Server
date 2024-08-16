package utilities

import (
	managerDtos "SleekSpace/dtos/manager"
	userDtos "SleekSpace/dtos/user"
	managerModels "SleekSpace/models/manager"
	userModels "SleekSpace/models/user"

	// "log"

	"math/rand"
	"os"
	"strconv"
	"time"
	// "github.com/joho/godotenv"
)

func GenerateVerificationCode() int {
	rand.Seed(time.Now().UnixNano())
	min := 100000
	max := 999999
	randomInt := rand.Intn(max-min) + min
	return randomInt
}

func ConvertIntToString(number int) string {
	return strconv.Itoa(number)
}

func ConvertStringToInt(number string) int {
	i, _ := strconv.Atoi(number)
	return i
}

func GenerateVerificationGracePeriod() time.Time {
	return time.Now().Add(time.Minute * 15)
}

type EnvVariables struct {
	DatabaseDetails     string
	Email               string
	EmailPassword       string
	TokenSecret         string
	LocationIQToken     string
	SupabaseApiKey      string
	SupabaseReferenceId string
}

func GetEnvVariables() EnvVariables {
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatalf("Error loading .env file: %s", err)
	// }
	databaseDetails := os.Getenv("DATABASE_DETAILS")
	email := os.Getenv("EMAIL")
	emailPassword := os.Getenv("EMAIL_PASSWORD")
	tokenSecret := os.Getenv("ACCESS_TOKEN_SECRET")
	locationIQToken := os.Getenv("LOCATION_IQ_ACCESS_TOKEN")
	return EnvVariables{
		DatabaseDetails: databaseDetails,
		Email:           email,
		EmailPassword:   emailPassword,
		TokenSecret:     tokenSecret,
		LocationIQToken: locationIQToken,
	}
}

func processedContactNumbers(contactNumbers []userModels.ContactNumber) []userDtos.ContactNumberDTO {
	contacts := []userDtos.ContactNumberDTO{}
	for i := 0; i < len(contactNumbers); i++ {
		contact := userDtos.ContactNumberDTO{
			Id:           contactNumbers[i].Id,
			CountryAbbrv: contactNumbers[i].CountryAbbrv,
			CountryCode:  contactNumbers[i].CountryCode,
			Number:       contactNumbers[i].Number,
			Type:         contactNumbers[i].Type,
			UserId:       contactNumbers[i].UserId,
		}
		contacts = append(contacts, contact)
	}
	return contacts
}

func processedManagerContactNumbers(contactNumbers []managerModels.ManagerContactNumber) []managerDtos.ManagerContactNumberDTO {
	contacts := []managerDtos.ManagerContactNumberDTO{}
	for i := 0; i < len(contactNumbers); i++ {
		contact := managerDtos.ManagerContactNumberDTO{
			Id:           contactNumbers[i].Id,
			CountryAbbrv: contactNumbers[i].CountryAbbrv,
			CountryCode:  contactNumbers[i].CountryCode,
			Number:       contactNumbers[i].Number,
			Type:         contactNumbers[i].Type,
			ManagerId:    contactNumbers[i].ManagerId,
		}
		contacts = append(contacts, contact)
	}
	return contacts
}

func UserResponseMapper(user *userModels.User, accessToken string) userDtos.UserResponseDTO {
	return userDtos.UserResponseDTO{
		Id:             user.Id,
		Email:          user.Email,
		ContactNumbers: processedContactNumbers(user.ContactNumbers),
		FamilyName:     user.FamilyName,
		GivenName:      user.GivenName,
		AccessToken:    user.AccessToken,
		ProfilePicture: userDtos.UserProfilePictureResponseAndUpdateDTO{
			Id:          user.ProfilePicture.Id,
			UserId:      user.ProfilePicture.UserId,
			Uri:         user.ProfilePicture.Uri,
			Name:        user.ProfilePicture.Name,
			FullPath:    user.ProfilePicture.FullPath,
			FileType:    user.ProfilePicture.FileType,
			Size:        user.ProfilePicture.Size,
			ContentType: user.ProfilePicture.ContentType,
		},
		Location: userDtos.LocationDTO{
			UserId:      user.Location.UserId,
			Lat:         user.Location.Lat,
			Lon:         user.Location.Lon,
			City:        user.Location.City,
			County:      user.Location.County,
			Country:     user.Location.Country,
			CountryCode: user.Location.CountryCode,
			Surburb:     user.Location.Surburb,
			Id:          user.Location.Id,
			Boundingbox: user.Location.Boundingbox,
			DisplayName: user.Location.DisplayName,
			Province:    user.Location.Province,
		},
	}
}
func ManagerResponse(manager *managerModels.Manager) managerDtos.ManagerResponseDTO {
	return managerDtos.ManagerResponseDTO{
		Id:     manager.Id,
		UserId: manager.UserId,
		Email:  manager.Email,
		Name:   manager.Name,
		ProfilePicture: managerDtos.ManagerProfilePictureResponseAndUpdateDTO{
			Id:          manager.ProfilePicture.Id,
			ManagerId:   manager.ProfilePicture.ManagerId,
			Uri:         manager.ProfilePicture.Uri,
			Name:        manager.ProfilePicture.Name,
			FullPath:    manager.ProfilePicture.FullPath,
			FileType:    manager.ProfilePicture.FileType,
			Size:        manager.ProfilePicture.Size,
			ContentType: manager.ProfilePicture.ContentType,
		},
		Contacts: processedManagerContactNumbers(manager.ManagerContactNumbers),
	}
}

func AddManagerIdToContacts(contacts []managerModels.ManagerContactNumber, managerId int) []managerModels.ManagerContactNumber {
	newContacts := []managerModels.ManagerContactNumber{}
	for i := 0; i < len(contacts); i++ {
		contact := managerModels.ManagerContactNumber{
			Id:           contacts[i].Id,
			CountryAbbrv: contacts[i].CountryAbbrv,
			CountryCode:  contacts[i].CountryCode,
			Number:       contacts[i].Number,
			Type:         contacts[i].Type,
			ManagerId:    managerId,
		}
		newContacts = append(newContacts, contact)
	}
	return newContacts
}
