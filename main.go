package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Danil-114195722/GoShortURL/db"
	"github.com/Danil-114195722/GoShortURL/handlers"
	"github.com/Danil-114195722/GoShortURL/settings"
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

	// файловый сервер для статики
	fileServer := http.FileServer(http.Dir("./frontend/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))

	// основные обработчики
	http.HandleFunc("/", handlers.Index)

	// на каком хосту запускается
	addr := fmt.Sprintf("%s:%s", settings.HostIp, settings.HostPort)
	
	// лог о запуске сервера с прослушкой настроенного адреса
	settings.InfoLog.Printf("Start server on %s...\n", addr)

	err = http.ListenAndServe(addr, nil)
	settings.DieIf(err)
}
