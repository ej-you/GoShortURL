package services

import (
	"net/http"
)


func CheckLinkIsValid(link string) error {
	_, err := http.Get(link)

	if err != nil {
		return err
	} else {
		return nil
	}
}
