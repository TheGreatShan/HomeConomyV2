package company

import (
	"HomeConomyv2GO/handlers"
	"HomeConomyv2GO/services"
	"database/sql"
	"encoding/hex"
)

func GetAllCompanies(connection *sql.DB) ([]Company, error) {
	var companies []Company

	query, err := connection.Query("SELECT * FROM companies")

	if err != nil {
		return nil, err
	}

	defer query.Close()
	for query.Next() {
		var company Company
		if err := query.Scan(&company.Id, &company.Name); err != nil {
			return nil, err
		}
		company = Company{Id: hex.EncodeToString([]byte(company.Id)), Name: company.Name}
		companies = append(companies, company)
	}
	return companies, nil
}

func GetCompanyById(connection *sql.DB, id []byte) (Company, error) {
	var company Company

	query, err := connection.Query("SELECT * FROM companies WHERE id = ?", id)

	if err != nil {
		return company, err
	}

	defer query.Close()

	if query.Next() {
		if err := query.Scan(&company.Id, &company.Name); err != nil {
			return company, err
		}
		company = Company{Id: hex.EncodeToString([]byte(company.Id)), Name: company.Name}
	}

	return company, nil
}
func CreateNewCompany(connection *sql.DB, company Company) (Company, error) {
	id, err := hex.DecodeString(services.CreateUuid())
	handlers.Panic(err)
	company.Id = hex.EncodeToString(id)

	result, err := connection.Query("INSERT INTO companies (id, name) VALUES (?, ?)", id, company.Name)
	if err != nil {
		return company, err
	}

	defer result.Close()

	return company, nil
}

func UpdateCompanyById(connection *sql.DB, id []byte, company Company) (Company, error) {
	result, err := connection.Query("UPDATE companies SET name = ? WHERE id = ?", company.Name, id)
	if err != nil {
		return company, err
	}

	defer result.Close()
	return company, nil
}