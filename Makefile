start:
	docker-compose up

stop:
	sudo docker-compose stop

migrateUp:
	sudo docker run -v /home/izifizik/Dev/GolangProjects/WB-test-L0/migrations:/migrations --network host migrate/migrate -path=/migrations/ -database "postgres://WB:WB@localhost:5432/WB?sslmode=disable" up


migrateDown:
	sudo docker run -v /home/izifizik/Dev/GolangProjects/WB-test-L0/migrations:/migrations --network host migrate/migrate -path=/migrations/ -database "postgres://WB:WB@localhost:5432/WB?sslmode=disable" down -all

test:
	go test ./internal/repository/database

.PHONY: start stop migrateUp migrateDown test
