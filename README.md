## Sharebite - Food Sharing

This is an application working as part of a food sharing product to reduce food waste.

Runs on:
- Go (Gin)
- Postgres
- Docker

To setup the app:
- Clone the repo
- Run `go mod download`
- If using Docker, run: `docker build -t sharebite-api .`
- Finally run the app. For normal setup run: `go build && ./main.go`
- For Docker: `docker-compose up --build`

ToDos:
- [ ] Complete the Docker Compose file
- [ ] Deploy Kubes
- [ ] Finish writing tests
- [ ] Do OpenAPI docs