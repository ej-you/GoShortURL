package main

import (
	"fmt"
	"os"

	"github.com/Danil-114195722/GoShortURL/db"
	"github.com/Danil-114195722/GoShortURL/settings"
)

func main() {
	// var err error

	fmt.Println("Hello!")

	args := os.Args

	// если при запуске указан аргумент "migrate"
	if len(args) > 1 {
		if args[1] == "migrate" {
			db.Migrate()
			return
		}
	}

	settings.InfoLog.Println("hahaha")
	settings.ErrorLog.Println("Oops")

	// dbConnect := db.GetConnection()
}
