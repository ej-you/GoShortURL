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

// путь до SQLite3 БД
var PathDB string = "./db.sqlite3"

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

