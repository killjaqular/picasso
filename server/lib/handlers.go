package lib

import (
    // Standard Golang
    "crypto/sha256"
    "database/sql"
    "fmt"
    "log"
    "net/http"

    // Non-standard Golang
    "github.com/google/uuid"
)

type Handlers struct{}

func (h Handlers) ServeStaticFile(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, r.URL.Path[1:])
}

func (h Handlers) HomePage(w http.ResponseWriter, r *http.Request) {
    // Generate the HTML page
    html := `
    <!DOCTYPE html>
    <html>
    <head>
        <title>Picasso</title>
        <style>
            @font-face {
                font-family: 'Title Font';
                src: url('/picasso/fonts/MaybeOneDay.ttf') format('truetype');
                font-weight: normal;
                font-style: normal;
            }
            @font-face {
                font-family: 'Regular Font';
                src: url('/picasso/fonts/ArchivoBlack.ttf') format('truetype');
                font-weight: normal;
                font-style: normal;
            }
            body {
                background-image: url('/picasso/images/homepage.jpg');
                background-size: cover;
                background-repeat: no-repeat;
                background-attachment: fixed;
            }
            h1 {
                text-align: center;
                padding: 20px;
                font-family: 'Title Font', sans-serif; /* Use the custom font */
                font-size: 150px;
                color: rgb(125, 25, 25); /* Adjust the RGB values for your desired color */
            }
            form {
                text-align: center; /* Center form content horizontally */
                padding: 20px;
                background-color: rgba(125, 25, 25, 0.6);
                border-radius: 10px;
                margin: 0 auto;
                width: 15%;
                min-width: 200px; /* Set a minimum width for the form */
                font-family: 'Regular Font', sans-serif;
                opacity: 0;
                transition: opacity 0.5s;
            }
            form.show {
                opacity: 1; /* When the "show" class is added, the form fades in */
            }
            /* Override default input and button behavior to center them */
            label, input, input[type="submit"] {
                display: block;
                margin: 10px auto; /* Center the elements horizontally */
                font-family: 'Regular Font', sans-serif;
            }
        </style>

        <script>
            // JavaScript code for fading in the form
            window.addEventListener('load', function() {
                var form = document.querySelector('form');
                form.style.transition = 'opacity 0.5s';
                form.style.opacity = '1';
            });
        </script>
    </head>

    <body>
        <h1>Picasso</h1>
        <form action="/create-account" method="post" style="opacity: 0;">
            <label for="username">Username:</label>
            <input type="text" id="username" name="username" required>
            <label for="password">Password:</label>
            <input type="password" id="password" name="password" required>
            <input type="submit" value="Create Account">
        </form>
    </body>
    </html>
    `

    // Set the response content type to HTML
    w.Header().Set("Content-Type", "text/html")
    // Write the HTML to the response
    fmt.Fprint(w, html)
}

func (h Handlers) CreateAccountHandler(databaseConnection *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        if r.Method == http.MethodPost {
            // Get the form values
            username := r.FormValue("username")
            password := r.FormValue("password")

            ////////////////////////////////////////////////////////////////
            // Create instances to hold the information
            ////////////////////////////////////////////////////////////////
            // Populate information for a User
            var newUser User
            // Generate UUID for the new user
            newUser.UUID = uuid.New()
            newUser.Username = username
            // Generate SHA256bit of the password
            newUser.Password = fmt.Sprintf("%x", sha256.Sum256([]byte(password)))
            // Populate information for a UUID
            var newUUID UUIDRecord
            newUUID.UUID = newUser.UUID

            ////////////////////////////////////////////////////////////////
            // Check if username is already used in the database
            ////////////////////////////////////////////////////////////////
            // 1. Query for username
            var existingUsername string
            err := databaseConnection.QueryRow("SELECT username FROM accounts.user WHERE username = $1",
                                               username).Scan(&existingUsername)
            if err == nil {
                // 2.a if present, reject user creation with pop-up and redirection
                userExistsError := `
                <!DOCTYPE html>
                <html>
                <head>
                    <title>Username Taken</title>
                    <style>
                        body {
                            background-image: url('/picasso/images/homepage.jpg');
                            background-size: cover;
                            background-repeat: no-repeat;
                            background-attachment: fixed;
                            display: flex;
                            justify-content: center;
                            align-items: center;
                            height: 100vh;
                            margin: 0;
                        }
                        .popup {
                            background-color: rgba(0, 0, 0, 0.8);
                            color: white;
                            padding: 20px;
                            border-radius: 10px;
                            text-align: center;
                        }
                        .button {
                            padding: 10px 20px;
                            background-color: #f44336;
                            border: none;
                            color: white;
                            border-radius: 5px;
                            cursor: pointer;
                        }
                    </style>
                </head>
                <body>
                    <div class="popup">
                        <h2>Username is already taken.</h2>
                        <p>Please choose a different username.</p>
                        <button class="button" onclick="redirectToHomepage()">OK</button>
                    </div>
                    <script>
                        function redirectToHomepage() {
                            window.location.href = "/";
                        }
                    </script>
                </body>
                </html>
                `
                fmt.Fprintln(w, userExistsError)
                return
            } else if err != sql.ErrNoRows {
                // Handle other query errors
                log.Fatal(err)
            }
            // 2.b if absent, insert new user and new uuid into database
            _, err = databaseConnection.Exec("INSERT INTO uuids.uuid (uuid, parentTable) VALUES ($1, $2)",
                                             newUUID.UUID, "accounts.user")
            if err != nil {
                log.Fatal(err)
            }
            log.Printf("[i] Inserted new uuid:")
            log.Printf("           uuid: %s\n", newUUID.UUID)
            log.Printf("    parentTable: %s\n", "accounts.user")
            _, err = databaseConnection.Exec("INSERT INTO accounts.user (uuid, username, password) VALUES ($1, $2, $3)",
                                             newUser.UUID, newUser.Username, newUser.Password)
            if err != nil {
                log.Fatal(err)
            }
            log.Printf("[i] Inserted new user:")
            log.Printf("        uuid: %s\n", newUUID.UUID)
            log.Printf("    username: %s\n", newUser.Username)
            log.Printf("    password: %s\n", newUser.Password)

            userCreateSuccess := `
            <!DOCTYPE html>
            <html>
            <head>
                <title>New User Created!</title>
                <style>
                    body {
                        background-image: url('/picasso/images/homepage.jpg');
                        background-size: cover;
                        background-repeat: no-repeat;
                        background-attachment: fixed;
                        display: flex;
                        justify-content: center;
                        align-items: center;
                        height: 100vh;
                        margin: 0;
                    }
                    .popup {
                        background-color: rgba(0, 0, 0, 0.8);
                        color: white;
                        padding: 20px;
                        border-radius: 10px;
                        text-align: center;
                    }
                    .button {
                        padding: 10px 20px;
                        background-color: #f44336;
                        border: none;
                        color: white;
                        border-radius: 5px;
                        cursor: pointer;
                    }
                </style>
            </head>
            <body>
                <div class="popup">
                    <h2>New User Created!</h2>
                    <p>You may now login with your account.</p>
                    <button class="button" onclick="redirectToHomepage()">OK</button>
                </div>
                <script>
                    function redirectToHomepage() {
                        window.location.href = "/";
                    }
                </script>
            </body>
            </html>
            `
            fmt.Fprintln(w, userCreateSuccess)
        } else {
            http.Error(w, "[x] Method not allowed", http.StatusMethodNotAllowed)
        }
    }
}
