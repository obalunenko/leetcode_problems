package wordbreak

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_wordBreak(t *testing.T) {
	type args struct {
		s        string
		wordDict []string
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: `Return true because "leetcode" can be segmented as "leet code".`,
			args: args{
				s:        "leetcode",
				wordDict: []string{"leet", "code"},
			},
			want: true,
		},
		{
			name: `Return true because "applepenapple" can be segmented as "apple pen apple".`,
			args: args{
				s:        "applepenapple",
				wordDict: []string{"apple", "pen"},
			},
			want: true,
		},
		{
			name: `Return false cause could not segment`,
			args: args{
				s:        "catsandog",
				wordDict: []string{"cats", "dog", "sand", "and", "cat"},
			},
			want: false,
		},
		{
			name: `Return true cause could be segmented, the same letter`,
			args: args{
				s:        "aaaaaaa",
				wordDict: []string{"aaaa", "aaa"},
			},
			want: true,
		},
		{
			name: `Return false for empty word`,
			args: args{
				s:        "",
				wordDict: []string{"aaaa", "aaa"},
			},
			want: false,
		},
		{
			name: `Return false for empty dict`,
			args: args{
				s:        "aasad s",
				wordDict: nil,
			},
			want: false,
		},
		{
			name: `Return true for repeated`,
			args: args{
				s:        "goalspecial",
				wordDict: []string{"go", "goal", "goals", "special"},
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := wordBreak(tt.args.s, tt.args.wordDict)
			assert.Equal(t, tt.want, got)
		})
	}
}
