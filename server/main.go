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

func main() {
    ////////////////////////////////////////////////////////////////
    // Initialize server
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
    log.Printf("[i] Server Port: %d\n", config.Server.Port)
    log.Printf("[i] Database Host: %s\n", config.Database.Host)
    log.Printf("[i] Database Port: %d\n", config.Database.Port)
    log.Printf("[i] Database Database: %s\n", config.Database.Database)
    log.Printf("[i] Database Username: %s\n", config.Database.Username)
    log.Printf("[i] Database Password: %s\n", config.Database.Password)
    port := fmt.Sprintf("%d", config.Server.Port)
    // 3. Try to connect to the database
    connStr := fmt.Sprintf(
        "host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
        config.Database.Host,
        config.Database.Port,
        config.Database.Username,
        config.Database.Password,
        config.Database.Database,
    )
    databaseConnection, err := sql.Open("postgres", connStr)
    if err != nil {
        log.Fatal(err)
    }
    defer databaseConnection.Close()
    // 4. Query the database to verify the connection
    rows, err := databaseConnection.Query("SELECT 1 FROM uuids.uuid;")
    if err != nil {
        log.Fatal(err)
    }else{
        log.Printf("[i] Connected to %s:%d", config.Database.Host, config.Database.Port)
    }
    defer rows.Close()

    ////////////////////////////////////////////////////////////////
    // Serve static files
    ////////////////////////////////////////////////////////////////
    imageFs := http.FileServer(http.Dir("/picasso/images/"))
    http.Handle("/picasso/images/", http.StripPrefix("/picasso/images/", imageFs))
    fontsFs := http.FileServer(http.Dir("/picasso/fonts/"))
    http.Handle("/picasso/fonts/", http.StripPrefix("/picasso/fonts/", fontsFs))

    ////////////////////////////////////////////////////////////////
    // Setup all handlers/endpoints
    ////////////////////////////////////////////////////////////////
    // Create an instance of the Handlers struct and register handlers
    handlers := lib.Handlers{}
    http.HandleFunc("/", handlers.HomePage)
    http.HandleFunc("/create-account", handlers.CreateAccountHandler(databaseConnection))

    ////////////////////////////////////////////////////////////////
    // Start service on host:port
    ////////////////////////////////////////////////////////////////
    log.Printf("[i] Picasso server on port %s\n", port)
    log.Fatal(http.ListenAndServe(":"+port, nil))
}
