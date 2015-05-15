package tabular

import (
	"strings"
)

func Tabular(r <-chan []string) <-chan string {
	out := make(chan string)

	go func() {
		defer close(out)

		for fields := range r {
			out <- strings.Join(fields, " ")
		}

	}()

	return out
}
