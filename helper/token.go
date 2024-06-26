package helper

import (
	"starter-gofiber/config"
	"starter-gofiber/dto"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(user dto.UserClaims) (string, error) {
	// Create the Claims
	claims := jwt.MapClaims{
		"id":    user.ID,
		"email": user.Email,
		"role":  user.Role,
		"exp":   time.Now().Add(time.Hour * 24 * 7).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	// Generate encoded token and send it as response.
	return token.SignedString(config.PRIVATE_KEY)
}
