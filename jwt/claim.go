package jwt

import jwtlib "github.com/golang-jwt/jwt"

func (j jwt) Claim(token *jwtlib.Token, key string) (string, error) {
	if claims, ok := token.Claims.(jwtlib.MapClaims); ok && token.Valid {
		val, claimOk := claims[key]
		if claimOk {
			casted, castOk := val.(string)
			if castOk {
				return casted, nil
			}
			return blankString, errorClaimCastingFailed
		}
		return blankString, errorClaimKeyNotFound
	}
	return blankString, errorClaimNotOK
}
