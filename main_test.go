package main

import (
	"log"
	"os"
	"testing"
)

var a App

const createTableQuery = `CREATE TABLE IF NOT EXISTS products 
(
	id SERIAL,
	name TEXT NOT NULL,
	price NUMERIC(10,2) NOT NULL DEFAULT 0.00,
	CONSTRAINT products_pkey PRIMARY KEY (id)
)`

const purgeTableQuery = `DELETE FROM products`

const resetPrimaryKey = `ALTER SEQUENCE products_id_seq RESTART WITH 1`

func TestMain(m *testing.M) {
	a.Initialize(
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"))

	ensureTableExist()
	code := m.Run()
	clearTable()
	os.Exit(code)
}

func ensureTableExist() {
	_, err := a.DB.Exec(createTableQuery)
	if err != nil {
		log.Fatal(err)
	}
}

func clearTable() {
	a.DB.Exec(purgeTableQuery)
	a.DB.Exec(resetPrimaryKey)
}
