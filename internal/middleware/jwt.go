package middleware

import (
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func RequireJWT(next echo.HandlerFunc) echo.HandlerFunc {
	secret := []byte(os.Getenv("JWT_SECRET"))
	return func(c echo.Context) error {
		auth := c.Request().Header.Get("Authorization")
		if auth == "" {
			return c.JSON(http.StatusUnauthorized, map[string]string{"message": "missing authorization header"})
		}
		tokenStr := auth
		// Accept both raw token or "Bearer <token>"
		if len(auth) > 7 && auth[:7] == "Bearer " {
			tokenStr = auth[7:]
		}

		tok, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) { return secret, nil })
		if err != nil || !tok.Valid {
			return c.JSON(http.StatusUnauthorized, map[string]string{"message": "invalid token"})
		}
		claims := tok.Claims.(jwt.MapClaims)
		idf, _ := claims["sub"].(float64)
		c.Set("user_id", int(idf))
		return next(c)
	}
}
