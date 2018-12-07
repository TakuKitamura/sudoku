GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=sudoku

build:
	$(GOBUILD) -o $(BINARY_NAME)
test:
	cd ./src/api && $(GOTEST)

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
run:
	echo "[GO TEST!]"
	cd ./src/api && $(GOTEST)
	echo "[GO BUILD!]"
	$(GOBUILD) -o $(BINARY_NAME) -v
	echo "[GO RUN!]"
	./$(BINARY_NAME)