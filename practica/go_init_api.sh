#!/bin/bash
echo "Creando e inicializando API"
go version
go mod init main
go get -u github.com/jackc/pgx/v5/pgxpool
go get -u github.com/joho/godotenv
mkdir domain handlers repository service config pkg validator auth middleware
touch auth/jwt.go config/connection.go validator/password.go middleware/auth.go handlers/login.go
touch .env README.md main.go
go get -u github.com/gorilla/securecookie
go get -u golang.org/x/crypto/bcrypt
go get -u github.com/golang-jwt/jwt/v5