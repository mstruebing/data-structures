SOURCES = $(shell find $(SRC_DIR) -type f -name '*.go')

test: 
	go test -cover

bin/linkedlist: $(SOURCES)
	go build -o ./bin/linkedlist cmd/linkedlist/main.go

build: bin/linkedlist

run: bin/linkedlist
	./bin/linkedlist
