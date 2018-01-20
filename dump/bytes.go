package dump

import (
	"bytes"
	"encoding/binary"
	"io"
)

func Bytes(enc binary.ByteOrder, v interface{}) ([]byte, error) {
	var buf = bytes.Buffer{}
	if err := binary.Write(&buf, enc, v); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func BytesTo(r io.Writer, enc binary.ByteOrder, v interface{}) error {
	return binary.Write(r, enc, v)
}
