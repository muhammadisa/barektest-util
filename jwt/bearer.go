package jwt

import (
	"errors"
	"golang.org/x/net/context"
	"google.golang.org/grpc/metadata"
	"strings"
)

func Bearer(ctx context.Context, secret string) (jwt, error) {
	var tkn jwt
	var tokenMetadata []string
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		tokenMetadata = md.Get("authorization")
	}
	if len(tokenMetadata) == 0 {
		return tkn, errors.New("no token given")
	}
	token := strings.Split(tokenMetadata[0], " ")
	if len(token) != 2 || token[0] != "Bearer" {
		return tkn, errors.New("invalid bearer token")
	}
	tkn.Bearer = token[1]
	tkn.Secret = secret
	return tkn, nil
}
