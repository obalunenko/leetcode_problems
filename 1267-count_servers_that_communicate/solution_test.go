package countserversthatcommunicate

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_countServers(t *testing.T) {
	type args struct {
		grid [][]int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "[[1,0],[0,1]] => 0",
			args: args{
				grid: [][]int{
					{1, 0},
					{0, 1},
				},
			},
			want: 0,
		},
		{
			name: "[[1,0],[1,1]] => 3",
			args: args{
				grid: [][]int{
					{1, 0},
					{1, 1},
				},
			},
			want: 3,
		},
		{
			name: "[[1,1,0,0],[0,0,1,0],[0,0,1,0],[0,0,0,1]] => 4",
			args: args{
				grid: [][]int{
					{1, 1, 0, 0},
					{0, 0, 1, 0},
					{0, 0, 1, 0},
					{0, 0, 0, 1},
				},
			},
			want: 4,
		},
		{
			name: "[[1,0,0,1,0],[0,0,0,0,0],[0,0,0,1,0]] => 3",
			args: args{
				grid: [][]int{
					{1, 0, 0, 1, 0},
					{0, 0, 0, 0, 0},
					{0, 0, 0, 1, 0},
				},
			},
			want: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := countServers(tt.args.grid)

			assert.Equal(t, tt.want, got)
		})
	}
}
