package main

import (
	"flag"
	"fmt"
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

	in := mlib.GetFilesReaderFromArgs()

	var reader <-chan []string
	switch {
	case *useCsv:
		reader = tabular.CsvReader(in)
	case *useTsv:
		reader = tabular.TsvReader(in)
	case *useSpaceSeparated:
		reader = tabular.RegexpSeparatedReader(in, `\s+`)
	case *useCustomSeparator != "":
		reader = tabular.RegexpSeparatedReader(in, *useCustomSeparator)
	}

	for l := range tabular.Tabular(reader) {
		fmt.Print(l)
		fmt.Print("\n")
	}
}
