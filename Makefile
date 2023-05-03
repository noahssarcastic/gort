default: build

run_path = ./cmd/trace

.PHONY: build
build:
	@go build -o ./build/trace $(run_path)

.PHONY: run
run:
	@go run $(run_path)

.PHONY: test
test:
	@go test ./...

.PHONY: setup
setup:
	@go work use -r .

clean:
	-rm *.ppm

smoketest:
	go run ./cmd/ppm
	go run ./cmd/clock
	go run ./cmd/projectile
	go run ./cmd/trace
