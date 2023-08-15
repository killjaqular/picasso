package main

import (
	// Standard Golang
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	// Non-standard Golang
	"main/lib"

	_ "github.com/lib/pq"
	"gopkg.in/yaml.v2"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Picasso")
}

func main() {
	////////////////////////////////////////////////////////////////
	// Initiazlie server
	////////////////////////////////////////////////////////////////
	// 1. Open config.yml
	inputFile, err := ioutil.ReadFile("/picasso/config.yml")
	if err != nil {
		log.Fatal(err)
	}
	// 2. Get configuration information
	var config lib.Config
	err = yaml.Unmarshal(inputFile, &config)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Server Port: %d\n", config.Server.Port)
	fmt.Printf("Database Host: %s\n", config.Database.Host)
	fmt.Printf("Database Port: %d\n", config.Database.Port)
	fmt.Printf("Database Password: %s\n", config.Database.Database)
	fmt.Printf("Database Username: %s\n", config.Database.Username)
	fmt.Printf("Database Password: %s\n", config.Database.Password)
	port := fmt.Sprintf("%d", config.Server.Port)
	// 3. Try to connect to the database
	connStr := "user=YOUR_databaseConnection_USER databaseConnectionname=YOUR_databaseConnection_NAME sslmode=disable"
	databaseConnection, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer databaseConnection.Close()

	// Ping the database to verify the connection
	err = databaseConnection.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to PostgreSQL!")

	// 4. Try to host service on host:port
	http.HandleFunc("/", homePage)
	log.Printf("Picasso server on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
