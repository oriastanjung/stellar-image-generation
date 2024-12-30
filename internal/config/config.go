package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	IMAGE_GENERATION_MODEL          string
	IMAGE_API_GENERATION_URL        string
	IMAGE_API_GENERATION_ORIGIN     string
	IMAGE_API_GENERATION_VALIDATED  string
	IMAGE_API_GENERATION_MESSAGE_ID string
	Port                            string
	DatabaseURL                     string
	JWTSecretKey                    string
	BcryptSalt                      int
	GoogleAuthClientID              string
	GoogleAuthClientSecret          string
	GoogleAuthRedirectURL           string
	GoogleOAuthStateString          string
	AESSecretKey                    string
	GmailEmail                      string
	GmailPassword                   string
	EmailVerificationLink           string
	EmailForgetPasswordFrontendLink string
}

func LoadEnv() *Config {
	// load .env file
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	// Konversi SALT_KEY dari string ke int
	saltKey, err := strconv.Atoi(getEnv("SALT_KEY", "12")) // Default 12
	if err != nil {
		log.Fatalf("Invalid SALT_KEY value: %v", err)
	}
	config := &Config{
		IMAGE_GENERATION_MODEL:          getEnv("IMAGE_GENERATION_MODEL", ""),
		IMAGE_API_GENERATION_URL:        getEnv("IMAGE_API_GENERATION_URL", ""),
		IMAGE_API_GENERATION_ORIGIN:     getEnv("IMAGE_API_GENERATION_ORIGIN", ""),
		IMAGE_API_GENERATION_VALIDATED:  getEnv("IMAGE_API_GENERATION_VALIDATED", ""),
		IMAGE_API_GENERATION_MESSAGE_ID: getEnv("IMAGE_API_GENERATION_MESSAGE_ID", ""),
		Port:                            getEnv("PORT", "2701"), // defaultnya 3000,
		DatabaseURL:                     getEnv("DATABASE_URL", ""),
		JWTSecretKey:                    getEnv("JWT_SECRET_KEY", ""),
		AESSecretKey:                    getEnv("AES_SECRET_KEY", ""),
		GoogleAuthClientID:              getEnv("GOOGLE_AUTH_CLIENT_ID", ""),
		GoogleAuthClientSecret:          getEnv("GOOGLE_AUTH_CLIENT_SECRET", ""),
		GoogleAuthRedirectURL:           getEnv("GOOGLE_AUTH_REDIRECT_URL", ""),
		GoogleOAuthStateString:          getEnv("GOOGLE_OAUTH_STATE_STRING", ""),
		GmailEmail:                      getEnv("GMAIL_EMAIL", ""),
		GmailPassword:                   getEnv("GMAIL_PASSWORD", ""),
		EmailVerificationLink:           getEnv("EMAIL_VERIFICATION_LINK", ""),
		EmailForgetPasswordFrontendLink: getEnv("EMAIL_FORGET_PASSWORD_FRONTEND_LINK", ""),
		BcryptSalt:                      saltKey,
	}

	return config
}

func getEnv(key, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		log.Printf("Environment variable %s not set, defaulting to %s", key, fallback)
		value = fallback
	}
	return value

}
