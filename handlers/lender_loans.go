package handlers

import (
	"asira_lender/models"
	"database/sql"
	"net/http"
	"strconv"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func LenderLoanRequestList(c echo.Context) error {
	defer c.Request().Body.Close()

	user := c.Get("user")
	token := user.(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)

	lenderID, _ := strconv.Atoi(claims["jti"].(string))

	// pagination parameters
	rows, err := strconv.Atoi(c.QueryParam("rows"))
	page, err := strconv.Atoi(c.QueryParam("page"))
	orderby := c.QueryParam("orderby")
	sort := c.QueryParam("sort")

	// filters
	status := c.QueryParam("status")
	owner := c.QueryParam("owner")

	type Filter struct {
		Bank   sql.NullInt64 `json:"bank"`
		Status string        `json:"status"`
		Owner  string        `json:"owner"`
	}

	loan := models.Loan{}
	result, err := loan.PagedFilterSearch(page, rows, orderby, sort, &Filter{
		Bank: sql.NullInt64{
			Int64: int64(lenderID),
			Valid: true,
		},
		Owner:  owner,
		Status: status,
	})

	if err != nil {
		return returnInvalidResponse(http.StatusInternalServerError, err, "query result error")
	}

	return c.JSON(http.StatusOK, result)
}

func LenderLoanRequestListDetail(c echo.Context) error {
	defer c.Request().Body.Close()

	user := c.Get("user")
	token := user.(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)

	lenderID, _ := strconv.Atoi(claims["jti"].(string))

	loan_id, err := strconv.Atoi(c.Param("loan_id"))

	type Filter struct {
		Bank sql.NullInt64 `json:"bank"`
		ID   int           `json:"id"`
	}

	loan := models.Loan{}
	result, err := loan.FilterSearchSingle(&Filter{
		Bank: sql.NullInt64{
			Int64: int64(lenderID),
			Valid: true,
		},
		ID: loan_id,
	})

	if err != nil {
		return returnInvalidResponse(http.StatusInternalServerError, err, "query result error")
	}

	return c.JSON(http.StatusOK, result)
}
