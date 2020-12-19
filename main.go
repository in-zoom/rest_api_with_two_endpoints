package main

import (
	"github.com/julienschmidt/httprouter"
	"github.com/subosito/gotenv"
	"log"
	"net/http"
	"os"
	"rest_api_with_two_endpoints/handlers"
	"rest_api_with_two_endpoints/middleware"
)

func main() {
	initialize()
	router := httprouter.New()
	router.POST("/login", handlers.LoginHandler())
	router.GET("/list", handlers.ListRecords())
	router.PUT("/update", middleware.AuthCheckMiddleware(middleware.UpdateEntry()))
	port := os.Getenv("PORT")
	log.Fatal(http.ListenAndServe(":"+port, router))
}

/*
Файл .env нужен для добавления пользователя, соли и порта!
Формат файла выглядит следующим образом:
NAME_USER=
USER_PASS=
SALT=
PORT=
*/
func initialize() {
	if err := gotenv.Load(); err != nil {
		log.Print("Файл .env не найден")
	}
}
