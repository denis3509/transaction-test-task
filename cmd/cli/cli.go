package main

import (
	"flag"
	"fmt"
	"log"
)

func help() {
	fmt.Println("recreate-test-db - пересоздать тестовую базу данных")
	flag.PrintDefaults()
}

var command string

func main() {
	flag.Parse()
	args := flag.Args()

	fmt.Println(args)

	if len(args) == 0 {
		help()
		log.Fatal("no args specified")
	}
	command = args[0]

	switch command {
	default:
		help()
	}
}
