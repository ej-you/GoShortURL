package services

import (
	"github.com/Danil-114195722/GoShortURL/db"
)


func GetExternalLink(linkAlias string) (string, error) {
	dbConnection := db.GetConnection()

	var linkEntry db.Link
	// получение записи ссылки из БД
	getResult := dbConnection.Where("link_alias = ?", linkAlias).Find(&linkEntry)
	err := getResult.Error
	
	if err != nil {
		return "", err
	} else {
		// получение сокращённой ссылки
		externalLink := linkEntry.ExternalLink
		return externalLink, nil
	}
}
