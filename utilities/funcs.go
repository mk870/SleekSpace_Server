package utilities

import (
	"SleekSpace/dtos"
	"SleekSpace/models"

	//	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
	//"github.com/joho/godotenv"
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
	supabaseApiKey := os.Getenv("SUPABASE_REF_ID")
	supabaseReferenceId := os.Getenv("SUPABASE_APIKEY")
	return EnvVariables{
		DatabaseDetails:     databaseDetails,
		Email:               email,
		EmailPassword:       emailPassword,
		TokenSecret:         tokenSecret,
		LocationIQToken:     locationIQToken,
		SupabaseApiKey:      supabaseApiKey,
		SupabaseReferenceId: supabaseReferenceId,
	}
}

func processedContactNumbers(contactNumbers []models.ContactNumber) []dtos.ContactNumberDTO {
	contacts := []dtos.ContactNumberDTO{}
	for i := 0; i < len(contactNumbers); i++ {
		contact := dtos.ContactNumberDTO{
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

func processedManagerContactNumbers(contactNumbers []models.ManagerContactNumber) []dtos.ManagerContactNumberDTO {
	contacts := []dtos.ManagerContactNumberDTO{}
	for i := 0; i < len(contactNumbers); i++ {
		contact := dtos.ManagerContactNumberDTO{
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

func UserResponseMapper(user *models.User, accessToken string) dtos.UserResponseDTO {
	return dtos.UserResponseDTO{
		Id:             user.Id,
		Email:          user.Email,
		ContactNumbers: processedContactNumbers(user.ContactNumbers),
		Avatar:         user.Avatar,
		FamilyName:     user.FamilyName,
		GivenName:      user.GivenName,
		AccessToken:    user.AccessToken,
		Location: dtos.LocationDTO{
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
func ManagerResponse(manager *models.Manager) dtos.ManagerResponseDTO {
	return dtos.ManagerResponseDTO{
		Id:       manager.Id,
		UserId:   manager.UserId,
		Email:    manager.Email,
		Name:     manager.Name,
		Contacts: processedManagerContactNumbers(manager.ManagerContactNumbers),
	}
}

func AddManagerIdToContacts(contacts []models.ManagerContactNumber, managerId int) []models.ManagerContactNumber {
	newContacts := []models.ManagerContactNumber{}
	for i := 0; i < len(contacts); i++ {
		contact := models.ManagerContactNumber{
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
