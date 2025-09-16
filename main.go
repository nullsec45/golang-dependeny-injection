package main

import (
	// "github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	// "nullsec45/golang-dependency-injection/app"
	// "nullsec45/golang-dependency-injection/controller"
	"nullsec45/golang-dependency-injection/helper"
	"nullsec45/golang-dependency-injection/middleware"
	// "nullsec45/golang-dependency-injection/repository"
	// "nullsec45/golang-dependency-injection/service"
)

func NewServer(authMiddleware *middleware.AuthMiddleware) *http.Server {
	return &http.Server{
		Addr:    "localhost:3000",
		Handler: authMiddleware,
	}
}

func main() {
	server := InitializedServer()
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
