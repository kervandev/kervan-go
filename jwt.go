package kervan

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type LicenceCheckJWT struct {
	LicenceKey string                 `json:"licence_key"`
	Data       map[string]interface{} `json:"data"`
	jwt.RegisteredClaims
}

func GenerateLicenceCheckJWT(licenceKey string, data map[string]interface{}, secret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, LicenceCheckJWT{
		LicenceKey: licenceKey,
		Data:       data,
		RegisteredClaims: jwt.RegisteredClaims{
			Audience:  jwt.ClaimStrings{"kervan"},
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Minute)),
		},
	})
	return token.SignedString([]byte(secret))
}
