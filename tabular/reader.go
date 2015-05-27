package tabular

import (
	"bufio"
	"encoding/csv"
	"io"
	"log"
	"regexp"
	"strings"
)

type LineReader interface {
	read(io.Reader) []string
	close()
}

func TsvReader(origReader io.ReadCloser) <-chan []string {
	return xsvReader(origReader, '\t')
}

func CsvReader(origReader io.ReadCloser) <-chan []string {
	return xsvReader(origReader, ',')
}

func xsvReader(origReader io.ReadCloser, sep rune) <-chan []string {
	r := csv.NewReader(origReader)
	r.FieldsPerRecord = -1
	r.Comma = sep
	out := make(chan []string)

	go func() {
		defer close(out)
		defer origReader.Close()
		for {
			rec, err := r.Read()
			if err == io.EOF {
				return
			}
			if err != nil {
				log.Printf("Error reading input: %s\n", err)
				return
			}
			out <- rec
		}
	}()

	return out
}

func RegexpSeparatedReader(origReader io.Reader, pattern string) <-chan []string {
	re := regexp.MustCompile(pattern)
	r := bufio.NewReader(origReader)
	out := make(chan []string)

	go func() {
		defer close(out)
		for {
			line, err := r.ReadString('\n')
			line = strings.TrimRight(line, "\r\n")
			if line != "" || err == nil {
				out <- re.Split(line, -1)
			}
			if err == io.EOF {
				return
			}
			if err != nil {
				log.Printf("Error reading input: %s\n", err)
				return
			}
		}
	}()

	return out
}
