# Variabel
APP_NAME=grab-your-labubu
SERVER_MAIN=cmd/server/main.go

# 1. Run in local
run:
	go run $(SERVER_MAIN)

# 2. Tidy up library
tidy:
	go mod tidy

# 3. Run unit test
test:
	go test ./...

# 4. Build binary (make sure no error before deploy)
build:
	go build -o bin/$(APP_NAME) $(SERVER_MAIN)

# 5. Deployment to Fly.io
deploy:
	fly deploy

# 6. Status check in Fly.io
status:
	fly status