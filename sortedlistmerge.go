package student

type NodeI struct {
	Data int
	Next *NodeI
}

func SortedListMerge(n1 *NodeI, n2 *NodeI) *NodeI {
	n1 = ListSort(n1)
	n2 = ListSort(n2)

	if n1 == nil {
		return n2
	}
	if n2 == nil {
		return n1
	}
	if n1.Data <= n2.Data {
		n1.Next = SortedListMerge(n1.Next, n2)
		return n1
	} else {
		n2.Next = SortedListMerge(n1, n2.Next)
		return n2
	}
}

func ListSort(l *NodeI) *NodeI {
	if l == nil {
		return nil
	}
	current := l
	current.Next = ListSort(current.Next)
	if current.Next != nil && current.Data > current.Next.Data {
		current = SwapList(current)
	}
	return current
}
func SwapList(l *NodeI) *NodeI {
	prev := l
	next := l.Next
	current := next

	for next != nil && l.Data > next.Data {
		prev = next
		next = next.Next
	}
	prev.Next = l
	l.Next = next
	return current
}
