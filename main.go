package main

import (
	"belajar-golang-database-migration/app"
	"belajar-golang-database-migration/controller"
	"belajar-golang-database-migration/helper"
	"belajar-golang-database-migration/middleware"
	"belajar-golang-database-migration/repository"
	"belajar-golang-database-migration/service"
	"net/http"

	"github.com/go-playground/validator"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := app.NewDB()
	validate := validator.New()

	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := app.NewRouter(categoryController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
