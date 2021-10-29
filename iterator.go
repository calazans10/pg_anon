package main

import (
	"regexp"
	"strings"
)

type Mapping struct {
	Col   string
	Faker string
}

type Iterator struct {
	Mappings     []Mapping
	table        string
	columns      []string
	transformers []Anonymizer
}

func (i *Iterator) ProcessLine(line string) string {
	if strings.HasPrefix(line, "COPY") {
		i.processTable(line)
		return line
	} else if i.table != "" && strings.TrimSpace(line) != "" {
		return i.processRow(line)
	} else {
		i.table = ""
		return line
	}
}

func (i *Iterator) processTable(line string) {
	i.table = Gsub(line, "^COPY (.*?) .*$")

	i.columns = getColumns(Gsub(line, `^COPY (?:.*?) \((.*)\).*$`))

	i.transformers = getTransformers(i.columns, i.Mappings)
}

func (i Iterator) processRow(line string) string {
	values := strings.Split(line, "\t")
	result := make([]string, len(values))

	for index, value := range values {
		result[index] = i.transformers[index].Fake(value)
	}

	return strings.Join(result[:], "\t")
}

func getColumns(columns string) []string {
	return Map(strings.Split(columns, ","), func(v string) string {
		re := regexp.MustCompile(`"`)

		s := strings.TrimSpace(v)
		s = re.ReplaceAllString(s, "")
		s = strings.ToLower(s)

		return s
	})
}

func getTransformers(columns []string, mappings []Mapping) []Anonymizer {
	result := make([]Anonymizer, len(columns))

	for index, column := range columns {
		mapping := getMapping(mappings, column)
		result[index] = Anonymizer{Type: mapping.Faker}
	}

	return result
}

func getMapping(mappings []Mapping, column string) Mapping {
	for _, mapping := range mappings {
		if mapping.Col == column {
			return mapping
		}
	}
	return Mapping{Col: column, Faker: ""}
}
