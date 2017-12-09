package tour

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r rot13Reader) Read(p []byte) (n int, err error) {
	n, err = r.r.Read(p)
	for i, c := range p {
		var start byte
		switch true {
		case 'a' <= c && c <= 'z':
			start = 'a' - 1
		case 'A' <= c && c <= 'Z':
			start = 'A' - 1
		default:
			continue
		}
		c = c + 13
		if c > start+26 {
			c = start + (c - (start + 26))
		}
		p[i] = c

	}
	return
}

// C61 ...
func C61(opts Options, args []string) int {
	s := strings.NewReader("Lbh penpexrq gur pbqr!")
	r := rot13Reader{s}
	_, err := io.Copy(os.Stdout, r)
	if err != nil {
		return 1
	}
	return 0
}
