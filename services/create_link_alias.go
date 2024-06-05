package services

import (
	"time"
	"fmt"
	base62 "github.com/jcoene/go-base62"

	"github.com/Danil-114195722/GoShortURL/db"
	"github.com/Danil-114195722/GoShortURL/settings"
)


func CreateLinkAlias(userLink string) (string, error) {
	dbConnection := db.GetConnection()

	// текущее время переводим в число с основанием 62
	nowTime := time.Now().UnixMilli()
	nowTimeBase62 := base62.Encode(nowTime)

	// полная сокращённая ссылка
	shortLink := fmt.Sprintf("%s/short/%s", settings.HostForShortLink, nowTimeBase62)

	// создание записи в БД
	newLinkEntry := db.Link{
		ExternalLink: userLink,
		RedirectLink: shortLink,
	}
	createResult := dbConnection.Create(&newLinkEntry)

	return nowTimeBase62, createResult.Error
}
