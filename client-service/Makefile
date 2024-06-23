
migrate_create:
	migrate create -ext sql -dir migrations -seq client

migrate_up:
	migrate -database postgres://postgres:3115@localhost:5432/client?sslmode=disable -path ./migrations up

migrate_down:
	migrate -database postgres://postgres:3115@localhost:5432/client?sslmode=disable -path ./migrations migrate_down

migrate_force:
	migrate -database postgres://postgres:3115@localhost:5432/client?sslmode=disable -path ./migrations force


