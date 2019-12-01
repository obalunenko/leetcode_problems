package addtwonumbers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_addTwoNumbers(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}

	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "add two lists with equal elements num and without overflow",
			args: args{
				l1: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val:  3,
							Next: nil,
						},
					},
				},
				l2: &ListNode{
					Val: 5,
					Next: &ListNode{
						Val: 6,
						Next: &ListNode{
							Val:  4,
							Next: nil,
						},
					},
				},
			},
			want: &ListNode{
				Val: 7,
				Next: &ListNode{
					Val: 0,
					Next: &ListNode{
						Val:  8,
						Next: nil,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := addTwoNumbers(tt.args.l1, tt.args.l2)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_addSlices(t *testing.T) {
	type args struct {
		n1 []int
		n2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "add two slices with equal elements num and without overflow",
			args: args{
				n1: []int{1, 2, 3},
				n2: []int{1, 2, 3},
			},
			want: []int{2, 4, 6},
		},
		{
			name: "add two lists with equal elements num and with overflow",
			args: args{
				n1: []int{9, 4, 7},
				n2: []int{1, 3, 9},
			},
			want: []int{0, 8, 6, 1},
		},
		{
			name: "add two lists with not equal elements num and with overflow",
			args: args{
				n1: []int{4, 7},
				n2: []int{1, 3, 9},
			},
			want: []int{5, 0, 0, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := addSlices(tt.args.n1, tt.args.n2)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestListNode_toNumber(t *testing.T) {
	type fields struct {
		Val  int
		Next *ListNode
	}
	tests := []struct {
		name   string
		fields fields
		want   []int
	}{
		{
			name: "convert list to slice",
			fields: fields{
				Val: 7,
				Next: &ListNode{
					Val: 0,
					Next: &ListNode{
						Val:  8,
						Next: nil,
					},
				},
			},
			want: []int{7, 0, 8},
		},
		{
			name: "convert list to slice",
			fields: fields{
				Val: 2,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  3,
						Next: nil,
					},
				},
			},
			want: []int{2, 4, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &ListNode{
				Val:  tt.fields.Val,
				Next: tt.fields.Next,
			}
			got := toNumber(l)

			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_makeList(t *testing.T) {
	type args struct {
		n []int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "",
			args: args{
				n: []int{7, 0, 8},
			},
			want: &ListNode{
				Val: 7,
				Next: &ListNode{
					Val: 0,
					Next: &ListNode{
						Val:  8,
						Next: nil,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := makeList(tt.args.n)

			assert.Equal(t, tt.want, got)
		})
	}
}
