debug:
	go run ./main.go

# заменить на запуск скомпилированного файла
prod:
	go run ./main.go >> ./logs/info-log.log 2>> ./logs/error-log.log