package asn1
import (
    "io"
    "bytes"
)

type ReReader struct {
    re []byte       // bytes to read again
    buf []byte      // bytes read so far
    rd io.Reader    // source
}

func (b *ReReader) Read(p []byte) (n int, err error) {
    if len(b.re) > 0 {
        n = copy(p, b.re)
        b.re = b.re[n:]
    } else {
        n, err = b.rd.Read(p)
    }
    b.buf = append(b.buf, p[:n]...)
    return
}

func (b *ReReader) ReadFull(p []byte) error {
    _, err := io.ReadFull(b, p)
    return err
}

func (b *ReReader) ReadN(n int) ([]byte, error) {
    buf := make([]byte, n)
    return buf, b.ReadFull(buf)
}

func (b *ReReader) ReadSoFar() []byte {
    return b.buf
}


func (b *ReReader) ReadByte() (byte, error) {
    t := make([]byte, 1)
    err := b.ReadFull(t)
    return t[0], err
}

func (b *ReReader) Reset() {
    b.re = b.buf
    b.buf = []byte{}
}

func NewReReader(r io.Reader) *ReReader {
    return &ReReader{[]byte{}, []byte{}, r}
}

func NewReReaderBytes(b []byte) *ReReader {
    return NewReReader(bytes.NewBuffer(b))
}
