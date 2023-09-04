package lib

import (
    // Standard Golang
    "crypto/rand"
    "crypto/sha256"
    "database/sql"
    "encoding/base64"
    "fmt"
    "log"
    "net/http"
    "time"

    // Non-standard Golang
    "github.com/google/uuid"
)

type Handlers struct{}

/*
Handles the "/" endpoint.
Servers the Home Page of the Picasso WebGUI.
Shows title of page and a Create Account/Login form.
In:  w - Object to write response into.
     r - Object of the HTTP request.
Out: w - Write HTML directly into the response.
*/
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
    </head>

    <body>
        <h1>Picasso</h1>
        <form action="/create-account" method="post" class="show"> <!-- Added class "show" for initial visibility -->
            <label for="username">Username:</label>
            <input type="text" id="username" name="username" required>
            <label for="password">Password:</label>
            <input type="password" id="password" name="password" required>
            <input type="submit" value="Create Account">
            <input type="button" value="Login" onclick="redirectToLogin()">
        </form>

        <script>
            // JavaScript code for fading in the form
            window.addEventListener('load', function() {
                var form = document.querySelector('form');
                form.style.transition = 'opacity 0.5s';
                form.style.opacity = '1';
            });
            function redirectToLogin() {
                var username = document.getElementById("username").value;
                var password = document.getElementById("password").value;
                var loginUrl = "/login-account?username=" + encodeURIComponent(username) + "&password=" + encodeURIComponent(password);
                window.location.href = loginUrl;
            }
            
        </script>
    </body>
    </html>
    `

    // Set the response content type to HTML
    w.Header().Set("Content-Type", "text/html")
    // Write the HTML to the response
    fmt.Fprint(w, html)
}

/*
Handles the "/create-account" endpoint.
Handles logic to create a new user in the database.
In:  w - Object to write response into.
     r - Object of the HTTP request.
Out: w - Write HTML directly into the response.
*/
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
            // Insert newUUID
            log.Printf("[i] New uuid:\n\tuuid: %s\n\tparentTable: %s\n",
                       newUUID.UUID, "accounts.user")
            insertQuery := "INSERT INTO uuids.uuid (uuid, parentTable) VALUES ($1, $2)"
            _, err = databaseConnection.Exec(insertQuery, newUUID.UUID, "accounts.user")
            if err != nil {
                log.Fatal(err)
            }
            log.Printf("[+] Inserted!\n")
            // Insert newUser
            log.Printf("[i] New user:\n\tuuid: %s\n\tusername: %s\n\tpassword: %s\n",
                       newUUID.UUID,
                       newUser.Username,
                       newUser.Password)
            insertQuery = "INSERT INTO accounts.user (uuid, username, password) VALUES ($1, $2, $3)"
            _, err = databaseConnection.Exec(insertQuery,
                                             newUser.UUID,
                                             newUser.Username,
                                             newUser.Password)
            if err != nil {
                log.Fatal(err)
            }
            log.Printf("[+] Inserted!\n")

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

/*
Handles the "/login-account" endpoint.
Handles logic to login a user and create a session key.
In:  w - Object to write response into.
     r - Object of the HTTP request.
Out: w - Write HTML directly into the response.
*/
func (h Handlers) LoginAccountHandler(databaseConnection *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        username := r.URL.Query().Get("username")
        password := r.URL.Query().Get("password")

        // Populate information for a User
        var loginUser User
        loginUser.Username = username
        // Generate SHA256bit of the password
        loginUser.Password = fmt.Sprintf("%x", sha256.Sum256([]byte(password)))
        // Populate information for a UUID
        var loginUUID UUIDRecord

        // Check if username and password are in the database
        log.Printf("[i] Validating username:password:\n\tusername: %s\n\tpassword: %s\n",
                   loginUser.Username,
                   loginUser.Password)
        query := "SELECT uuid FROM accounts.user WHERE username = $1 AND password = $2"
        err := databaseConnection.QueryRow(query,
                                           loginUser.Username,
                                           loginUser.Password).Scan(&loginUUID.UUID)
        if err != nil {
            // 2.a if abscent, reject user login with pop-up and redirection
            invalidUsernamePassword := `
            <!DOCTYPE html>
            <html>
            <head>
                <title>Invalid Username or Password</title>
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
                    <h2>Invalid username or password.</h2>
                    <p>Please use valid login credentials or create a new user.</p>
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
            fmt.Fprintln(w, invalidUsernamePassword)
            return
        }
        loginUser.UUID = loginUUID.UUID
        log.Printf("[+] Valid!\n")
        // Create a new session key and store it in the database
        sessionKey := generateSessionKey()
        expiration := time.Now().Add(time.Hour) // Set session expiration to 1 hour
        var loginSession Session
        loginSession.SessionKey = sessionKey
        loginSession.UserUUID   = loginUser.UUID
        loginSession.Expiration = expiration
        log.Printf("[i] New key:\n\tsessionKey: (256 Raw Bytes, not worth printing.)\n\tuserUUID: %s\n\texpiration: %s\n",
                   loginSession.UserUUID,
                   loginSession.Expiration)
        insertQuery := "INSERT INTO sessionKeys.session (sessionKey, userUUID, expiration) VALUES ($1, $2, $3)"
        _, err = databaseConnection.Exec(insertQuery,
                                         loginSession.SessionKey,
                                         loginSession.UserUUID,
                                         loginSession.Expiration)
        if err != nil {
            http.Error(w, "[x] Login failed. Unable to create session.", http.StatusInternalServerError)
            log.Printf("[x] Login failed. Unable to create session.\n")
            return
        }
        log.Printf("[+] Inserted!\n")

        // Set the session key as a cookie and redirect user
        encodedSessionKey := base64.StdEncoding.EncodeToString(loginSession.SessionKey)
        http.SetCookie(w, &http.Cookie{
            Name:    "sessionKey",
            Value:   encodedSessionKey,
            Expires: loginSession.Expiration,
            // Other cookie settings as needed
        })

        loginAccountSuccess := `
        <!DOCTYPE html>
        <html>
        <head>
            <title>Login Account!</title>
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
                <h2>Welcome!</h2>
                <p>You may now manage your account.</p>
                <button class="button" onclick="redirectToDashboard()">OK</button>
            </div>
            <script>
                function redirectToDashboard() {
                    window.location.href = "/dashboard";
                }
            </script>
        </body>
        </html>
        `
        fmt.Fprintln(w, loginAccountSuccess)
    }
}

