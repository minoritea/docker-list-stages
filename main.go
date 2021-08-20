package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
)

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}
}

func run() error {
	dockerfile := "Dockerfile"
	if len(os.Args) >= 2 {
		dockerfile = os.Args[1]
	}

	f, err := os.Open(dockerfile)
	if err != nil {
		return fmt.Errorf("cannot open Dockerfile(%s): %w)", dockerfile, err)
	}
	defer f.Close()

	stages, err := parse(f)
	if err != nil {
		return fmt.Errorf("Failed to parse Dockerfile(%s): %w)", dockerfile, err)
	}

	for _, stage := range stages {
		fmt.Println(stage)
	}
	return nil
}

var matcher = regexp.MustCompile(`^(FROM|from) ([\w:.-]+) (AS|as) (\w+)$`)

func parse(r io.Reader) ([]string, error) {
	scanner := bufio.NewScanner(r)
	var result []string
	for scanner.Scan() {
		line := scanner.Bytes()
		m := matcher.FindAllSubmatch(line, -1)
		if len(m) > 0 && len(m[0]) > 4 {
			result = append(result, string(m[0][4]))
		}
	}
	return result, scanner.Err()
}
