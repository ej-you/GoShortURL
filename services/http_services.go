package services

import (
	"fmt"
	"net/http"
	"strings"
)


// получение удалённого IP и строки запроса
func GetRequestInfo(request *http.Request) (string, string) {
	// удалённый IP
	remoteHost := strings.Split(request.RemoteAddr, ":")[0]
	// полная строка запроса с хостом и URI
	fullRequestInfo := fmt.Sprintf("%s %s %s", request.Method, request.URL, request.Proto)

	return remoteHost, fullRequestInfo
}
