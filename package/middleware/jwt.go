package middleware
import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)
func CreateToken(operatorID int, username string) (string, error) {
	claims := jwt.MapClaims{}
	claims["operatorID"] = operatorID
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(constants.SECRET_JWT))
}