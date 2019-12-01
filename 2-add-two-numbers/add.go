package addtwonumbers

// ListNode represents singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

type list struct {
	head   *ListNode
	length int
}

func (l *list) Append(ln *ListNode) {
	if ln != nil {
		if l.head == nil {
			l.head = ln
			l.length++
			return
		}

		cur := l.head

		for cur.Next != nil {
			cur = cur.Next
		}

		cur.Next = ln
		l.length++
	}
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	res := addSlices(toNumber(l1), toNumber(l2))

	return makeList(res)
}

func makeList(n []int) *ListNode {
	var (
		list list
	)
	for _, v := range n {
		list.Append(&ListNode{
			Val:  v,
			Next: nil,
		})
	}

	return list.head
}

func toNumber(l *ListNode) []int {
	var n []int
	for l != nil {
		n = append(n, l.Val)
		l = l.Next
	}

	return n
}

// addSlices accepts numbers in reversed order and adds them.
// example:
// 342+256 = 598
// should be passed as:
// addSlices([]int{2,4,3}, []int{6,5,2})
func addSlices(n1, n2 []int) []int {
	var res []int
	var add int

	diff := len(n1) - len(n2)
	if diff < 0 {
		n1 = addZeros(n1, diff)
	}

	if diff > 0 {
		n2 = addZeros(n2, diff)
	}

	for i := 0; i < len(n1); i++ {
		sum := n1[i] + n2[i] + add

		if sum >= 10 {
			add = 1
			sum = sum % 10
		} else {
			add = 0
		}

		res = append(res, sum)
	}

	if add != 0 {
		res = append(res, add)
	}

	return res
}

func addZeros(n []int, diff int) []int {
	if diff < 0 {
		diff = diff * -1
	}

	x := make([]int, diff)
	n = append(n, x...)

	return n
}
