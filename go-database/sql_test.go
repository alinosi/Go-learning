package belajar_golang_database

import (
	"context"
	"fmt"
	"testing"
)

func TestExecSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "INSERT INTO customer(id, name) VALUES('Noby', 'Begone')"
	_, err := db.ExecContext(ctx, script) // intializes a connection to the database if there's no connection yet.
	if err != nil {
		panic(err)
	} // execution method doesn't requie variable to store data because there's no data return from DB.

	fmt.Println("Success insert new customer")
}

func TestQuerySql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "SELECT id, name FROM customer"
	rows, err := db.QueryContext(ctx, script) // read method need variable to store data return from DB.
	if err != nil {
		panic(err)
	}
	defer rows.Close() // always remember to close rows(data store variable) after using it.

	for rows.Next() { // iterate the rows.
		var id, name string

		// this method will manipulate data from another scope, that's why we use pointers
		err = rows.Scan(&id, &name)

		if err != nil {
			panic(err)
		}

		fmt.Println("Id:", id)
		fmt.Println("Name:", name)
	}
}

//  vulnerable to sql injection
