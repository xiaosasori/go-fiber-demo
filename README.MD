# Golang - Fiber - Postgres - Docker

## Run the project

- MacOS: `go run *.go`
- Windows: `go run ./.`

## Make docker update the containers after changing Dockerfile
- Run `docker-compose up --build`

## Run backend in docker
- `docker-compose exec backend sh`

## gorm
- `Preload("prop")` to preload associations
- `gorm:"-"`: dont add this prop to the db

## models
- `models/user.go` use `revenue *float64 json:"omitempty"` pointer here to make response return `revenue: 0` for ambassador