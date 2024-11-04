package db

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/ej-you/GoShortURL/settings"
)


// создание таблиц в БД по структурам в GO
func Migrate() {
	fmt.Println("Start migration...")
	
	fmt.Printf("Connectiong to %q database...\n", settings.PathDB)
	db := GetConnection()

	err := db.AutoMigrate(&Link{})
	settings.DieIf(err)

	fmt.Println("DB -- Migrated successfully!")
}

// получение соединения с БД
func GetConnection() (db *gorm.DB) {
	connection, err := gorm.Open(sqlite.Open(settings.PathDB), &gorm.Config{})
	settings.DieIf(err)

	return connection
}
