package main

import (
	"bufio"
	"fmt"
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

	fmt.Fprintln(os.Stdout, "Running...")

	iterator := Iterator{Mappings: mappings}

	file, err := os.Open(p.File)
	check(err)
	defer file.Close()

	output, err := os.Create(p.Output)
	check(err)
	defer output.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		_, err := output.WriteString(fmt.Sprintf("%s\n", iterator.ProcessLine(scanner.Text())))
		check(err)
	}

	err = scanner.Err()
	check(err)

	output.Sync()

	fmt.Fprintln(os.Stdout, "Done")
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

func check(e error) {
	if e != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", e)
		os.Exit(1)
	}
}
