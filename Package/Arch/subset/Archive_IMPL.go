package subset

import "io"

type Arch interface {
	Compress(dir string, buf io.Writer) (err error)
	Decompress(file io.Reader, dst string) (err error)
}
