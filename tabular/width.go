package tabular

import (
	// "github.com/omakoto/mlib"
	"golang.org/x/text/width"
	// "unicode/utf8"
)

func stringWidth(s string) int {
	// numChars := utf8.RuneCountInString(s)

	bytes := []byte(s)

	w := 0
	for i := 0; i < len(s); {
		p, size := width.Lookup(bytes[i:])
		switch p.Kind() {
		case width.Neutral, width.EastAsianAmbiguous, width.EastAsianNarrow, width.EastAsianHalfwidth:
			w += 1
		case width.EastAsianWide, width.EastAsianFullwidth:
			w += 2
		}
		i += size
	}
	// mlib.Debug("  string=\"%s\", width=%d\n", s, w)

	return w
}
