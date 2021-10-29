package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Processor struct {
	File   string
	Output string
	Fields []string
}

func (p Processor) Run() {
	mappings := make([]Mapping, len(p.Fields))

	for index, field := range p.Fields {
		col := getColumn(field)

		faker := getFaker(field)
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

func getColumn(field string) string {
	col := Replace(field, ":(?:.*)$", "")
	return strings.ToLower(col)
}

func getFaker(field string) string {
	if !strings.Contains(field, ":") {
		return ""
	}

	return Replace(field, ":(?:.*)$", "")
}
