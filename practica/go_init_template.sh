#!/bin/bash
echo "Creando e inicializando template"
go version
go mod init main
go get -u github.com/jackc/pgx/v5/pgxpool
go get -u github.com/joho/godotenv
mkdir domain handlers repository service config pkg validator templates
touch README.md .env main.go templates/login.html templates/index.html templates/logout.html
touch handlers/login.go validator/password.go pkg/hashpass.go config/connection.go
go get -u github.com/gorilla/securecookie
go get -u golang.org/x/crypto/bcrypt
go get -u github.com/golang-jwt/jwt/v5