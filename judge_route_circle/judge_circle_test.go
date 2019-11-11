package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type args struct {
	moves string
}

type want struct {
	circle bool
}

type test struct {
	id          int
	description string
	args        args
	want        want
}

func TestHamming(t *testing.T) {
	tests := hummingTests()
	for _, tc := range tests {
		tc := tc
		t.Run(fmt.Sprintf("Test%d: %s", tc.id, tc.description), func(t *testing.T) {
			actualResult := judgeCircle(tc.args.moves)
			assert.Equal(t, tc.want.circle, actualResult)
		})
	}
}

func hummingTests() []test {
	return []test{
		{
			id:          1,
			description: "test case from task - returned to start",
			args: args{
				moves: "UD",
			},
			want: want{
				circle: true,
			},
		},
		{
			id:          2,
			description: "test case from task - not returned to start",
			args: args{
				moves: "LL",
			},
			want: want{
				circle: false,
			},
		},
		{
			id:          3,
			description: "Not returned to start",
			args: args{
				moves: "ULUUDRLD",
			},
			want: want{
				circle: false,
			},
		},
		{
			id:          4,
			description: "Not returned to start on y",
			args: args{
				moves: "ULUDUDRLD",
			},
			want: want{
				circle: false,
			},
		},
		{
			id:          5,
			description: "Not returned to start on x",
			args: args{
				moves: "ULUUDRLDR",
			},
			want: want{
				circle: false,
			},
		},
		{
			id:          6,
			description: "Returned to start after Long walk",
			args: args{
				moves: "ULUUDRLDRD",
			},
			want: want{
				circle: true,
			},
		},
	}
}
