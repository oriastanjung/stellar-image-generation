package middleware

import (
	"context"

	"github.com/oriastanjung/stellar/internal/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func TokenValidationUnaryInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	// List of methods to skip token validation
	skipMethods := map[string]bool{
		"/auth.AuthServiceRoutes/SignUpAdmin":                true,
		"/auth.AuthServiceRoutes/LoginAdmin":                 true,
		"/auth.AuthServiceRoutes/SignUpUser":                 true,
		"/auth.AuthServiceRoutes/LoginUser":                  true,
		"/auth.AuthServiceRoutes/VerifyUser":                 true,
		"/auth.AuthServiceRoutes/LoginUserViaGoogle":         true,
		"/auth.AuthServiceRoutes/LoginUserViaGoogleCallback": true,
		"/auth.AuthServiceRoutes/RequestForgetPassword":      true,
		"/auth.AuthServiceRoutes/ResetPasswordByToken":       true,
	}

	// Check if the method should skip token validation
	if skipMethods[info.FullMethod] {
		// Skip token validation
		return handler(ctx, req)
	}

	// Extract metadata from the context
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "metadata not provided")
	}

	// Get the authorization header value
	values := md["authorization"]
	if len(values) == 0 {
		return nil, status.Errorf(codes.Unauthenticated, "invalid or missing token")
	}

	// Extract and clean up the token
	token := values[0]

	// Verify the token
	claims, err := utils.VerifyTokenJWT(token)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "invalid token: %v", err)
	}
	// Add claims to context for use in downstream handlers
	ctx = context.WithValue(ctx, "claims", claims)

	// Proceed to the handler
	return handler(ctx, req)
}
