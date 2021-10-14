package jwt

import (
	jwtlib "github.com/golang-jwt/jwt"
)

func (j jwt) Parser() (*jwtlib.Token, error) {
	token, err := jwtlib.Parse(j.Bearer, func(token *jwtlib.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwtlib.SigningMethodHMAC); !ok {
			return nil, errorUnexpectSigningMethod
		}
		return []byte(j.Secret), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}
