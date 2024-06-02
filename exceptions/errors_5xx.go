package exceptions

import (
	"net/http"

	"github.com/Danil-114195722/GoShortURL/services"
	"github.com/Danil-114195722/GoShortURL/settings"
)


func TemplateParseError(response http.ResponseWriter, request *http.Request, htmlPath string) {
	// даннные запроса
	remoteHost, fullRequestInfo := services.GetRequestInfo(request)
	// статус ответа
	respStatus := 500

	// лог о неудавшемся запросе
	settings.ErrorLog.Printf("Request: %q %d %d (from %s)", fullRequestInfo, respStatus, request.ContentLength, remoteHost)
	// лог с описанием ошибки
	settings.ErrorLog.Printf("Failed to parse template file: %s", htmlPath)
	http.Error(response, "500: Server internal error", respStatus)
}


func TemplateExecuteError(response http.ResponseWriter, request *http.Request, htmlPath string) {
	// даннные запроса
	remoteHost, fullRequestInfo := services.GetRequestInfo(request)
	// статус ответа
	respStatus := 500

	// лог о неудавшемся запросе
	settings.ErrorLog.Printf("Request: %q %d %d (from %s)", fullRequestInfo, respStatus, request.ContentLength, remoteHost)
	// лог с описанием ошибки
	settings.ErrorLog.Printf("Failed to execute template file: %s", htmlPath)
	http.Error(response, "500: Server internal error", respStatus)
}
