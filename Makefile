run_module = github.com/noahssarcastic/tddraytracer/cmd/run

test_modules = ./geometry/ray \
	./geometry/sphere \
	./image/canvas \
	./image/color \
	./image/ppm \
	./math/matrix \
	./math/transform \
	./math/tuple \
	./math/utils

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
		go test $$mod; \
	done
