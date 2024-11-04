package handlers

import (
	"html/template"
	"net/http"
	"path"

	"github.com/ej-you/GoShortURL/settings"
	"github.com/ej-you/GoShortURL/exceptions"
)


func RenderHTML(response http.ResponseWriter, request *http.Request, htmlTemplate string, data interface{}) error {
	// путь к HMTL файлу
	htmlPath := path.Join(settings.TemplatesPath, htmlTemplate)
	
	// чтение HTML файла
	html, err := template.ParseFiles(htmlPath)
	if err != nil {
		exceptions.TemplateParseError(response, request, htmlPath)
		return err
	}

	// рендеринг HTML файла
	err = html.Execute(response, data)
	if err != nil {
		exceptions.TemplateExecuteError(response, request, htmlPath)
		return err
	}

	return nil
}
