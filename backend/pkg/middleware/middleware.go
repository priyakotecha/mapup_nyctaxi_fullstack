package middleware

import (
	"net/http"
	"nyctaxi_mapup/pkg/cache"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var jwtKey = []byte("your_secret_key") // Change to a secure key

type Claims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

func AuthMiddleware(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			c.Abort()
			return
		}

		claims, err := ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
			c.Abort()
			return
		}

		user, exists := cache.GetUserCredentials(claims.Username)
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid user"})
			c.Abort()
			return
		}

		for _, role := range allowedRoles {
			if user.Role == role {
				c.Set("username", claims.Username)
				c.Set("role", user.Role)
				c.Next()
				return
			}
		}

		c.JSON(http.StatusForbidden, gin.H{"message": "Forbidden"})
		c.Abort()
	}
}

func GenerateToken(username string, role string) (string, error) {
	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &Claims{
		Username: username,
		Role:     role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	cache.SetJWTToken(username, token.Raw)
	return token.SignedString(jwtKey)
}

func ValidateToken(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}
	return claims, nil
}
