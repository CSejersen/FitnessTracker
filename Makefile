build: 
	@go build -o bin/fitnessTracker

run: build 
	@./bin/fitnessTracker

test:
	@go test -v ./...
