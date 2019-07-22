package migration

import (
	"asira_lender/asira"
	"asira_lender/models"
	"database/sql"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/jinzhu/gorm/dialects/postgres"
)

func Seed() {
	seeder := asira.App.DB.Begin()
	defer seeder.Commit()

	if asira.App.ENV == "development" {
		// seed lender
		services := []int{1, 2, 3, 5, 8}
		jMarshal, _ := json.Marshal(services)
		lenders := []models.Bank{
			models.Bank{
				Name:     "Bank A",
				Type:     "BDP",
				Address:  "Bank A Address",
				Province: "Province A",
				City:     "City A",
				Services: postgres.Jsonb{jMarshal},
				PIC:      "Bank A PIC",
				Phone:    "081234567890",
				Username: "Banktoib",
				Password: "password",
			},
			models.Bank{
				Name:     "Bank B",
				Type:     "BDP",
				Address:  "Bank B Address",
				Province: "Province B",
				City:     "City B",
				Services: postgres.Jsonb{jMarshal},
				PIC:      "Bank B PIC",
				Phone:    "081234567891",
				Username: "Banktoic",
				Password: "password",
			},
		}
		for _, lender := range lenders {
			lender.Create()
		}

		// seed loans
		fees := []models.LoanFee{
			models.LoanFee{
				Description: "fee 1",
				Amount:      1000,
			},
		}
		jMarshal, _ = json.Marshal(fees)
		loans := []models.Loan{
			models.Loan{
				Bank: sql.NullInt64{
					Int64: 1,
					Valid: true,
				},
				LoanAmount:       5000000,
				Installment:      8,
				LoanIntention:    "a loan 1 intention",
				IntentionDetails: "a loan 1 intention details",
				Fees:             postgres.Jsonb{jMarshal},
				Interest:         1.5,
				TotalLoan:        float64(6500000),
				LayawayPlan:      500000,
			},
		}
		for _, loan := range loans {
			loan.Create()
		}
	}
}

// truncate defined tables. []string{"all"} to truncate all tables.
func Truncate(tableList []string) (err error) {
	if len(tableList) > 0 {
		if tableList[0] == "all" {
			tableList = []string{
				"banks",
			}
		}

		tables := strings.Join(tableList, ", ")
		sqlQuery := fmt.Sprintf("TRUNCATE TABLE %s RESTART IDENTITY CASCADE", tables)
		err = asira.App.DB.Exec(sqlQuery).Error
		return err
	}

	return fmt.Errorf("define tables that you want to truncate")
}
