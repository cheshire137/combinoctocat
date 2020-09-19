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

	if len(inputPath) < 1 {
		flag.Usage()
		return nil, errors.New("No input path given")
	}

	flag.Parse()

	return &Options{InputPath: inputPath}, nil
}
