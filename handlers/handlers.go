package handlers

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"

	"github.com/Danil-114195722/GoShortURL/exceptions"
	"github.com/Danil-114195722/GoShortURL/services"
	"github.com/Danil-114195722/GoShortURL/settings"
)


func Index(response http.ResponseWriter, request *http.Request) {
	// рендеринг HTML файла
	err := RenderHTML(response, request, "index.html", nil)

	if err == nil {
		// даннные запроса
		remoteHost, fullRequestInfo := services.GetRequestInfo(request)
		// лог об успешном запросе
		settings.InfoLog.Printf("Request: %q %d %d (from %s)", fullRequestInfo, 200, request.ContentLength, remoteHost)
	}
}

func CreateShortLink(response http.ResponseWriter, request *http.Request) {
	var err error
	// псевдоним для внешней ссылки
	var linkAlias string

	// парсинг формы запроса
	if err = request.ParseForm(); err != nil {
		exceptions.FormParseError(response, request)
		return
	}

	// достаём внешнюю ссылку для сокращения из запроса
	externalLink := request.PostForm["link"][0]

	// проверка внешней ссылки юзера на валидность
	err = services.CheckLinkIsValid(externalLink)
	if err != nil {
		exceptions.InvalidLinkError(response, request)
		return
	}

	// если запись с такой внешней ссылкой ужже есть в БД, то достаём для неё псевдоним
	linkAlias, err = services.GetLinkAliasIfEntryExist(externalLink)
	if err != nil {
		exceptions.UnknownError(response, request, err)
		return
	}

	if len(linkAlias) == 0 {
		// создание записи в БД
		linkAlias, err = services.CreateLinkAlias(externalLink)
		if err != nil {
			exceptions.CreateLinkEntryError(response, request, err)
			return
		}
	}

	// переадресация для вывода результата
	redirectUrl := fmt.Sprintf("/result?linkAlias=%s", linkAlias)
	http.Redirect(response, request, redirectUrl, 301)
}

func Result(response http.ResponseWriter, request *http.Request) {
	var err error

	// парсинг формы запроса
	if err = request.ParseForm(); err != nil {
		exceptions.FormParseError(response, request)
		return
	}

	// достаём код для сокращённой ссылки из параметров запроса
	data := struct {
		ShortLink string
	}{
		ShortLink: fmt.Sprintf("%s/short/%s", settings.HostForShortLink, request.Form["linkAlias"][0]),
	}

	// рендеринг HTML файла
	err = RenderHTML(response, request, "result.html", data)

	if err == nil {
		// даннные запроса
		remoteHost, fullRequestInfo := services.GetRequestInfo(request)
		// лог об успешном запросе
		settings.InfoLog.Printf("Request: %q %d %d (from %s)", fullRequestInfo, 201, request.ContentLength, remoteHost)
	}
}

func Short(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)

	// получение псевдонима из запроса
	linkAlias := vars["linkAlias"]

	// достаём из БД внешнюю ссылку для редиректа
	externalLink, err := services.GetExternalLink(linkAlias)
	if err != nil {
		exceptions.GetExternalLinkError(response, request, err)
		return
	}

	http.Redirect(response, request, externalLink, 301)
}
