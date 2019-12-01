package twosum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_twoSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "sum exist",
			args: args{
				nums:   []int{2, 7, 11, 15},
				target: 9,
			},
			want: []int{0, 1},
		},
		{
			name: "empty slice",
			args: args{
				nums:   nil,
				target: 9,
			},
			want: nil,
		},
		{
			name: "one element",
			args: args{
				nums:   []int{9},
				target: 9,
			},
			want: []int{0},
		},
		{
			name: "no target sum",
			args: args{
				nums:   []int{9,2,99},
				target: 1000,
			},
			want: nil,
		},
		{
			name: "3,2,4",
			args: args{
				nums:   []int{3,2,4},
				target: 6,
			},
			want: []int{1,2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := twoSum(tt.args.nums, tt.args.target)
			assert.Equal(t, tt.want, got)
		})
	}
}
