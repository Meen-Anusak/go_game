package middleware

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"os"
	"strings"
)

func Protect() func(ctx *fiber.Ctx) error {
	signature := []byte(os.Getenv("SECRET_KEY"))
	return func(ctx *fiber.Ctx) error {
		auth := ctx.Get("Authorization")
		token := strings.TrimPrefix(auth, "Bearer ")
		_, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpeceted signing method: %v\n", token.Header["alg"])
			}
			return signature, nil
		})

		if err != nil {
			return ctx.Status(fiber.StatusUnauthorized).JSON("Unauthorized")

		}

		return ctx.Next()
	}
}
