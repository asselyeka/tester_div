package tests

import (
	"testing"

	student ".."
)

func TestSortedListMerge(t *testing.T) {
	tables := []struct {
		list1  []int
		list2  []int
		answer string
	}{
		{
			[]int{},
			[]int{},
			"<nil>",
		},
		{
			[]int{},
			[]int{2, 2, 4, 9, 12, 12, 19, 20},
			"2 -> 2 -> 4 -> 9 -> 12 -> 12 -> 19 -> 20 -> <nil>",
		},
		{
			[]int{4, 4, 6, 9, 13, 18, 20, 20},
			[]int{},
			"4 -> 4 -> 6 -> 9 -> 13 -> 18 -> 20 -> 20 -> <nil>",
		},
		{
			[]int{0, 7, 39, 92, 97, 93, 91, 28, 64},
			[]int{80, 23, 27, 30, 85, 81, 75, 70},
			"0 -> 7 -> 23 -> 27 -> 28 -> 30 -> 39 -> 64 -> 70 -> 75 -> 80 -> 81 -> 85 -> 91 -> 92 -> 93 -> 97 -> <nil>",
		},
		{
			[]int{0, 2, 11, 30, 54, 56, 70, 79, 99},
			[]int{23, 28, 38, 67, 67, 79, 95, 97},
			"0 -> 2 -> 11 -> 23 -> 28 -> 30 -> 38 -> 54 -> 56 -> 67 -> 67 -> 70 -> 79 -> 79 -> 95 -> 97 -> 99 -> <nil>",
		},
		{
			[]int{0, 3, 8, 8, 13, 19, 34, 38, 46},
			[]int{7, 39, 45, 53, 59, 70, 76, 79},
			"0 -> 3 -> 7 -> 8 -> 8 -> 13 -> 19 -> 34 -> 38 -> 39 -> 45 -> 46 -> 53 -> 59 -> 70 -> 76 -> 79 -> <nil>",
		},
	}

	for _, table := range tables {
		var link *student.NodeI
		var link2 *student.NodeI

		for _, d1 := range table.list1 {
			link = ListPushBack(link, d1)
		}

		for _, d2 := range table.list2 {
			link2 = ListPushBack(link2, d2)
		}

		res := GetList(student.SortedListMerge(link2, link))

		if res != table.answer {
			t.Errorf("SortedListMerge (%d, %d): [your answer - (%s)] != [(%s) - test's answer]", table.list1, table.list2, res, table.answer)
		}
	}
}
