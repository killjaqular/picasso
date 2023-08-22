package lib

import (
    "fmt"
    "net/http"
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

func (h Handlers) CreateAccountHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodPost {
        // Get the form values
        username := r.FormValue("username")
        password := r.FormValue("password")

        ////////////////////////////////////////////////////////////////
        // TODO: Check if username is already taken
        ////////////////////////////////////////////////////////////////

        ////////////////////////////////////////////////////////////////
        // TODO: Create hash for password
        ////////////////////////////////////////////////////////////////

        ////////////////////////////////////////////////////////////////
        // TODO: Create UUID for username
        ////////////////////////////////////////////////////////////////

        ////////////////////////////////////////////////////////////////
        // TODO: Insert UUID, username, password into 
        ////////////////////////////////////////////////////////////////

        // Respond with a message
        message := fmt.Sprintf("[i] Account created for %s:%s", username, password)
        fmt.Fprintln(w, message)
    } else {
        http.Error(w, "[x] Method not allowed", http.StatusMethodNotAllowed)
    }
}
