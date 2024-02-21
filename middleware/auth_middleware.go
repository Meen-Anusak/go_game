package middleware

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"go_game/database"
	"go_game/database/crm_model"
	"os"
	"strings"
)

func AuthMiddleWare() func(ctx *fiber.Ctx) error {
	signature := []byte(os.Getenv("SECRET_KEY"))
	return func(ctx *fiber.Ctx) error {
		auth := ctx.Get("Authorization")
		token := strings.TrimPrefix(auth, "Bearer ")
		tokenByte, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpeceted signing method: %v\n", token.Header["alg"])
			}
			return signature, nil
		})

		if err != nil {
			return ctx.Status(fiber.StatusUnauthorized).JSON("Unauthorized")

		}

		claims, ok := tokenByte.Claims.(jwt.MapClaims)
		if !ok || !tokenByte.Valid {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": "invalid token claim"})

		}

		var user crm_model.User
		db, err := database.NewDatabase()
		if err != nil {
			panic("failed to connect database")
		}

		if result := db.First(&user, "user_id", claims["jti"]); result.Error != nil {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": "invalid user_id claim"})
		}

		ctx.Locals("user", crm_model.FilterUserRecord(&user))
		return ctx.Next()
	}

}
