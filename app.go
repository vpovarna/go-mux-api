package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

//App structure defines application components.
//A sql DB connection and router
type App struct {
	Router *mux.Router
	DB     *sql.DB
}

//Initialize function is responsable for establish a DB connection
func (a *App) Initialize(user string, password string, dbname string) {
	connectionString := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, password, dbname)

	var err error
	a.DB, err = sql.Open("postgres", connectionString)

	if err != nil {
		log.Fatal(err)
	}

	a.Router = mux.NewRouter()
}

//Run function will start the application
func (a *App) Run(addr string) {

}
