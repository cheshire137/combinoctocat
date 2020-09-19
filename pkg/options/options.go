package options

import (
	"errors"
	"flag"
)

type Options struct {
	InputPath string
}

func ParseOptions() (*Options, error) {
	var inputPath string
	flag.StringVar(&inputPath, "in", "",
		"Path to input HTML for myOctocat page source")

	flag.Parse()

	if len(inputPath) < 1 {
		flag.Usage()
		return nil, errors.New("No input path given")
	}

	return &Options{InputPath: inputPath}, nil
}
