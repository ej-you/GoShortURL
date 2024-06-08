info_log = "/logs/info-log.log"
error_log = "/logs/error-log.log"


dev:
	go run ./main.go

migrate:
	go run ./main.go migrate

compile:
	go build -o ./GoShortURL ./main.go

prod:
	@echo "Running migrations..."
	/root/GoShortURL migrate
	@echo "Running main app..."
	/root/GoShortURL >> $(info_log) 2>> $(error_log)
