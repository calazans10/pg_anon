package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/jessevdk/go-flags"
)

type Options struct {
	File   string
	Output string
	Fields string
}

func main() {
	opts, err := parseArgs()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	p := Processor{
		File:   opts.File,
		Output: opts.Output,
		Fields: strings.Split(opts.Fields, ","),
	}
	p.Run()
}

func parseArgs() (*Options, error) {
	var opts struct {
		File   string `short:"d" long:"dump" description:"Path to the dump file"`
		Output string `short:"o" long:"output" description:"Path to the output file"`
		Fields string `short:"f" long:"fields" description:"List of fields to anonymize"`
		Help   bool   `long:"help" description:"Show help"`
	}

	parser := flags.NewParser(&opts, flags.None)
	parser.Usage = "[arguments]"

	_, err := parser.Parse()
	if err != nil {
		parser.WriteHelp(os.Stderr)
		return nil, err
	}

	if opts.Help {
		parser.WriteHelp(os.Stdout)
		os.Exit(0)
	}

	if opts.File == "" {
		opts.File = "./in.sql"
	}

	if opts.Output == "" {
		opts.Output = "./out.sql"
	}

	if opts.Fields == "" {
		parser.WriteHelp(os.Stderr)
		return nil, fmt.Errorf("required flag `-f, --fields` not specified")
	}

	return &Options{
		File:   opts.File,
		Output: opts.Output,
		Fields: opts.Fields,
	}, nil
}
