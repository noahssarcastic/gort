package main

import (
	"flag"
	"log"
	"os"
)

type config struct {
	filePath string
	file     *os.File

	cpuProfilePath string
	cpuProfile     *os.File
	memProfilePath string
	memProfile     *os.File
}

var cfg config

func init() {
	flag.StringVar(&cfg.filePath, "o", "", "write image to `file`")
	flag.StringVar(
		&cfg.cpuProfilePath, "cpuprofile", "", "write cpu profile to `file`")
	flag.StringVar(
		&cfg.memProfilePath,
		"memprofile",
		"",
		"write memory profile to `file`")
}

func initConfig() {
	flag.Parse()

	if cfg.filePath == "" {
		log.Println("Set output file with -o flag.")
		os.Exit(0)
	}
	cfg.file = openFile(cfg.filePath)

	if cfg.cpuProfilePath != "" {
		cfg.cpuProfile = openFile(cfg.cpuProfilePath)
	}
	if cfg.memProfilePath != "" {
		cfg.memProfile = openFile(cfg.memProfilePath)
	}
}

func openFile(path string) *os.File {
	f, err := os.OpenFile(path,
		os.O_RDWR|os.O_CREATE|os.O_TRUNC,
		0755)
	if err != nil {
		log.Panic(err)
	}
	return f
}

func cleanUp() {
	if cfg.file != nil {
		closeFile(cfg.file)
	}
	if cfg.cpuProfile != nil {
		closeFile(cfg.cpuProfile)
	}
	if cfg.memProfile != nil {
		closeFile(cfg.memProfile)
	}
}

func closeFile(f *os.File) {
	if f == nil {
		return
	}
	err := f.Close()
	if err != nil {
		log.Panic(err)
	}
}
