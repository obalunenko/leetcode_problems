package palindromic

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_longestPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Example 1",
			args: args{
				s: "babad",
			},
			want: "bab",
		},
		{
			name: "Example 2",
			args: args{
				s: "cbbd",
			},
			want: "bb",
		},
		{
			name: "Example 3",
			args: args{
				s: "a",
			},
			want: "a",
		},
		{
			name: "Example 4",
			args: args{
				s: "ac",
			},
			want: "a",
		},
		{
			name: "Example 5",
			args: args{
				s: "bb",
			},
			want: "bb",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := longestPalindrome(tt.args.s)

			assert.Equal(t, tt.want, got)
		})
	}
}
