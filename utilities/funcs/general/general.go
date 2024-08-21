package general

import (
	//"log"

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
	return EnvVariables{
		DatabaseDetails: databaseDetails,
		Email:           email,
		EmailPassword:   emailPassword,
		TokenSecret:     tokenSecret,
		LocationIQToken: locationIQToken,
	}
}
