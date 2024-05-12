package main

import (
	"fmt"

	"github.com/Danil-114195722/GoShortURL/settings"
)

func main() {
	fmt.Println("Hello, world!")

	settings.InfoLog.Printf("hahaha")
	settings.ErrorLog.Printf("Oops")
}
