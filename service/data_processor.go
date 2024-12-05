package service

import (
	"bytes"
	"compress/zlib"
	"fmt"
	"io"
)

func UnzlibData(data []byte) ([]byte, error) {
	if len(data) < 4 || string(data[:4]) != "ZLIB" {
		return nil, fmt.Errorf("invalid zlib data header")
	}

	b := bytes.NewReader(data[4:])
	r, err := zlib.NewReader(b)
	if err != nil {
		return nil, err
	}
	defer r.Close()

	return io.ReadAll(r)
} 