.PHONY: clean go_mod_tidy_vendor init test_unit test_integration generate generate_mock generate_wire migrate_version migrate_up

clean:
	rm -rf vendor
	find . -name "*.mock.gen.go" -type f -delete
	find . -name "wire_gen.go" -type f -delete

go_mod_tidy_vendor:
	-go mod init
	go mod tidy
	go mod vendor

init: generate_mock \
	go_mod_tidy_vendor \
	generate_wire

test_unit:
	go test -json -short -coverprofile coverage.out -v ./... | grep FAIL

test_integration:
	go test -json -coverprofile coverage_integration.out -v ./... -tags=integration 2>&1

generate: generate_mock \
	generate_wire

INTERFACE_GO_FILES := $(shell find internal -name "interface.go")
INTERFACE_GEN_GO_FILES := $(INTERFACE_GO_FILES:%.go=%.mock.gen.go)
generate_mock: $(INTERFACE_GEN_GO_FILES)
$(INTERFACE_GEN_GO_FILES): %.mock.gen.go: %.go
	@echo "Generating mock $@ for $<"
	mockgen -source=$< -destination=$@ -package=$(shell basename $(dir $<))

generate_wire:
	wire ./...

migrate_version:
	migrate -path migration -database ${DATABASE_URL} version

migrate_up:
	migrate -path migration -database ${DATABASE_URL} up
