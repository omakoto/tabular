package main

import (
	"flag"
	"fmt"
	"github.com/omakoto/bashcomp"
	"github.com/omakoto/mlib"
	"github.com/omakoto/tabular/tabular"
)

var (
	useCsv             = flag.Bool("c", false, "input is CSV")
	useTsv             = flag.Bool("t", false, "input is TSV")
	useSpaceSeparated  = flag.Bool("s", true, "input is space-separated")
	useCustomSeparator = flag.String("p", "", "use regex as a separator")
)

func main() {
	flag.Parse()
	bashcomp.HandleBashCompletion()

	in := mlib.GetFilesReaderFromArgs()

	var reader <-chan []string
	switch {
	case *useCsv:
		reader = tabular.CsvReader(in)
	case *useTsv:
		reader = tabular.TsvReader(in)
	case *useCustomSeparator != "":
		reader = tabular.RegexpSeparatedReader(in, *useCustomSeparator)
	case *useSpaceSeparated: // Default, check it last.
		reader = tabular.RegexpSeparatedReader(in, `\s+`)
	}

	for l := range tabular.Tabular(reader) {
		fmt.Print(l)
		fmt.Print("\n")
	}
}
