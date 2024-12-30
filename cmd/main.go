package main

import (
	"log"
	"net"

	"github.com/oriastanjung/stellar/internal/config"
	"github.com/oriastanjung/stellar/internal/database"
	serverAuth "github.com/oriastanjung/stellar/internal/grpc/auth"
	"github.com/oriastanjung/stellar/internal/middleware"
	repositoryAuth "github.com/oriastanjung/stellar/internal/repository/auth"
	servicesAuth "github.com/oriastanjung/stellar/internal/services/auth"
	usecaseAuth "github.com/oriastanjung/stellar/internal/usecase/auth"
	pbAuth "github.com/oriastanjung/stellar/proto/auth"

	serverImage "github.com/oriastanjung/stellar/internal/grpc/image"
	servicesImage "github.com/oriastanjung/stellar/internal/services/image"
	usecaseImage "github.com/oriastanjung/stellar/internal/usecase/image"
	pbImage "github.com/oriastanjung/stellar/proto/image"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
)

func main() {
	// load env
	config := config.LoadEnv()
	var port string = config.Port
	var addr string = "0.0.0.0:" + port

	// load DB
	database.InitDB()

	// Handle graceful shutdown
	database.GracefulShutdown()

	// auth service
	authRepository := repositoryAuth.NewAuthRepository(database.DB)
	authUseCase := usecaseAuth.NewAuthUseCase(authRepository)
	authService := servicesAuth.NewAuthService(authUseCase)
	authServer := serverAuth.NewAuthServer(authService, config.BcryptSalt)
	// end auth service

	//image service
	imageUseCase := usecaseImage.NewImageUseCase()
	imageService := servicesImage.NewImageService(imageUseCase)
	imageServer := serverImage.NewImageServer(imageService)
	//end image service

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Printf("Error on Listening %v\n", err)
	}

	defer listener.Close()
	log.Printf("listening on %s\n", addr)

	// set ssl certificate
	options := []grpc.ServerOption{}
	tls := true

	if tls {
		certFile := "ssl/server.crt"
		keyFile := "ssl/server.pem"
		creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)
		if err != nil {
			log.Fatalf("Failed login certificate %v", err)
		}
		options = append(options, grpc.Creds(creds))
	}

	// register middleware
	options = append(options, grpc.UnaryInterceptor(middleware.TokenValidationUnaryInterceptor))
	serverInstance := grpc.NewServer(options...)

	pbAuth.RegisterAuthServiceRoutesServer(serverInstance, authServer)
	pbImage.RegisterImageServiceServer(serverInstance, imageServer)
	// pbFinance.RegisterFinanceRoutesServiceServer(serverInstance, fincanceServer)
	// pbBusiness.RegisterBusinessRoutesServiceServer(serverInstance, businessServer)

	reflection.Register(serverInstance)
	if err = serverInstance.Serve(listener); err != nil {
		log.Fatalf("Failed on Serve %v\n", err)
	}
}
