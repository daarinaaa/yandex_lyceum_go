package main

import "testing"

func TestDeleteVowels(t *testing.T) {
	cases := []struct {
		name   string
		values string
		want   string
	}{
		{
			name:   "test1",
			values: "qwertyuiop1",
			want:   "qwrtyp1",
		},

		{
			name:   "test2",
			values: "ma ma ma",
			want:   "m m m",
		},
		{
			name:   "test3",
			values: "apsdfa",
			want:   "psdf",
		}}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			got := DeleteVowels(tc.values)
			if got != tc.want {
				t.Errorf("DeleteVowels(%v) = %v; want %v", tc.values, got, tc.want)
			}
		})
	}
}
