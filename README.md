# asn1
Fork of encoding/asn1 library (Go v 1.4.1), customized to accept io.Reader as well as byte slices.



For more info, see the proposal on golang-dev: https://groups.google.com/forum/#!topic/golang-dev/8x6vjAJvTKQ

### Usage
```
func MarshalToWriter(val interface{}, w io.Writer) error
func UnmarshalFromReader(val interface{}, r io.Reader) error
```

### Installation

```sh
go get github.com/eipifi/asn1
```
