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
	-rm -f ./**/*.ppm

.PHONY: smoketest
smoketest:
	@mkdir -p output
	go run ./cmd/ppm -o output/test.ppm
	go run ./cmd/clock -o output/clock.ppm
	go run ./cmd/projectile -o output/projectile.ppm
	go run ./cmd/trace -o output/trace.ppm
