run_module=github.com/noahssarcastic/tddraytracer/cmd/run

default: build

.PHONY: run
run:
	@go run $(run_module)

.PHONY: build
build:
	@go build -o ./build/tddraytracer $(run_module)