package readline

import (
	"io"
	"bufio"
	"bytes"
)

func ReadLine(reader io.Reader, f func(string)) {
	buf := bufio.NewReader(reader)
	line, err := buf.ReadBytes('\n')
	for err == nil {
		line = bytes.TrimRight(line, "\n")
		if line[len(line)-1] == 13 { //'\r'
			line = bytes.TrimRight(line, "\r")
		}
		if len(line) > 0 {
			f(string(line))
		}
		line, err = buf.ReadBytes('\n')
	}

	if len(line) > 0 {
		f(string(line))
	}
}
