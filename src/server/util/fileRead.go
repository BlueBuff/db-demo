package util

import (
	"os"
	"fmt"
	"bufio"
	"io"
)

type FileReadUtil interface {
	Read() ([]byte, error)
}

type BufferFileReader struct {
	path string
}

func NewBufferFileReader(path string) FileReadUtil {
	read := new(BufferFileReader)
	read.path = path
	return read
}

func (reader *BufferFileReader) Read() ([]byte, error) {
	fi, err := os.Open(reader.path)
	if err != nil {
		return nil, fmt.Errorf("read file failed==>%s", err.Error())
	}
	r := bufio.NewReader(fi)
	buf := make([]byte, 1024, 1024)
	dest := make([]byte, 0, 1024)
	for {
		n, err := r.Read(buf)
		if err != nil && err != io.EOF {
			return nil, err
		}
		if n == 0 {
			break
		}
		dest = append(dest, buf[:n]...)
	}
	return dest, nil
}
