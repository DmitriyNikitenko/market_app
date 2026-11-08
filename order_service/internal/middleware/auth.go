package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(secretkey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, "not authorized")
			c.Abort()
			return
		}

		stringData := strings.Split(authHeader, " ")
		if len(stringData) != 2 || stringData[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, "not authorized")
			c.Abort()
			return
		}

		tokenString := stringData[1]
		//метод Parse вернет нам объект токена и ошибку
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			//проверяем, что алгоритм шифрования HMAC (защита от злоумышленников)
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method")
			}
			//отдаем наш секретный ключ, превратив его в массив байт
			return []byte(secretkey), nil
		})

		//проверяем результат парсинга
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, "not authorized")
			c.Abort()
			return
		}

		//достанем мапу с данными, чтобы узнать, кто делает запрос
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, "not authorized")
			c.Abort()
			return
		}

		userIDFloat, ok := claims["user_id"].(float64)
		if !ok {
			c.JSON(http.StatusUnauthorized, "not authorized")
			c.Abort()
			return
		}

		c.Set("userID", int64(userIDFloat))
		c.Next()
	}

}
