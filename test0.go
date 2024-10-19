package main

import (
	"errors"
	"testing"
	"unicode/utf8"
)

var ErrInvalidUTF8 = errors.New("invalid utf8")

func GetUTFLength(input []byte) (int, error) {
	if !utf8.Valid(input) {
		return 0, ErrInvalidUTF8
	}

	return utf8.RuneCount(input), nil
}

func TestGetUTFLength(t *testing.T) {
	cases := []struct {
		name   string
		values string
		want   int
		err    error
	}{
		{
			name:   "test1",
			values: "qwerty",
			want:   6,
			err:    nil,
		},
		{
			name:   "test2",
			values: "e",
			want:   1,
			err:    nil,
		},
		{
			name:   "test3",
			values: "qazxderty ",
			want:   10,
			err:    nil,
		},
		{name: "test4",
			values: string([]byte{0xff, 0xfe}),
			want:   0,
			err:    ErrInvalidUTF8,
		}}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			got, err := GetUTFLength([]byte(tc.values))
			if got != tc.want || (err != nil && err.Error() != tc.err.Error()) {
				t.Errorf("GetUTFLength(%v) = %v, %v; want %v, %v", tc.values, got, err, tc.want, tc.err)
			}
		})
	}
}
