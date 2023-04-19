package kervan

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type LicenceCheckJWT struct {
	LicenceKey string            `json:"licence_key"`
	IpAddress  string            `json:"ip_address"`
	Data       map[string]string `json:"data"`
	jwt.RegisteredClaims
}

type LicenceCheckResponseJWT struct {
	IsValid  bool   `json:"is_valid"`
	PlanCode string `json:"plan_code"`
	jwt.RegisteredClaims
}

func GenerateLicenceCheckJWT(licenceKey, ip_address string, data map[string]string, secret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, LicenceCheckJWT{
		LicenceKey: licenceKey,
		IpAddress:  ip_address,
		Data:       data,
		RegisteredClaims: jwt.RegisteredClaims{
			Audience:  jwt.ClaimStrings{"kervan"},
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Minute)),
		},
	})
	return token.SignedString([]byte(secret))
}

func ParseLicenceCheckResponseJWT(tokenString, secret string) (claims *LicenceCheckResponseJWT, err error) {
	if tokenString == "" {
		return nil, errors.New("tokenString is empty")
	}
	if secret == "" {
		return nil, errors.New("secret is empty")
	}
	token, err := jwt.ParseWithClaims(
		tokenString,
		&LicenceCheckResponseJWT{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		},
	)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*LicenceCheckResponseJWT)
	if !ok {
		return nil, errors.New("claims not LicenceCheckResponseJWT")
	}

	return claims, nil
}
