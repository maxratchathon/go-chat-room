hello:
	echo "Hello World!"

dev: 
	go run cmd/chat-room/main.go 

# the user are "postgres"
postgres:
	docker run --name go-chat-rooms -e POSTGRES_PASSWORD=secret -p 5432:5432 -d postgres