// TODO: Implement this
/*
Handles the "/dashboard" endpoint.
Handles logic for a user to manage their account.
Displays options for a user to 'delete', 'download', 'upload', and 'view' their images.
In:  w - Object to write response into.
     r - Object of the HTTP request.
Out: w - Write HTML directly into the response.
*/
func (h Handlers) Dashboard(w http.ResponseWriter, r *http.Request) {
    ////////////////////////////////////////////////////////////////
    // Validate Session Key
    ////////////////////////////////////////////////////////////////
    // 1. Get browser cookie
    // 2. Validate sessionKey

    ////////////////////////////////////////////////////////////////
    // Handle user button selection
    ////////////////////////////////////////////////////////////////
    // 1. Delete Image
    // 2. Download Image
    // 3. Upload Image
    // 4. View Image

    http.Error(w, "[x] Not yet implemented.", http.StatusMethodNotAllowed)
}

// TODO: Implement this
/*
Handles the "/logout-account" endpoint.
Handles logic to logout a user from the server.
Invalidates all session keys and require a user to visit "/login-account" to re-authorize.
In:  w - Object to write response into.
     r - Object of the HTTP request.
Out: w - Write HTML directly into the response.
*/
func (h Handlers) LogoutAccountHandler(databaseConnection *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        http.Error(w, "[x] Not yet implemented.", http.StatusMethodNotAllowed)
    }
}

// TODO: Implement this
/*
Handles the "/upload-image" endpoint.
Handles logic to upload an image to the server.
In:  w - Object to write response into.
     r - Object of the HTTP request.
Out: w - Write HTML directly into the response.
*/
func (h Handlers) UploadImageHandler(databaseConnection *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        http.Error(w, "[x] Not yet implemented.", http.StatusMethodNotAllowed)
    }
}

// TODO: Implement this
/*
Handles the "/download-image" endpoint.
Handles logic to download an image from the server.
In:  w - Object to write response into.
     r - Object of the HTTP request.
Out: w - Write HTML directly into the response.
*/
func (h Handlers) DownloadImageHandler(databaseConnection *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        http.Error(w, "[x] Not yet implemented.", http.StatusMethodNotAllowed)
    }
}

// TODO: Implement this
/*
Handles the "/delete-image" endpoint.
Handles logic to delete an image from the server.
In:  w - Object to write response into.
     r - Object of the HTTP request.
Out: w - Write HTML directly into the response.
*/
func (h Handlers) DeleteImageHandler(databaseConnection *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        http.Error(w, "[x] Not yet implemented.", http.StatusMethodNotAllowed)
    }
}

// TODO: Implement this
/*
Handles the "/dashboard" endpoint.
The dashboard is where a user may perform almost all functions on their account:
1. View their images.
2. Upload an image.
3. Download an image.
4. Delete an image.
5. Logout.
In:  w - Object to write response into.
     r - Object of the HTTP request.
Out: w - Write HTML directly into the response.
*/
func (h Handlers) DashboardHandler(databaseConnection *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        http.Error(w, "[x] Not yet implemented.", http.StatusMethodNotAllowed)
    }
}

/*
Generates 256 random bytes used to identify a valid session of a logged-in user.
In:  NONE
Out: sessionKeyBytes - 256 random bytes.
*/
func generateSessionKey() []byte {
    // Generate a byte slice of 256 random bytes
    sessionKeyBytes := make([]byte, 256)
    _, err := rand.Read(sessionKeyBytes)
    if err != nil {
        // Handle error
        panic(err)
    }
    return sessionKeyBytes
}
