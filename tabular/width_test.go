package tabular

import (
	// "github.com/omakoto/tabular/tabular"
	"testing"
)

func expect(t *testing.T, expectedWidth int, s string) {
	actual := stringWidth(s)
	if actual != expectedWidth {
		t.Errorf("Fail expected width=%d, actual=%d for string=\"%s\"", expectedWidth, actual, s)
	}
}

func TestStringWidth(t *testing.T) {
	expect(t, 0, "")
	expect(t, 1, "a")
	expect(t, 3, "abc")
	expect(t, 2, "平")
	expect(t, 8, "平板粒子")
	expect(t, 21, "ｆｕｌｌｗｉｄｔｈ123")
	expect(t, 24, "ｆｕｌｌｗｉｄｔｈ１２３")
}
