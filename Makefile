
.PHONY: install
install:
	go install ./cmd/protoc-gen-goose

.PHONY: test
test:
	go test -v ./...

.PHONY: example
example:
	protoc \
	--proto_path=. \
	--proto_path=./third_party \
	--proto_path=./../ \
	--go_out=. \
	--go_opt=paths=source_relative \
	--goose_out=. \
	--goose_opt=paths=source_relative \
	example/*/*.proto

all: install example
