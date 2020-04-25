package main

import (
	"database/sql"
	_ "database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

// Config params for db
const (
	HOST     = "localhost"
	PORT     = 54320
	USERNAME = "postgres"
	PASSWORD = "password"
	DBNAME   = "hn_db"
)

func main() {
	_connect()
}

func _connect() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		HOST, PORT, USERNAME, PASSWORD, DBNAME)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	_createTable(db)

}

func _createTable(db *sql.DB) {
	tableName := "posts"
	tables := _getTables(db)
	postsTableExists := false
	for _, value := range tables {
		if value == tableName{
			postsTableExists = true
		}
	}
	if postsTableExists {
		log.Println("Table " + tableName + " already created")

	} else {
		stmt, err := db.Prepare(`CREATE TABLE IF NOT EXISTS ` + tableName + ` (
			id SERIAL PRIMARY KEY,
			title TEXT,
			url TEXT,
			type TEXT,
			score INT,
			timestamp TIMESTAMP);`)

		if err != nil {
			log.Fatal(err)
		}

		defer stmt.Close()

		_, err = stmt.Exec()
		if err != nil {
			log.Fatal(err)
		}

		log.Println("Table " + tableName + " succesfully created")

	}

}

func _getTables(db *sql.DB) []string {
	var tableName string
	var tables []string

	query := `SELECT table_name
	FROM information_schema.tables
	WHERE table_schema = 'public'
	ORDER BY table_name;`

	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&tableName)
		tables = append(tables,tableName)
		if err != nil {
			log.Fatal(err)
		}
	}

	return tables

}
