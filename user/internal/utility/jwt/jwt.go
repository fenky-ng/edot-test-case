package jwt

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"

	in_err "github.com/fenky-ng/edot-test-case/user/internal/error"
	"github.com/fenky-ng/edot-test-case/user/internal/model"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var jwtSecret []byte

func init() {
	jwtSecret = []byte(os.Getenv("JWT_SECRET"))
}

func GenerateJWT(userID uuid.UUID) (string, error) {
	claims := jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(7 * 24 * time.Hour).UnixMilli(), // Expiration time (1 week)
		"iat": time.Now().UnixMilli(),                         // Issued At
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign token with secret key
	return token.SignedString(jwtSecret)
}

func VerifyJWT(tokenString string) (model.JWTClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if token.Method != jwt.SigningMethodHS256 {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Method)
		}
		return jwtSecret, nil
	})
	if err != nil {
		return model.JWTClaims{}, errors.Join(in_err.ErrInvalidAuthToken, err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		jsonbody, err := json.Marshal(claims)
		if err != nil {
			return model.JWTClaims{}, errors.Join(in_err.ErrJWT, err)
		}

		var jc model.JWTClaims
		if err := json.Unmarshal(jsonbody, &jc); err != nil {
			return model.JWTClaims{}, errors.Join(in_err.ErrJWT, err)
		}

		return jc, nil
	}

	return model.JWTClaims{}, in_err.ErrInvalidAuthToken
}
