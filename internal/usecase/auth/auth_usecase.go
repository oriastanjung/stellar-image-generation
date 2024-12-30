package usecase

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/oriastanjung/stellar/internal/config"
	"github.com/oriastanjung/stellar/internal/entities"
	repository "github.com/oriastanjung/stellar/internal/repository/auth"
	"github.com/oriastanjung/stellar/internal/utils"
	"github.com/oriastanjung/stellar/internal/utils/smtp"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type AuthUseCase interface {
	RegisterAdmin(user *entities.User, passwordSalt int) error
	LoginAdmin(user *entities.User) (string, error)
	RegisterUser(user *entities.User, passwordSalt int) error
	LoginUser(user *entities.User) (string, error)
	VerifyUser(token string) error
	RequestForgetPassword(token string) error
	ResetPasswordByToken(token string, password string) error
	LoginUserViaGoogle(ctx context.Context) (string, error)
	LoginUserViaGoogleCallback(ctx context.Context, email string, username string, pictureUrl string) (string, error)
}

type authUseCase struct {
	authRepo repository.AuthRepository
}

func NewAuthUseCase(authRepo repository.AuthRepository) AuthUseCase {
	return &authUseCase{
		authRepo: authRepo,
	}
}

func (usecase *authUseCase) RegisterAdmin(user *entities.User, passwordSalt int) error {
	// 1. Generate ID unik untuk pengguna
	user.ID = utils.GenerateIDbyKSUID()

	// 2. Set peran pengguna sebagai admin
	user.Role = string(entities.AdminRole)
	user.IsVerified = true

	// 3. Hash password untuk keamanan
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), passwordSalt)
	if err != nil {
		return status.Errorf(codes.Internal, fmt.Sprintf("Error hashing password: %v", err))
	}
	user.Password = string(hashedPassword)

	// 4. Simpan data pengguna ke database melalui repository
	return usecase.authRepo.RegisterAdmin(user)
}

func (usecase *authUseCase) LoginAdmin(user *entities.User) (string, error) {
	dbUser := &entities.User{}
	dbUser.Email = user.Email
	err := usecase.authRepo.LoginAdmin(dbUser)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", status.Errorf(codes.NotFound, "User Not Found")
		}
		return "", err
	}

	if dbUser.IsVerified == false {
		return "", status.Errorf(codes.Unauthenticated, "User Not Verified")
	}

	if dbUser.Role != string(entities.AdminRole) {
		return "", errors.New("User Not Admin")
	}

	err = bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password))
	if err != nil {
		return "", status.Errorf(codes.Unauthenticated, "Invalid Password")
	}

	token, err := utils.GenerateTokenJWT(*dbUser)
	if err != nil {
		return "", status.Errorf(codes.Internal, fmt.Sprintf("Error : %v", err))
	}
	return token, nil
}
func (usecase *authUseCase) RegisterUser(user *entities.User, passwordSalt int) error {
	cfg := config.LoadEnv()
	// 1. Generate ID unik untuk pengguna
	user.ID = utils.GenerateIDbyKSUID()

	// 2. Set peran pengguna sebagai User
	user.Role = string(entities.UserRole)
	user.IsVerified = false
	verificationToken := uuid.New().String()
	user.VerificationToken = verificationToken

	// 3. Hash password untuk keamanan
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), passwordSalt)
	if err != nil {
		return status.Errorf(codes.Internal, fmt.Sprintf("Error hashing password: %v", err))
	}
	user.Password = string(hashedPassword)

	// 3. Kirim email verifikasi// Send verification email asynchronously
	go func() {
		err := smtp.SendEmailVerification(user.Email, verificationToken, cfg.EmailVerificationLink)
		if err != nil {
			log.Printf("Error sending verification email to %s: %v", user.Email, err)
		}
	}()

	// 4. Simpan data pengguna ke database melalui repository
	return usecase.authRepo.RegisterUser(user)
}

