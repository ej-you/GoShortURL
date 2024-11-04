package main

import (
	"fmt"
	"net/http"
	"os"
	"github.com/gorilla/mux"

	"github.com/ej-you/GoShortURL/db"
	"github.com/ej-you/GoShortURL/handlers"
	"github.com/ej-you/GoShortURL/settings"
)

func main() {
	var err error

	args := os.Args

	// если при запуске указан аргумент "migrate"
	if len(args) > 1 {
		if args[1] == "migrate" {
			db.Migrate()
			return
		}
	}

	// маршрутизатор для запросов
	router := mux.NewRouter()

	// файловый сервер для статики
	fileServer := http.FileServer(http.Dir("./frontend/static"))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fileServer)).Methods("GET")
	router.Handle("/favicon.ico", fileServer).Methods("GET")

	// основные обработчики
	router.HandleFunc("/", handlers.Index).Methods("GET")
	router.HandleFunc("/create-short-link", handlers.CreateShortLink).Methods("POST")
	router.HandleFunc("/result", handlers.Result).Methods("GET")
	router.HandleFunc("/short/{linkAlias}", handlers.Short).Methods("GET")

	// на каком хосту запускается
	addr := fmt.Sprintf("%s:%s", settings.HostIp, settings.HostPort)
	
	// лог о запуске сервера с прослушкой настроенного адреса
	settings.InfoLog.Printf("Start server at http://%s/\n", addr)

	err = http.ListenAndServe(addr, router)
	settings.DieIf(err)
}
