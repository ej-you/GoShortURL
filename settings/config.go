package settings

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)


// загрузка переменных окружения
var _ error = godotenv.Load("./.env")


// распаковка переменных окружения по переменным
var HostIp string = os.Getenv("HOST_IP")
var HostPort string = os.Getenv("HOST_PORT")
// URL для вывода сокращённой ссылки
var HostForShortLink string = os.Getenv("HOST_FOR_SHORT_LINK")

// путь до SQLite3 БД
var PathDB string = "./db.sqlite3"

var TemplatesPath = "./frontend/templates"
var StaticPath = "./frontend/static"

// логеры
var InfoLog *log.Logger = log.New(os.Stdout, "[INFO]\t", log.Ldate|log.Ltime)
var ErrorLog *log.Logger = log.New(os.Stderr, "[ERROR]\t", log.Ldate|log.Ltime|log.Lshortfile)

// функция для обработки критических ошибок
func DieIf(err error) {
	if err != nil {
		fatalLog := log.New(os.Stderr, "[FATAL]\t", log.Ldate|log.Ltime|log.Lshortfile)
		fatalLog.Panic(err)
	}
}

