package auth

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

type Service interface {
	GenerateToken(userID int) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtService struct {
}

// var SECRET_KEY = []byte("RAMIDARANGSIT_th776_875")

func NewService() *jwtService {
	return &jwtService{}
}

func (s *jwtService) GenerateToken(UserID int) (string, error) {

	// buat expired time jwt
	expiredTime := time.Now().Add(5 * time.Minute)

	// ambil input id parameter untuk generate token
	claim := jwt.MapClaims{}

	claim["user_id"] = UserID
	claim["exp"] = expiredTime.Unix()

	// buat token atau generate claim token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	secKey := os.Getenv("SECRET_KEY")

	// tandatangi atau setujui token
	signedToken, err := token.SignedString([]byte(secKey))

	if err != nil {
		return signedToken, err
	}

	return signedToken, nil

}

func (s *jwtService) ValidateToken(token string) (*jwt.Token, error) {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	secKey := os.Getenv("SECRET_KEY")

	tokenJWT, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("Invalid Token")
		}

		return []byte(secKey), nil
	})

	if err != nil {
		return tokenJWT, err
	}

	return tokenJWT, nil
}
