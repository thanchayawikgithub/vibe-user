package jwt

import (
	"time"
	"vibe-user/internal/config"
	"vibe-user/internal/modules/user/entity"

	"github.com/golang-jwt/jwt/v5"
)

type JwtClaims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

func GenerateToken(cfg *config.Jwt, user *entity.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, JwtClaims{
		UserID: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * time.Duration(cfg.ExpiresIn))),
		},
	})

	return token.SignedString([]byte(cfg.Secret))
}
