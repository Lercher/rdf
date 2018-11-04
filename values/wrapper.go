package values

import "io"

func toByteReader(r io.Reader) io.ByteReader {
	if br, ok := r.(io.ByteReader); ok {
		return br
	}
	return wrapper{r}
}

type wrapper struct {
	io.Reader
}

func (w wrapper) ReadByte() (byte, error) {
	b := []byte{0}
	_, err := w.Read(b)
	return b[0], err
}
