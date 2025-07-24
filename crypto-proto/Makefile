# Makefile

PROTO_DIR := proto
PROTO_FILES := $(shell find $(PROTO_DIR) -name '*.proto')

.PHONY: proto clean help

proto:
	@for file in $(PROTO_FILES); do \
		echo "Generating $$file..."; \
		protoc --go_out=. --go_opt=paths=source_relative \
		       --go-grpc_out=. --go-grpc_opt=paths=source_relative \
		       $$file; \
	done

clean:
	@find . -name '*.pb.go' -delete

help:
	@echo "Available targets:"
	@echo "  make proto    # Generate Go code from .proto files"
	@echo "  make clean    # Delete all generated .pb.go files"
