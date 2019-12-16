package parser

import "io"

var (
	rByte byte = 13 // the byte that corresponds to the '\r' rune.
	nByte byte = 10 // the byte that corresponds to the '\n' rune.
)

type reader struct {
	r io.Reader
}

func createReader(r io.Reader) io.Reader {
	return &reader{r: r}
}

func (r reader) Read(p []byte) (n int, err error) {
	n, err = r.r.Read(p)
	for i, b := range p {
		if b == rByte {
			p[i] = nByte
		}
	}
	return
}
