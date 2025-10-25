package belajar_golang_database

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestExecSql(t *testing.T) {
	db := GetConnection() // return a database pool
	defer db.Close() // close the db pool

	ctx := context.Background()

	script := "INSERT INTO customer(id, name) VALUES('Noby', 'Mjolnir')"
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

func TestQuerySqlComplex(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "SELECT id, name, email, balance, rating, birth_date, married, created_at FROM customer"
	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id, name string
		var email sql.NullString // if the column allows NULL value, use custom data type from sql package
		var balance int32
		var rating float64
		var birthDate sql.NullTime
		var createdAt time.Time
		var married bool

		err = rows.Scan(&id, &name, &email, &balance, &rating, &birthDate, &married, &createdAt)
		if err != nil {
			panic(err)
		}
		fmt.Println("================")
		fmt.Println("Id:", id)
		fmt.Println("Name:", name)
		if email.Valid {
			fmt.Println("Email:", email.String)
		}
		fmt.Println("Balance:", balance)
		fmt.Println("Rating:", rating)
		if birthDate.Valid {
			fmt.Println("Birth Date:", birthDate.Time)
		}
		fmt.Println("Married:", married)
		fmt.Println("Created At:", createdAt)
	}
}

func TestSqlInjection(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "admin'; #"
	password := "salah"

	// use concatenation instead of substitution(preapre statement)
	script := "SELECT username FROM user WHERE username = '" + username +
		"' AND password = '" + password + "' LIMIT 1" //  this script is vulnerable to sql injection
	fmt.Println(script)
	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("Sukses Login", username)
	} else {
		fmt.Println("Gagal Login")
	}
}

func TestSqlInjectionSafe(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "admin'; #"
	password := "salah"

	// user parameter instead of concatenation(substitution)
	script := "SELECT username FROM user WHERE username = ? AND password = ? LIMIT 1"
	fmt.Println(script)
	rows, err := db.QueryContext(ctx, script, username, password) // variable position must be sequential
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("Sukses Login", username)
	} else {
		fmt.Println("Gagal Login")
	}
}

func TestExecSqlParameter(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "noby'; DROP TABLE user; #"
	password := "sasuke12"

	script := "INSERT INTO user(username, password) VALUES(?, ?)"
	_, err := db.ExecContext(ctx, script, username, password)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new user")
}

func TestAutoIncrement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	email := "noby@gmail.com"
	comment := "Comment test"

	script := "INSERT INTO comments(email, comment) VALUES(?, ?)"
	result, err := db.ExecContext(ctx, script, email, comment)
	if err != nil {
		panic(err)
	}

	// this method that will return lastest number of column that set to auto_increment
	insertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new comment with id", insertId)
}

func TestPrepareStatement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	// ExecContext and QueryContext implicitly prepare the statement each time,
	// but they don't guarantee reusing the same DB connection.
	// Use PrepareContext if you need a reusable prepared statement.

	ctx := context.Background()
	script := "INSERT INTO comments(email, comment) VALUES(?, ?)"
	statement, err := db.PrepareContext(ctx, script) // bind one connection
	if err != nil {
		panic(err)
	}
	defer statement.Close() // return the connection to the db pool

	for i := 0; i < 10; i++ {
		email := "noby" + strconv.Itoa(i) + "@gmail.com"
		comment := "Komentar ke " + strconv.Itoa(i)

		// use the one connection that has been prepared
		result, err := statement.ExecContext(ctx, email, comment) 
		if err != nil {
			panic(err)
		}

		id, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}

		fmt.Println("Comment Id ", id)
	}

	/*
		The prepared method creates a dedicated connection for the input query,
		allowing it to be reused. Suitable for queries where parameter values ​​frequently change.
	*/

	/*
		Prepared statement cocok untuk repeat queries on the same connection, bukan sekadar untuk “banyak traffic”.
		Scope efektifnya adalah repeated execution in one session, bukan “many users hitting the same endpoint”.
	*/
}
