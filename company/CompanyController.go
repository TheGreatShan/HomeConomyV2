package company

import (
	"HomeConomyv2GO/handlers"
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
	handlers.Panic(err)

	companies, err = GetAllCompanies(connection)
	handlers.Panic(err)

	c.IndentedJSON(http.StatusOK, companies)
}

func GetCompany(c *gin.Context) {
	var company Company

	connection, err := services.GetDbConnection(c)

	handlers.Panic(err)

	id, err := hex.DecodeString(c.Param("id"))
	handlers.Panic(err)

	exists := handlers.CheckIfExistsById(c, "companies")
	if exists == false {
		c.IndentedJSON(http.StatusNotFound, gin.H{"Message": "company not found"})
		return
	}
	company, err = GetCompanyById(connection, id)
	handlers.Panic(err)

	c.IndentedJSON(http.StatusOK, company)
}

func CreateCompany(c *gin.Context) {
	var company Company

	connection, err := services.GetDbConnection(c)

	handlers.Panic(err)

	if err := c.ShouldBindJSON(&company); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	result := handlers.CheckIfExistsByName(c, "companies", company.Name)

	if result == true {
		c.IndentedJSON(http.StatusConflict, gin.H{"Message": "company already exists"})
		return
	}

	id, err := hex.DecodeString(services.CreateUuid())
	company.Id = hex.EncodeToString(id)
	connection.Query("INSERT INTO companies (id, name) VALUES (?, ?)", id, company.Name)

	c.IndentedJSON(http.StatusCreated, company)
}

func UpdateCompany(c *gin.Context) {
	var company Company

	connection, err := services.GetDbConnection(c)

	handlers.Panic(err)

	result := handlers.CheckIfExistsById(c, "companies")

	if result == false {
		c.IndentedJSON(http.StatusNotFound, gin.H{"Message": "company not found"})
		return
	}

	if err := c.ShouldBindJSON(&company); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	id, err := hex.DecodeString(c.Param("id"))

	handlers.Panic(err)

	connection.Query("UPDATE companies SET name = ? WHERE id = ?", company.Name, id)

	c.IndentedJSON(http.StatusCreated, company)
}

func DeleteCompany(c *gin.Context) {
	connection, err := services.GetDbConnection(c)

	handlers.Panic(err)

	id, err := hex.DecodeString(c.Param("id"))

	handlers.Panic(err)

	result := handlers.CheckIfExistsById(c, "companies")

	if result == false {
		c.IndentedJSON(http.StatusNotFound, gin.H{"Message": "company not found"})
		return
	}

	connection.Query("DELETE FROM companies WHERE id = ?", id)

	c.IndentedJSON(http.StatusNoContent, gin.H{})
}
