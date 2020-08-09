package controllers

import (
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

// login function
func Login(c echo.Context) error {
	// get submitted form value
	username := c.FormValue("username")
	password := c.FormValue("password")

	// TODO: using user table
	if username == "jon" && password == "shhh!" {
		// Create token
		token := jwt.New(jwt.SigningMethodHS256)

		// Set claims
		claims := token.Claims.(jwt.MapClaims)
		claims["name"] = "Jon Snow"
		claims["access"] = "1"
		claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

		// Generate encoded token and send it as response.
		t, err := token.SignedString([]byte("secret"))
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, map[string]string{
			"token": t,
		})
	}

	return echo.ErrUnauthorized
}

func GetToken(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	al := claims["access"].(string)
	return c.String(http.StatusOK, "Welcome "+al+name+"!")
}

func AccessMe(c echo.Context) error {
	return c.JSON(http.StatusOK, "welcome")
}
