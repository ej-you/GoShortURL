package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"path"
	"strings"

	"github.com/Danil-114195722/GoShortURL/settings"
)


func Index(response http.ResponseWriter, request *http.Request) {
	// путь к HMTL файлу
	indexHtmlPath := path.Join(settings.TemplatesPath, "index.html")
	// чтение HTML файла
	html, err := template.ParseFiles(indexHtmlPath)
	if err != nil {
		settings.ErrorLog.Printf("Failed to parse template file: %s", indexHtmlPath)
		http.Error(response, err.Error(), 500)
		return
	}
	// рендеринг HTML файла
	err = html.Execute(response, nil)
	if err != nil {
		settings.ErrorLog.Printf("Failed to execute template file: %s", indexHtmlPath)
		http.Error(response, err.Error(), 500)
	}

	// удалённый IP
	remoteHost := strings.Split(request.RemoteAddr, ":")[0]
	// полная строка запроса с хостом и URI
	fullURL := fmt.Sprintf("http://%s:%s%s", settings.HostIp, settings.HostPort, request.URL)
	// лог о запросе
	settings.InfoLog.Printf("Request: %s (%s) from %s", fullURL, request.Method, remoteHost)
}
