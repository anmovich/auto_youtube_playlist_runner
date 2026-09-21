.SILENT:
build: cmd/main.go
	go build cmd/main.go

run: build
	./main
cont_stop:
	docker-compose down
cont_run: cont_stop
	docker-compose up
migrate_up: 
	migrate -path ./schema -database "postgres://anme:1234@localhost:5432/postgres?sslmode=disable" up
migrate_down:
	migrate -path ./schema -database "postgres://anme:1234@localhost:5432/postgres?sslmode=disable" down

