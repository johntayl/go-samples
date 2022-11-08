package main

import (
	"errors"
	"fmt"

	jwt "github.com/golang-jwt/jwt/v4"
)

const (
	JwtSecret = "secret"
)

// Struct to hold the token info
type JwtTokenInfo struct {
	UserId string `json:"user_id"`
	Role   string `json:"role"`
}

// SignToken signs a new JWT token
func SignToken(info *JwtTokenInfo) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": info.UserId,
		"role":    info.Role,
	})

	tokenString, err := token.SignedString([]byte(JwtSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// VerifyToken verifies a JWT token
func VerifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(JwtSecret), nil
	})

	return token, err
}

// ExtractTokenInfo extracts the token info from a JWT token
func ExtractTokenInfo(token *jwt.Token) (*JwtTokenInfo, error) {
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	userId, ok := claims["user_id"].(string)
	if !ok {
		return nil, errors.New("invalid user_id")
	}

	role, ok := claims["role"].(string)
	if !ok {
		return nil, errors.New("invalid role")
	}

	return &JwtTokenInfo{
		UserId: userId,
		Role:   role,
	}, nil
}

func main() {

	// Sign a new JWT token
	tokenString, err := SignToken(&JwtTokenInfo{
		UserId: "123",
		Role:   "admin",
	})

	if err != nil {
		panic(err)
	}

	fmt.Println("Token: " + tokenString)

	// Verify the token
	token, err := VerifyToken(tokenString)
	if err != nil {
		panic(err)
	}

	// Extract the token info
	info, err := ExtractTokenInfo(token)
	if err != nil {
		panic(err)
	}

	fmt.Println("Decoded token info")
	fmt.Println("User ID: " + info.UserId)
	fmt.Println("Role: " + info.Role)
}
