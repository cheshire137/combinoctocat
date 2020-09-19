package models

import "flag"

type Options struct {
	InputPath string
}

func ParseOptions() *Options {
	var inputPath string
	flag.StringVar(&inputPath, "in", "",
		"Path to input HTML for myOctocat page source")

	flag.Parse()

	return &Options{InputPath: inputPath}
}

func DisplayOptions() {
	flag.Usage()
}
