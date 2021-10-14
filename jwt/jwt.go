package jwt

import (
	"errors"
	jwtlib "github.com/golang-jwt/jwt"
	"time"
)

const (
	blankString = ``
)

var (
	errorNoTokenGiven          = errors.New("no bearer token given")
	errorUnexpectSigningMethod = errors.New("unexpected signing method")
	errorClaimNotOK            = errors.New("claim not ok or token invalid")
	errorClaimKeyNotFound      = errors.New("fail claim key or key not recognized")
	errorClaimCastingFailed    = errors.New("claim key found but cast error")
)

type JWT interface {
	Parser() (*jwtlib.Token, error)
	Claim(*jwtlib.Token, string) (string, error)
	ExtractKey(string) (string, error)
	ExtractKeys([]string) (map[string]string, error)
	Generate(map[string]string) (*Data, error)
}

type Data struct {
	RefreshToken string `json:"refresh_token"`
	Token        string `json:"token"`
}

type jwt struct {
	Bearer, Secret string
	Exp            time.Duration
}

func NewJWT(bearer, secret string) JWT {
	return &jwt{Bearer: bearer, Secret: secret}
}

func NewJWTGenerateMode(exp time.Duration, secret string) JWT {
	return &jwt{Exp: exp, Secret: secret}
}
