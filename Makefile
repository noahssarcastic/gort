default: build

run_path = ./cmd/trace

.PHONY: build
build:
	@go build -o ./build/trace $(run_path)

.PHONY: run
run:
	@go run $(run_path) -o trace.ppm

.PHONY: cpuprofile
cpuprofile:
	@go run $(run_path) -o trace.ppm -prof trace.prof

.PHONY: test
test:
	@go test ./...

.PHONY: setup
setup:
	@go work use -r .

clean:
	rm -f *.ppm *.prof
	rm -rf output

.PHONY: smoketest
smoketest:
	@mkdir -p output
	go run ./cmd/ppm -o output/test.ppm
	go run ./cmd/clock -o output/clock.ppm
	go run ./cmd/projectile -o output/projectile.ppm
	go run ./cmd/scene -o output/scene.ppm

# .PHONY: docs
# docs:
# 	xdg-open https://pkg.go.dev/github.com/noahssarcastic/gort