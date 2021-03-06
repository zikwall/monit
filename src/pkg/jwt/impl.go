package jwt

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/zikwall/monit/src/pkg/exceptions"
)

type Claims struct {
	UUID int64 `json:"uuid"`
	jwt.StandardClaims
}

const (
	AuthHeaderName = "Authorization"
	AuthTokenType  = "Bearer"
)

func AuthHeader(header string) (string, bool) {
	parts := strings.Split(header, " ")
	if len(parts) != 2 || !strings.EqualFold(parts[0], AuthTokenType) {
		return "", false
	}

	return parts[1], true
}

func CreateJwtToken(claims *Claims, duration int64, privateBytes []byte) (string, error) {
	key, err := jwt.ParseRSAPrivateKeyFromPEM(privateBytes)

	if err != nil {
		return "", exceptions.ThrowPublicError(err)
	}

	claims.StandardClaims = jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Duration(duration) * time.Second).Unix(),
		Issuer:    "monit",
	}

	withClaimsToken := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	return withClaimsToken.SignedString(key)
}

func VerifyJwtToken(token string, publicKeyBytes []byte) (*Claims, error) {
	withClaimsToken, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		key, err := jwt.ParseRSAPublicKeyFromPEM(publicKeyBytes)
		if err != nil {
			return nil, exceptions.ThrowPublicError(err)
		}

		if key == nil || key.N == nil {
			return nil, exceptions.ThrowPublicError(errors.New("JWT token is not defined"))
		}

		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, exceptions.ThrowPublicError(fmt.Errorf("unexpected signing method: %v", token.Header["alg"]))
		}

		return key, nil
	})

	if err != nil {
		return nil, exceptions.ThrowPublicError(err)
	}

	if claims, ok := withClaimsToken.Claims.(*Claims); ok && withClaimsToken.Valid {
		return claims, nil
	}

	return nil, exceptions.ThrowPublicError(errors.New("failed to get the source data from the JWT token"))
}
