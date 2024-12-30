package utils

import (
	"github.com/oriastanjung/stellar/internal/config"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func GetGoogleOAuthConfig() *oauth2.Config {
	cfg := config.LoadEnv()
	return &oauth2.Config{
		RedirectURL:  cfg.GoogleAuthRedirectURL,
		ClientID:     cfg.GoogleAuthClientID,
		ClientSecret: cfg.GoogleAuthClientSecret,
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.profile", "https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}
}
