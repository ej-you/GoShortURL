info_log = "$(GOPATH)/src/github.com/Danil-114195722/GoShortURL/logs/info-log.log"
error_log = "$(GOPATH)/src/github.com/Danil-114195722/GoShortURL/logs/error-log.log"


dev:
	go run ./main.go

migrate:
	go run ./main.go migrate

compile:
	go install .

prod:
	$(GOPATH)/bin/GoShortURL migrate >> $(info_log) 2>> $(error_log)
	$(GOPATH)/bin/GoShortURL >> $(info_log) 2>> $(error_log)
