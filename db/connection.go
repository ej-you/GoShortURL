package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/Danil-114195722/GoShortURL/settings"
)


// создание таблиц в БД по структурам в GO
func Migrate() {
	db := GetConnection()

	err := db.AutoMigrate(&Link{})
	settings.DieIf(err)

	settings.InfoLog.Println("DB -- Migrated successfully!")
}

// получение соединения с БД
func GetConnection() (db *gorm.DB) {
	connection, err := gorm.Open(sqlite.Open(settings.PathDB), &gorm.Config{})
	settings.DieIf(err)

	return connection
}
