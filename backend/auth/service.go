package auth

import (
	"errors"
	"log"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

type Service interface {
	GenerateToken(userID uuid.UUID) (string, error)
	ValidateToken(encodedToken string) (*jwt.Token, error)
}

type jwtService struct {
	secretKey []byte
}

func NewService() *jwtService {
	err := godotenv.Load("../.env")
	if err != nil {
		// log.Fatal("Error loading .env file")
		log.Fatal("Error loading .env file:", err)
	}

	secret := os.Getenv("SECRET_KEY")
	if secret == "" {
		log.Fatal("SECRET_KEY is not set in environment variables")
	}

	return &jwtService{
		secretKey: []byte(secret),
	}
}

func (s *jwtService) GenerateToken(userID uuid.UUID) (string, error) {

	var secretKey = []byte(s.secretKey)
	
	claim := jwt.MapClaims{}
	claim["user_id"] = userID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}

func (s *jwtService) ValidateToken(encodedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("Invalid token")
		}

		return []byte(s.secretKey), nil
	})

	if err != nil {
		return token, err
	}

	return token, nil
}