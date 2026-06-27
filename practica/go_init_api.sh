#!/bin/bash
echo "Creando e inicializando API"
go version
go mod init main
go get -u github.com/jackc/pgx/v5/pgxpool
go get -u github.com/joho/godotenv
mkdir domain handlers repository service config pkg validator
go get -u github.com/gorilla/securecookie
go get -u golang.org/x/crypto/bcrypt
go get -u github.com/golang-jwt/jwt/v5