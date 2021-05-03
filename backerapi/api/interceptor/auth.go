package interceptor

import (
	"context"

	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Auth(ctx context.Context) (context.Context, error) {
	token, err := grpc_auth.AuthFromMD(ctx, "bearer")
	if err != nil {
		return nil, err
	}

	if token != "hi/mi/tsu" {
		return nil, status.Errorf(codes.Unauthenticated, "invalid bearer token")
	}
	return context.WithValue(ctx, "UserName", "God"), nil
}
