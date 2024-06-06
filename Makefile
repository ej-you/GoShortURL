info_log = "./logs/info-log.log"
error_log = "./logs/error-log.log"


dev:
	go run ./main.go

migrate:
	go run ./main.go migrate

compile:
	go install .

prod:
	/go/bin/GoShortURL migrate >> $(info_log) 2>> $(error_log)
	/go/bin/GoShortURL >> $(info_log) 2>> $(error_log)
