package handlers

// import (
// 	"github.com/labstack/echo"
// )

// func LenderLoanRequestList(c echo.Context) error {
// 	defer c.Request().Body.Close()

// 	user := c.Get("user")
// 	token := user.(*jwt.Token)
// 	claims := token.Claims.(jwt.MapClaims)

// 	lenderModel := models.Bank{}

// 	lenderID, _ := strconv.Atoi(claims["jti"].(string))
// }