func (usecase *authUseCase) LoginUser(user *entities.User) (string, error) {
	dbUser := &entities.User{}
	dbUser.Email = user.Email
	err := usecase.authRepo.LoginUser(dbUser)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", status.Errorf(codes.NotFound, "User Not Found")
		}
		return "", err
	}

	if dbUser.IsVerified == false {
		return "", status.Errorf(codes.Unauthenticated, "User Not Verified")
	}

	if dbUser.Role != string(entities.UserRole) {
		return "", status.Errorf(codes.Unauthenticated, ("Account Not User Role"))
	}

	err = bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password))
	if err != nil {
		return "", status.Errorf(codes.Unauthenticated, "Invalid Password")
	}

	token, err := utils.GenerateTokenJWT(*dbUser)
	if err != nil {
		return "", status.Errorf(codes.Internal, fmt.Sprintf("Error : %v", err))
	}
	return token, nil
}
func (usecase *authUseCase) VerifyUser(token string) error {
	return usecase.authRepo.VerifyUser(token)
}
func (usecase *authUseCase) RequestForgetPassword(email string) error {
	cfg := config.LoadEnv()
	user, err := usecase.authRepo.FindUserByEmail(email)
	if err != nil {
		return status.Errorf(codes.NotFound, "User Not Found")
	}
	forgetPasswordToken := uuid.New().String()
	user.ForgetPasswordToken = forgetPasswordToken
	err = usecase.authRepo.UpdateUserByEmail(email, user)
	if err != nil {
		return status.Errorf(codes.Internal, fmt.Sprintf("Error saving user: %v", err))
	}
	go func() {
		err := smtp.SendEmailForgetPassword(email, forgetPasswordToken, cfg.EmailForgetPasswordFrontendLink)
		if err != nil {
			log.Printf("Error sending forget password email to %s: %v", email, err)
		}
	}()

	return nil
}

func (usecase *authUseCase) ResetPasswordByToken(token string, password string) error {
	cfg := config.LoadEnv()
	user, err := usecase.authRepo.FindOneUserByKey("forget_password_token", token)
	if err != nil {
		return status.Errorf(codes.NotFound, "User Not Found")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), cfg.BcryptSalt)
	if err != nil {
		return status.Errorf(codes.Internal, fmt.Sprintf("Error hashing password: %v", err))
	}
	user.Password = string(hashedPassword)
	err = usecase.authRepo.UpdateUserByEmail(user.Email, user)
	if err != nil {
		return status.Errorf(codes.Internal, fmt.Sprintf("Error saving user: %v", err))
	}
	return nil
}

func (usecase *authUseCase) LoginUserViaGoogle(ctx context.Context) (string, error) {
	cfg := config.LoadEnv()
	googleOauthConfig := utils.GetGoogleOAuthConfig()
	url := googleOauthConfig.AuthCodeURL(cfg.GoogleOAuthStateString)
	return string(url), nil
}

func (usecase *authUseCase) LoginUserViaGoogleCallback(ctx context.Context, email string, username string, pictureUrl string) (string, error) {
	cfg := config.LoadEnv()
	user, err := usecase.authRepo.FindUserByEmail(email)
	if err != nil {
		var newUser entities.User
		newUser.ID = utils.GenerateIDbyKSUID()
		newUser.Email = email
		newUser.Username = username
		newUser.ProfilePictureUrl = pictureUrl
		newUser.IsVerified = true
		newUser.Role = string(entities.UserRole)
		hashedPassword, errPassword := bcrypt.GenerateFromPassword([]byte("email:"+email), cfg.BcryptSalt)
		if errPassword != nil {
			return "", status.Errorf(codes.Internal, fmt.Sprintf("Error hashing password: %v", errPassword))
		}
		newUser.Password = string(hashedPassword)

		errRegister := usecase.authRepo.RegisterUser(&newUser)
		if errRegister != nil {
			return "", status.Errorf(codes.Internal, fmt.Sprintf("Error : %v", errRegister))
		}

		token, err := utils.GenerateTokenJWT(newUser)
		if err != nil {
			return "", status.Errorf(codes.Internal, fmt.Sprintf("Error : %v", err))
		}
		return token, nil
	}
	token, err := utils.GenerateTokenJWT(*user)
	if err != nil {
		return "", status.Errorf(codes.Internal, fmt.Sprintf("Error : %v", err))
	}
	return token, nil
}
