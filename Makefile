dev:
	go run ./main.go

migrate:
	go run ./main.go migrate

# заменить на запуск скомпилированного файла, добавить выполнение в фоне
prod:
	go run ./main.go >> ./logs/info-log.log 2>> ./logs/error-log.log
