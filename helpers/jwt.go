package helpers

import (
	"errors"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func GenerateToken(id uint, email string) string {
	var claims = jwt.MapClaims{
		"id":    id,
		"email": email,
	}

	var parseToken = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	var signedToken, _ = parseToken.SignedString([]byte(os.Getenv("SECRET_KEY")))

	return signedToken
}

func VerifyToken(ctx *gin.Context) (any, error) {
	var headerToken = ctx.GetHeader("Authorization")
	var bearer = strings.HasPrefix(headerToken, "Bearer ")
	var err = errors.New("sign in to proceed")

	if !bearer {
		return nil, err
	}

	var stringToken = strings.Split(headerToken, " ")[1]
	var token, _ = jwt.Parse(stringToken, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, err
		}

		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	var claims, ok = token.Claims.(jwt.MapClaims)
	if !ok && !token.Valid {
		return nil, err
	}

	return claims, nil
}
