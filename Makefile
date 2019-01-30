GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOGET=$(GOCMDX) get
BINARY_NAME=app
BINARY_UNIX=$(BINARY_NAME)_unix

all: build run
		@echo "Starting application server..."

get-deps:
				go get github.com/julienschmidt/httprouter
				go get github.com/lib/pq
build:
			@echo "Building application..."
			$(GOBUILD) -o $(BINARY_NAME) -v
run: build
		@echo "Starting application server..."
		./$(BINARY_NAME)
		@echo "Done!"

clean:
			@echo "Hello from target one..."
			$(GOCLEAN)
			rm -f $(BINARY_NAME)
			rm -f $(BINARY_UNIX)
			@echo "Done cleaning"
