package longestsubstring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_lengthOfLongestSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "abcabcbb",
			args: args{
				s: "abcabcbb",
			},
			want: 3,
		},
		{
			name: "bbbbb",
			args: args{
				s: "bbbbb",
			},
			want: 1,
		},
		{
			name: "pwwkew",
			args: args{
				s: "pwwkew",
			},
			want: 3,
		},
		{
			name: "empty string",
			args: args{
				s: "",
			},
			want: 0,
		},
		{
			name: "empty string with spaces",
			args: args{
				s: " ",
			},
			want: 1,
		},
		{
			name: "one letter",
			args: args{
				s: "a",
			},
			want: 1,
		},
		{
			name: "no repeated",
			args: args{
				s: "au",
			},
			want: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lengthOfLongestSubstring(tt.args.s)
			assert.Equal(t, tt.want, got)
		})
	}
}
