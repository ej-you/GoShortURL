package services

import (
	"time"
	base62 "github.com/jcoene/go-base62"

	"github.com/Danil-114195722/GoShortURL/db"
)


func CreateLinkAlias(userLink string) (string, error) {
	dbConnection := db.GetConnection()

	// текущее время переводим в число с основанием 62
	nowTime := time.Now().UnixMilli()
	nowTimeBase62 := base62.Encode(nowTime)

	// создание записи в БД
	newLinkEntry := db.Link{
		ExternalLink: userLink,
		LinkAlias: nowTimeBase62,
	}
	createResult := dbConnection.Create(&newLinkEntry)

	return nowTimeBase62, createResult.Error
}

func GetLinkAliasIfEntryExist(externalLink string) (string, error) {
	dbConnection := db.GetConnection()

	var linkEntry db.Link
	// получение записи ссылки из БД
	getResult := dbConnection.Where("external_link = ?", externalLink).Find(&linkEntry)
	err := getResult.Error
	
	if err != nil {
		return "", err
	} else {
		// получение псевддонима ссылки
		linkAlias := linkEntry.LinkAlias
		return linkAlias, nil
	}
}

