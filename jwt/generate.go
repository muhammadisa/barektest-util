package jwt

import (
	jwtlib "github.com/golang-jwt/jwt"
	"time"
)

func (j jwt) Generate(data map[string]string) (*Data, error) {
	claims := jwtlib.MapClaims{}
	claims["exp"] = time.Now().Add(j.Exp).Unix()
	for k, v := range data {
		claims[k] = v
	}
	jwtToken := jwtlib.NewWithClaims(jwtlib.SigningMethodHS256, claims)
	token, err := jwtToken.SignedString([]byte(j.Secret))
	if err != nil {
		return nil, err
	}

	refreshToken := jwtlib.New(jwtlib.SigningMethodHS256)
	rtClaims := refreshToken.Claims.(jwtlib.MapClaims)
	rtClaims["exp"] = time.Now().Add(j.Exp * 2).Unix()
	for k, v := range data {
		rtClaims[k] = v
	}
	refresh, err := refreshToken.SignedString([]byte(j.Secret))
	if err != nil {
		return nil, err
	}

	return &Data{
		Token:        token,
		RefreshToken: refresh,
	}, nil
}
