package protocol

import "io"

type DataTypeReader interface {
	Decode(r io.Reader) (int64, error)
}

func Unmarshal(r io.Reader, values ...DataTypeReader) error {
	for _, value := range values {
		if _, err := value.Decode(r); err != nil {
			return err
		}
	}

	return nil
}
