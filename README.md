# Frank's MicroService Project

Demo project for deploying microservices

## Web UI

Login & Registration Page
Dashboard

### Running

From the `/web/ui` folder run `npm install && npm run build` to download dependencies and build the production 
application. This should create `index.html`, `styles.css`, and `[hash].js` files in a new folder `public`. This
is where the Go application serves files from.

Then simply run the Go Webapp: `go run main.go` in the `/web` folder 

## Auth Service

Simple authentication stub available over RPC.

## GitHub Issues Fetcher

Grabs open issues for the given repository ID. Available over RPC

