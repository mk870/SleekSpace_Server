package utilities

import (
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

func GenerateVerificationGracePeriod() time.Time {
	return time.Now().Add(time.Minute * 15)
}

type EnvVariables struct {
	DatabaseDetails string
	Email           string
	EmailPassword   string
	TokenSecret     string
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
	return EnvVariables{
		DatabaseDetails: databaseDetails,
		Email:           email,
		EmailPassword:   emailPassword,
		TokenSecret:     tokenSecret,
	}
}
