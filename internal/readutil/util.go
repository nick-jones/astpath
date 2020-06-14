package readutil

import (
	"errors"
	"io"
)

// ReadLine reads a full line from an arbitrary position. If the offset is a newline character the line preceding that
// character will be returned. Surely there is an easier way to achieve this?
func ReadLine(r io.ReaderAt, offset int64) ([]byte, error) {
	char := make([]byte, 1)

	// read forwards of the position until hitting a newline
	end := make([]byte, 0)
	for i := int64(0); ; i++ {
		if _, err := r.ReadAt(char, offset+i); err != nil {
			if i > 0 && errors.Is(err, io.EOF) {
				break
			}
			return nil, err
		}
		if char[0] == '\n' {
			break
		}
		end = append(end, char[0])
	}

	// read backwards from the position until hitting a newline
	start := make([]byte, 0)
	for off := offset - 1; off >= 0; off-- {
		if _, err := r.ReadAt(char, off); err != nil {
			return nil, err
		}
		if char[0] == '\n' {
			break
		}
		start = append([]byte{char[0]}, start...)
	}

	return append(start, end...), nil
}
