package main

import (
	"fmt"
	"os"

	"flag"

	"io"

	"github.com/alistanis/cf_parameter_generator/lib"
)

var (
	inFile                       string
	outFile                      string
	min                          bool
	numIndentSpaces              int
	overwrite                    bool
	removeOldParamsNotInTemplate bool
	oyaml                        bool
	inyaml                       bool
	verbose                      bool
	printVersion                 bool
)

const (
	version = "cf_parameter_generator: v1.9"
)

func init() {
	flag.StringVar(&inFile, "f", "", "The file to read from to generate parameters.")
	flag.StringVar(&outFile, "o", "", "Optional: Specify a file name to write out parameters.")
	flag.BoolVar(&min, "min", false, "If given, will generate minified output.")
	flag.IntVar(&numIndentSpaces, "spaces", 2, "The number of spaces used to indent the file if not generating minified output.")
	flag.BoolVar(&overwrite, "overwrite", false, "By default, will update an existing parameters file with newly found parameters, but will not overwrite.")
	flag.BoolVar(&removeOldParamsNotInTemplate, "r", false, "Removes old entries from parameters found in old parameters files.")
	flag.BoolVar(&oyaml, "outyaml", false, "Will output in yaml instead of json.")
	flag.BoolVar(&inyaml, "inyaml", false, "Will expect input as yaml instead of json.")
	flag.BoolVar(&verbose, "v", false, "Places verbose output in the ParameterValue field to be replaced.")
	flag.BoolVar(&printVersion, "version", false, "Print the version and exits.")
	flag.Parse()
	if printVersion {
		fmt.Println(version)
		os.Exit(0)
	}
}

func config() *cfpgen.Config {
	return &cfpgen.Config{InFile: inFile, OutFile: outFile, Minimize: min, Indent: numIndentSpaces, Overwrite: overwrite, RemoveOldParamsNotInTemplate: removeOldParamsNotInTemplate, OutYaml: oyaml, InYaml: inyaml, Verbose: verbose}
}

func main() {
	var reader io.Reader
	stat, err := os.Stdin.Stat()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		reader = os.Stdin
	}
	err = cfpgen.Generate(config(), reader)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		flag.Usage()
		os.Exit(-1)
	}
}
