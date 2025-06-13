package middleware

import (
	"fmt"
	"net/http"

	//"go/token"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("12345678901234567890123456789012") // 32 bytes == 256 bits == 32 chars

// CustomClaims structure
type CustomClaims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

// function to Generate jwt
func GenerateJWT(email string, expirationMinutes time.Duration) (string, error) {
	expirationTime := time.Now().Add(expirationMinutes * time.Minute)

	claims := CustomClaims{
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// checking validity and before expiration of jwt
func IsTokenValid(tokenstr string) (bool, error) {
	claims := &CustomClaims{}

	token, err := jwt.ParseWithClaims(tokenstr, claims, func(t *jwt.Token) (interface{}, error) { return jwtKey, nil })
	if err != nil {
		return false, err
	}

	if !token.Valid {
		return false, fmt.Errorf("invalid token")
	}
	if claims.ExpiresAt == nil || claims.ExpiresAt.Time.Before(time.Now()) {
		return false, fmt.Errorf("token expired at %v", claims.ExpiresAt.Time)
	}
	return true, nil
}

// ExtractEmail extracts the email from a JWT
func ExtractEmail(tokenStr string) (string, error) {
	claims := &CustomClaims{}

	_, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) { return jwtKey, nil })

	if err != nil {
		return "", err
	}

	return claims.Email, nil
}

// calling all fns
func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		tokenStr, err1 := GenerateJWT("palnutsam2002@gmail.com", 10)
		if err1 != nil {
			c.JSON(500, gin.H{"error": "Error in generating jwt"})
			return
		}
		fmt.Println("jwt-generated : ",tokenStr)
		valid, err := IsTokenValid(tokenStr)
		if err != nil || !valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			return
		}

		email, _ := ExtractEmail(tokenStr)
		c.Set("email", email)
		fmt.Println("email : ",email)
		if email != "palnutsam2002@gmail.com" {
			c.JSON(500,"Internal error")
		}
		c.Next()
	}
}
