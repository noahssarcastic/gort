//go:build mage
// +build mage

package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/magefile/mage/sh"
)

const runPath = "./cmd/trace"

// build ray-tracer executable under ./build/
func Build() error {
	if err := sh.Run("go", "mod", "download"); err != nil {
		return err
	}
	return sh.Run("go", "build", "-o", "./build/trace", runPath)
}

func goRun(cmd, output string, extra ...string) error {
	args := []string{"run", cmd, "-o", output}
	return sh.Run("go", append(args, extra...)...)
}

// run the ray-tracer and outputs image to out.ppm
func Run() error {
	return goRun(runPath, "out.ppm")
}

// Profiling

// run the ray-tracer with cpu profiling
func Cpuprofile() error {
	return goRun(runPath, "out.ppm", "-cpuprofile", "cpu.prof")
}

// run the ray-tracer with memory profiling
func Memprofile() error {
	return goRun(runPath, "out.ppm", "-memprofile", "mem.prof")
}

// Testing

var goTest = sh.OutCmd("go", "test")

// run unit tests
func Test() error {
	out, err := goTest("./pkg/...")
	fmt.Println(out)
	return err
}

// run unit tests with test coverage
func Cov() error {
	_, err := goTest("-coverprofile", "cover.out", "./pkg/...")
	if err != nil {
		return err
	}
	return sh.Run("go", "tool", "cover", "-html", "./cover.out")
}

// run smoke-tests and output images to ./output/
func Smoke() error {
	err := os.Mkdir("./output", os.ModeDir)
	if err != nil && !os.IsExist(err) {
		return err
	}
	for _, t := range []struct{ cmd, output string }{
		{"./cmd/ppm", "output/test.ppm"},
		{"./cmd/clock", "output/clock.ppm"},
		{"./cmd/projectile", "output/projectile.ppm"},
	} {
		err = goRun(t.cmd, t.output)
		if err != nil {
			return err
		}
	}
	return nil
}

// Misc.

// remove build artifacts and output images
func Clean() error {
	if err := os.RemoveAll("build"); err != nil {
		return err
	}
	if err := os.RemoveAll("output"); err != nil {
		return err
	}
	if err := removeGlob("*.prof"); err != nil {
		return err
	}
	if err := removeGlob("*.ppm"); err != nil {
		return err
	}
	return removeGlob("*.out")
}

func removeGlob(glob string) error {
	fs, err := filepath.Glob(glob)
	if err != nil {
		return err
	}
	for _, f := range fs {
		if err := os.Remove(f); err != nil {
			return err
		}
	}
	return nil
}

// open documentation
func Docs() error {
	url := "https://pkg.go.dev/github.com/noahssarcastic/gort"
	switch runtime.GOOS {
	case "linux":
		return sh.Run("xdg-open", url)
	case "windows":
		return sh.Run("rundll32", "url.dll,FileProtocolHandler", url)
	case "darwin":
		return sh.Run("open", url)
	default:
		return errors.New("unsupported platform")
	}
}
