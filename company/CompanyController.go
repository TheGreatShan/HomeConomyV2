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

func GetCompany(c *gin.Context) {
	var company Company

	connection, err := services.GetDbConnection(c)

	if err != nil {
		panic(err)
	}
	id, err := hex.DecodeString(c.Param("id"))

	if err != nil {
		panic(err)
	}
	query, err := connection.Query("SELECT * FROM companies WHERE id = ?", id)

	defer query.Close()
	if query.Next() {
		if err := query.Scan(&company.Id, &company.Name); err != nil {
			panic(err)
		}
		company = Company{Id: hex.EncodeToString([]byte(company.Id)), Name: company.Name}
		c.IndentedJSON(http.StatusOK, company)
	} else {
		c.IndentedJSON(http.StatusNotFound, gin.H{"Message": "company not found"})
	}
}

func CreateCompany(c *gin.Context) {
	var company Company

	connection, err := services.GetDbConnection(c)

	if err != nil {
		panic(err)
	}

	if err := c.ShouldBindJSON(&company); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	id, err := hex.DecodeString(services.CreateUuid())
	company.Id = hex.EncodeToString(id)
	connection.Query("INSERT INTO companies (id, name) VALUES (?, ?)", id, company.Name)

	c.IndentedJSON(http.StatusCreated, company)
}
