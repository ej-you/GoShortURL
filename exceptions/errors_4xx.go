package exceptions

import (
	"net/http"

	"github.com/Danil-114195722/GoShortURL/services"
	"github.com/Danil-114195722/GoShortURL/settings"
)


func FormParseError(response http.ResponseWriter, request *http.Request) {
	// даннные запроса
	remoteHost, fullRequestInfo := services.GetRequestInfo(request)
	// статус ответа
	respStatus := 400

	// лог о неудавшемся запросе
	settings.ErrorLog.Printf("Request: %q %d %d (from %s)", fullRequestInfo, respStatus, request.ContentLength, remoteHost)
	// лог с описанием ошибки
	settings.ErrorLog.Print("Failed to parse Form data")
	http.Error(response, "400: Failed to parse Form data", respStatus)
}


func InvalidLinkError(response http.ResponseWriter, request *http.Request) {
	// даннные запроса
	remoteHost, fullRequestInfo := services.GetRequestInfo(request)
	// статус ответа
	respStatus := 400

	// лог о неудавшемся запросе
	settings.ErrorLog.Printf("Request: %q %d %d (from %s)", fullRequestInfo, respStatus, request.ContentLength, remoteHost)
	// лог с описанием ошибки
	settings.ErrorLog.Print("Invalid link was given")
	http.Error(response, "400: Invalid link was given", respStatus)
}
