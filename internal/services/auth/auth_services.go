package services

import (
	"context"

	"github.com/oriastanjung/stellar/internal/entities"
	usecase "github.com/oriastanjung/stellar/internal/usecase/auth"
)

type AuthService interface {
	RegisterAdmin(ctx context.Context, user *entities.User, salt int) error
	LoginAdmin(ctx context.Context, user *entities.User) (string, error)
	RegisterUser(ctx context.Context, user *entities.User, salt int) error
	LoginUser(ctx context.Context, user *entities.User) (string, error)
	VerifyUser(ctx context.Context, token string) error
	RequestForgetPassword(ctx context.Context, email string) error
	ResetPasswordByToken(ctx context.Context, token string, password string) error
	LoginUserViaGoogle(ctx context.Context) (string, error)
	LoginUserViaGoogleCallback(ctx context.Context, email string, username string, pictureUrl string) (string, error)
}

type authService struct {
	authUseCase usecase.AuthUseCase
}

func NewAuthService(authUseCase usecase.AuthUseCase) AuthService {
	return &authService{
		authUseCase: authUseCase,
	}
}

func (service *authService) RegisterAdmin(ctx context.Context, user *entities.User, salt int) error {
	return service.authUseCase.RegisterAdmin(user, salt)
}

func (service *authService) LoginAdmin(ctx context.Context, user *entities.User) (string, error) {
	return service.authUseCase.LoginAdmin(user)
}

func (service *authService) RegisterUser(ctx context.Context, user *entities.User, salt int) error {
	return service.authUseCase.RegisterUser(user, salt)
}

func (service *authService) LoginUser(ctx context.Context, user *entities.User) (string, error) {
	return service.authUseCase.LoginUser(user)
}

func (service *authService) VerifyUser(ctx context.Context, token string) error {
	return service.authUseCase.VerifyUser(token)
}
func (service *authService) RequestForgetPassword(ctx context.Context, email string) error {
	return service.authUseCase.RequestForgetPassword(email)
}
func (service *authService) ResetPasswordByToken(ctx context.Context, token string, password string) error {
	return service.authUseCase.ResetPasswordByToken(token, password)
}

func (service *authService) LoginUserViaGoogle(ctx context.Context) (string, error) {
	return service.authUseCase.LoginUserViaGoogle(ctx)
}

func (service *authService) LoginUserViaGoogleCallback(ctx context.Context, email string, username string, pictureUrl string) (string, error) {
	return service.authUseCase.LoginUserViaGoogleCallback(ctx, email, username, pictureUrl)
}
