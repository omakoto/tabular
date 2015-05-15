package tabular

import (
	"bytes"
	"github.com/omakoto/mlib"
	"strings"
)

func Tabular(r <-chan []string) <-chan string {
	out := make(chan string)

	go func() {
		defer close(out)
		doTabular(r, out)
	}()

	return out
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func doTabular(r <-chan []string, out chan<- string) {
	all := make([][]string, 0, 1024)
	widths := make([]int, 0)

	mlib.Debug("Reading input...\n")

	for fields := range r {
		// mlib.DebugDump(fields)
		all = append(all, fields)
		for i := 0; i < len(fields); i++ {
			// Make sure widths has enough elements.
			if len(widths) < len(fields) {
				old := widths
				widths = make([]int, len(fields))
				copy(widths, old)
			}
			mlib.Debug("  index=%d, cur=%d\n", i, widths[i])
			widths[i] = max(widths[i], stringWidth(fields[i]))
		}
	}

	mlib.Debug("Read all lines\n")

	mlib.DebugDump(all)
	mlib.DebugDump(widths)

	outBuffer := bytes.Buffer{}

	for _, fields := range all {
		outBuffer.Reset()
		for i := 0; i < len(fields); i++ {
			if i > 0 {
				outBuffer.WriteString(" ")
			}
			w := stringWidth(fields[i])
			outBuffer.WriteString(fields[i])
			outBuffer.WriteString(strings.Repeat(" ", widths[i]-w))
		}
		out <- outBuffer.String()
	}
}
