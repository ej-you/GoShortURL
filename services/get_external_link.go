package services

import (
	"fmt"

	"github.com/Danil-114195722/GoShortURL/db"
	"github.com/Danil-114195722/GoShortURL/settings"
)


func GetExternalLink(linkAlias string) (string, error) {
	dbConnection := db.GetConnection()

	// полная сокращённая ссылка
	shortLink := fmt.Sprintf("%s/short/%s", settings.HostForShortLink, linkAlias)
	
	var linkEntry db.Link
	// получение записи ссылки из БД
	getResult := dbConnection.Where("RedirectLink = ?", shortLink).Find(&linkEntry)

	// получение сокращённой ссылки
	linkToRedirect := linkEntry.ExternalLink

	return linkToRedirect, getResult.Error
}
