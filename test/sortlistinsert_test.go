package tests

import (
	"strconv"
	"testing"

	student ".."
)

func TestSortListInsert(t *testing.T) {
	tables := []struct {
		list    []int
		dataRef int
		answer  string
	}{
		{
			[]int{0},
			39,
			"0 -> 39 -> <nil>",
		},
		{
			[]int{0, 1, 2, 3, 4, 5, 24, 25, 54},
			33,
			"0 -> 1 -> 2 -> 3 -> 4 -> 5 -> 24 -> 25 -> 33 -> 54 -> <nil>",
		},
		{
			[]int{0, 2, 18, 33, 37, 37, 39, 52, 53, 57},
			53,
			"0 -> 2 -> 18 -> 33 -> 37 -> 37 -> 39 -> 52 -> 53 -> 53 -> 57 -> <nil>",
		},
		{
			[]int{0, 5, 18, 24, 28, 35, 42, 45, 52},
			52,
			"0 -> 5 -> 18 -> 24 -> 28 -> 35 -> 42 -> 45 -> 52 -> 52 -> <nil>",
		},
		{
			[]int{0, 12, 20, 23, 23, 24, 30, 41, 53, 57, 59},
			38,
			"0 -> 12 -> 20 -> 23 -> 23 -> 24 -> 30 -> 38 -> 41 -> 53 -> 57 -> 59 -> <nil>",
		},
	}

	for _, table := range tables {
		var link *student.NodeI

		for _, data := range table.list {
			link = ListPushBack(link, data)
		}

		link = student.SortListInsert(link, table.dataRef)

		res := GetList(link)

		if res != table.answer {
			t.Errorf("SortInsertList (%d, %d): [your answer - (%s)] != [(%s) - test's answer]", table.list, table.dataRef, res, table.answer)
		}
	}
}

func GetList(l *student.NodeI) string {
	list := ""
	it := l
	for it != nil {
		list += strconv.Itoa(it.Data) + " -> "
		it = it.Next
	}
	list += "<nil>"
	return list
}

func ListPushBack(l *student.NodeI, data int) *student.NodeI {
	n := &student.NodeI{Data: data}

	if l == nil {
		return n
	}
	iterator := l
	for iterator.Next != nil {
		iterator = iterator.Next
	}
	iterator.Next = n
	return l
}
