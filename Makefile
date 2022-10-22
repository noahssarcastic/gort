run_module = github.com/noahssarcastic/tddraytracer/cmd/run

test_modules = ./tuple ./color ./canvas

default: build

.PHONY: run
run:
	@go run $(run_module)

.PHONY: build
build:
	@go build -o ./build/tddraytracer $(run_module)

.PHONY: test
test:
	@for mod in $(test_modules); do \
		go test -v $$mod; \
	done
