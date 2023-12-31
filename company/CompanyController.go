package company

import (
	"HomeConomyv2GO/services"
	"encoding/hex"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Company struct {
	Id   string
	Name string
}

func GetCompanies(c *gin.Context) {
	var companies []Company

	connection, err := services.GetDbConnection(c)

	if err != nil {
		panic(err)
	}

	query, err := connection.Query("SELECT * FROM companies")
	if err != nil {
		panic("laksdj")
	}
	defer query.Close()
	for query.Next() {
		var company Company
		if err := query.Scan(&company.Id, &company.Name); err != nil {
			panic(err)
		}
		company = Company{Id: hex.EncodeToString([]byte(company.Id)), Name: company.Name}
		companies = append(companies, company)
	}

	c.IndentedJSON(http.StatusOK, companies)
}
