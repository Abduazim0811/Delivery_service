
migrate_create:
	migrate create -ext sql -dir migrations -seq driver

migrate_up:
	migrate -database postgres://postgres:3115@localhost:5432/driver?sslmode=disable -path ./migrations up

migrate_down:
	migrate -database postgres://postgres:3115@localhost:5432/driver?sslmode=disable -path ./migrations down

migrate_force:
	migrate -database postgres://postgres:3115@localhost:5432/driver?sslmode=disable -path ./migrations force


