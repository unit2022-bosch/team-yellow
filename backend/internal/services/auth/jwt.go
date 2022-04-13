package auth

import (
	"github.com/golang-jwt/jwt"
	"github.com/pkg/errors"
	"github.com/unit2022-bosch/teapot/backend/internal/entity"
)

func (s *authService) createJwtToken(user *entity.User) (string, error) {
	jwtInfo := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
	})
	token, err := jwtInfo.SignedString([]byte(s.jwtSecret))
	return token, errors.Wrap(err, "failed to create jwt token")
}

func (s *authService) parseJwtToken(token string) (uint, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.jwtSecret), nil
	})
	if err != nil {
		return 0, errors.Wrap(err, "failed to parse jwt token")
	}
	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		return uint(claims["user_id"].(float64)), nil
	}
	return 0, errors.New("invalid jwt token")
}
