package readutil_test

import (
	"bytes"
	"errors"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/nick-jones/astpath/internal/readutil"
)

func TestReadLine(t *testing.T) {
	testCases := []struct {
		name         string
		bytes        []byte
		offset       int64
		expectedLine []byte
		expectedErr  error
	}{
		{
			name:         "zero offset, no newline",
			bytes:        []byte("foo bar baz"),
			offset:       0,
			expectedLine: []byte("foo bar baz"),
		},
		{
			name:         "zero offset, with newline",
			bytes:        []byte("foo bar baz\nhello world"),
			offset:       0,
			expectedLine: []byte("foo bar baz"),
		},
		{
			name:         "first line, mid offset",
			bytes:        []byte("foo bar baz\nhello world"),
			offset:       6,
			expectedLine: []byte("foo bar baz"),
		},
		{
			name:         "first line, last offset",
			bytes:        []byte("foo bar baz\nhello world"),
			offset:       10,
			expectedLine: []byte("foo bar baz"),
		},
		{
			name:         "first line, offset is newline",
			bytes:        []byte("foo bar baz\nhello world"),
			offset:       11,
			expectedLine: []byte("foo bar baz"),
		},
		{
			name:         "second line, first offset",
			bytes:        []byte("foo bar baz\nhello world"),
			offset:       12,
			expectedLine: []byte("hello world"),
		},
		{
			name:         "second line, mid offset",
			bytes:        []byte("foo bar baz\nhello world"),
			offset:       17,
			expectedLine: []byte("hello world"),
		},
		{
			name:         "second line, last offset",
			bytes:        []byte("foo bar baz\nhello world"),
			offset:       22,
			expectedLine: []byte("hello world"),
		},
		{
			name:        "second line, beyond last offset",
			bytes:       []byte("foo bar baz\nhello world"),
			offset:      23,
			expectedErr: io.EOF,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			buf := bytes.NewReader(tc.bytes)
			res, err := readutil.ReadLine(buf, tc.offset)
			if tc.expectedErr == nil {
				require.NoError(tt, err)
				assert.Equal(tt, tc.expectedLine, res)
			} else {
				require.Error(tt, err)
				assert.True(tt, errors.Is(err, tc.expectedErr))
			}
		})
	}
}
