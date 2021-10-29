package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	colPattern   = ":(?:.*)$"
	fakerPattern = ":(?:.*)$"
)

type processor struct {
	File   string
	Output string
	Fields []string
}

func NewProcessor(file, output, fields string) processor {
	if file == "" {
		file = "./in.sql"
	}

	if output == "" {
		output = "./out.sql"
	}

	return processor{File: file, Output: output, Fields: strings.Split(fields, ",")}
}

func (p processor) Run() {
	mappings := make([]Mapping, len(p.Fields))

	for index, field := range p.Fields {
		col := getColumn(field, colPattern)

		faker := getFaker(field, fakerPattern)
		if faker == "" {
			faker = col
		}

		mappings[index] = Mapping{Col: col, Faker: faker}
	}

	fmt.Fprintln(os.Stderr, "Running...")

	iterator := Iterator{Mappings: mappings}

	file, err := os.Open(p.File)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	output, err := os.Create(p.Output)
	if err != nil {
		log.Fatal(err)
	}
	defer output.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		_, err := output.WriteString(fmt.Sprintf("%s\n", iterator.ProcessLine(scanner.Text())))
		if err != nil {
			log.Fatal(err)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	output.Sync()

	fmt.Fprintln(os.Stderr, "Done")
}

func getColumn(line, pattern string) string {
	col := Replace(line, pattern)
	return strings.ToLower(col)
}

func getFaker(line, pattern string) string {
	if !strings.Contains(line, ":") {
		return ""
	}

	return Replace(line, pattern)
}
