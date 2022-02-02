package protocol

import (
	"bytes"
	"io"
)

type DataTypeWriter interface {
	Encode(w io.Writer) (int64, error)
}

func Marshal(values ...DataTypeWriter) ([]byte, error) {
	buffer := &bytes.Buffer{}

	for _, value := range values {
		if _, err := value.Encode(buffer); err != nil {
			return nil, err
		}
	}

	return buffer.Bytes(), nil
}
