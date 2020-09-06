package auth

import (
	"context"
	"strings"

	"golang.org/x/oauth2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

var (
	errMissingMetadata = status.Errorf(codes.InvalidArgument, "missing metadata")
	errInvalidToken    = status.Errorf(codes.Unauthenticated, "invalid token")
)

type Authentication struct {
	Token string
}

func (a *Authentication) FetchToken() *oauth2.Token {
	return &oauth2.Token{
		AccessToken: a.Token,
	}
}

func (a *Authentication) Auth(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errMissingMetadata
	}

	if !valid(md["authorization"], a.Token) {
		return nil, errInvalidToken
	}

	return handler(ctx, req)
}

func valid(authorization []string, token string) bool {
	if len(authorization) < 1 {
		return false
	}

	auth := strings.TrimPrefix(authorization[0], "Bearer ")
	return auth == token
}
