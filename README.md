# Picasso

### Problem:

<p>
Ability to host images on a networked service.
</p>

### Solution:

<p>
HTTP Server to host image storage. A user will create an account, save, delete, and download their images.
</p>

## Development Notes:

<p>
Golang documentation can be found at: https://go.dev/doc/<br>
PostgreSQL documentation can be found at: https://www.postgresql.org/docs/
</p>

## Summary of solution

<p>
We will be using Docker containers to host 2 services:
<ol>
<li>Database, `postgresql`:</li>
The `Database` will receive data from:
<ol type='i'>
<li>The `Subscriber`.</li>
<li>The `HTTP Server`.</li>
</ol>

<li>Server, Custom HTTP Server written in Golang:</li>
The Server will:
<ol type='i'>
<li>Require a user to register an account.</li>
<li>Require a user to sign-in to their account.</li>
<li>Allow users to upload images.</li>
<li>Allow users to see a list of their saved images on their account from their machine.</li>
<li>Allow users to delete an image on their account.</li>
<li>Allow users to download an image on their account to their machine.</li>
</ol>
</ol>
</p>

## `database`

### Building

Use the `make` recipe:
> `user@host:$>make database`

### Developing

All related material for the development of the `Database` will be found in the `database/` directory.
All configurations required during the runtime of `Database` will be found in the `configs/` directory.

### Testing

<p>
Minimum testing is required for:
<ol>
<li>Logging into accounts that do not exist.</li>
<li>Inserting images of 0 bytes.</li>
<li>Inserting images of 1 Gigabyte. (1,073,741,824 bytes)</li>
<li>Deleting images that do not exist.</li>
<li>Downloading images that do not exist.</li>
</ol>
</p>

## `server`

### Building

Use the `make` recipe:
> `user@host:$>make server`

### Developing

All related material for the development of the `Server` will be found in the `server/` directory.
All configurations required during the runtime of `Server` will be found in the `configs/` directory.

### Testing

<p>
Minimum testing is required for:
<ol>
<li>Logging into accounts that do not exist.</li>
<li>Inserting images of 0 bytes.</li>
<li>Inserting images of 1 Gigabyte. (1,073,741,824 bytes)</li>
<li>Deleting images that do not exist.</li>
<li>Downloading images that do not exist.</li>
TODO: Add to the list as test cases are created.
</ol>
</p>

## Integration

### Deployment

Use the `make` recipe:
> `user@host:$>make run`

### Testing

<p>
Minimum testing is required for:
<ol>
<li>Logging into accounts that do not exist.</li>
<li>Inserting images of 0 bytes.</li>
<li>Inserting images of 1 Gigabyte. (1,073,741,824 bytes)</li>
<li>Deleting images that do not exist.</li>
<li>Downloading images that do not exist.</li>
TODO: Add to the list as test cases are created.
</ol>
</p>

## `Makefile`

### Recipes:

`database` - Pulls the Docker image of a `postgresql` free to use database and includes the project's tables, triggers, configurations, etc.<br>
`server` - Builds our HTTP Server solution.
# picasso
