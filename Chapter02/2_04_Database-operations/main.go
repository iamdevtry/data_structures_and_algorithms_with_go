package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

// Customer Class
type Customer struct {
	CustomerId   int
	CustomerName string
	SSN          string
}

// GetConnection method which returns sql.DB
func GetConnection() (database *sql.DB) {
	databaseDriver := "mysql"
	databaseUser := "root"
	databasePass := "pass"
	databaseName := "customer_example"

	database, err := sql.Open(databaseDriver, databaseUser+":"+databasePass+"@/"+databaseName)
	if err != nil {
		panic(err.Error())
	}

	return database
}

// GetCustomers method returns Customer Array
func GetCustomers() []Customer {
	var database *sql.DB
	database = GetConnection()
	//var err error
	//var rows *sql.Rows
	rows, err := database.Query("SELECT * FROM customer ORDER BY CustomerId DESC")
	if err != nil {
		panic(err.Error())
	}
	customer := Customer{}

	var customers []Customer
	for rows.Next() {
		var customerId int
		var customerName string
		var ssn string

		err = rows.Scan(&customerId, &customerName, &ssn)
		if err != nil {
			panic(err.Error())
		}
		customer.CustomerId = customerId
		customer.CustomerName = customerName
		customer.SSN = ssn
		customers = append(customers, customer)

	}

	defer database.Close()
	return customers
}

// main method
func main() {
	var customers []Customer
	customers = GetCustomers()
	fmt.Println("Customers", customers)
}
