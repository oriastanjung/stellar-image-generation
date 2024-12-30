package utils

import (
	"context"

	"github.com/segmentio/ksuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func GetUserId(ctx context.Context) (ksuid.KSUID, error) {
	claims, ok := ctx.Value("claims").(*JWTClaims)
	if !ok {
		return ksuid.Nil, status.Errorf(codes.Unauthenticated, "unauthenticated")
	}
	userId := claims.UserId
	return userId, nil
}
