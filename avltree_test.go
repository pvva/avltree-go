package avltree

import (
	"strings"
	"testing"
)

type ComparableString string

func (sv ComparableString) CompareTo(i interface{}) int {
	if ts, ok := i.(ComparableString); ok {
		return strings.Compare(string(sv), string(ts))
	}

	return -1
}

func TestComparableString(t *testing.T) {
	svA1 := ComparableString("A")
	svA2 := ComparableString("A")
	svB := ComparableString("B")

	if svA1.CompareTo(svA2) != 0 {
		t.Fatal("CompareTo failed")
	}

	if svA1.CompareTo(svB) != -1 || svB.CompareTo(svA1) != 1 {
		t.Fatal("CompareTo failed")
	}
}

func TestAvlTreeInsert2(t *testing.T) {
	svA := ComparableString("A")
	svB := ComparableString("B")

	tree := NewAvlTree()
	tree.Insert(svA)
	tree.Insert(svB)

	if tree.Height() != 1 {
		t.Fatal("Tree height is incorrect")
	}
	tree.Traverse(func(comparable Comparable, i int) bool {
		if (i == 0 && comparable.CompareTo(svB) != 0) || (i == 1 && comparable.CompareTo(svA) != 0) {
			t.Fatal("Incorrect tree structure")
		}

		return true
	})
}

func TestAvlTreeInsert4(t *testing.T) {
	svA := ComparableString("A")
	svB := ComparableString("B")
	svC := ComparableString("C")

	tree := NewAvlTree()
	tree.Insert(svA)
	tree.Insert(svB)
	tree.Insert(svC)
	tree.Insert(svA)

	// tree looks like this:
	//       B        level 2
	//    A           level 1
	//     A   C       level 0

	if tree.Height() != 2 {
		t.Fatal("Tree height is incorrect")
	}
	tree.Traverse(func(comparable Comparable, i int) bool {
		if (i == 0 && !(comparable.CompareTo(svC) == 0 || comparable.CompareTo(svA) == 0)) ||
			(i == 1 && comparable.CompareTo(svA) != 0) ||
			(i == 2 && comparable.CompareTo(svB) != 0) {

			t.Fatal("Incorrect tree structure")
		}

		return true
	})
}

func TestAvlTreeRemove4(t *testing.T) {
	svA := ComparableString("A")
	svB := ComparableString("B")
	svC := ComparableString("C")

	tree := NewAvlTree()
	tree.Insert(svA)
	tree.Insert(svB)
	tree.Insert(svC)
	tree.Insert(svA)

	// tree looks like this:
	//       B        level 2
	//    A    C      level 1
	//     A          level 0

	tree.Remove(svA)

	// tree should look like this:
	//      B      level 1
	//        C    level 0
	if tree.Height() != 1 {
		t.Fatal("Tree height is incorrect")
	}
	tree.Traverse(func(comparable Comparable, i int) bool {
		if (i == 0 && comparable.CompareTo(svC) != 0) || (i == 1 && comparable.CompareTo(svB) != 0) {
			t.Fatal("Incorrect tree structure")
		}

		return true
	})
}

func TestAvlTreeRemove5(t *testing.T) {
	svA := ComparableString("A")
	svB := ComparableString("B")
	svC := ComparableString("C")
	svD := ComparableString("D")
	svE := ComparableString("E")

	tree := NewAvlTree()
	tree.Insert(svA)
	tree.Insert(svB)
	tree.Insert(svC)
	tree.Insert(svD)
	tree.Insert(svE)

	// tree looks like this:
	//      B        level 3
	//    A   C      level 2
	//          D    level 1
	//            E  level 0

	tree.Remove(svA)
	tree.Remove(svD)

	// tree should look like this:
	//      C      level 1
	//    B   E    level 0
	if tree.Height() != 1 {
		t.Fatal("Tree height is incorrect")
	}
	tree.Traverse(func(comparable Comparable, i int) bool {
		if (i == 0 && !(comparable.CompareTo(svB) == 0 || comparable.CompareTo(svE) == 0)) ||
			(i == 1 && comparable.CompareTo(svC) != 0) {

			t.Fatal("Incorrect tree structure")
		}

		return true
	})
}

func TestAvlTreeRemove6(t *testing.T) {
	svA := ComparableString("A")
	svB := ComparableString("B")
	svC := ComparableString("C")

	tree := NewAvlTree()
	tree.Insert(svA)
	tree.Insert(svB)
	tree.Insert(svC)
	tree.Insert(svA)
	tree.Insert(svA)
	tree.Insert(svA)

	// tree looks like this:
	//      A        level 2
	//    A   B      level 1
	//  A    A C     level 0

	tree.Remove(svA)

	// tree should look like this:
	//      B      level 1
	//        C    level 0
	if tree.Height() != 1 {
		t.Fatal("Tree height is incorrect")
	}
	tree.Traverse(func(comparable Comparable, i int) bool {
		if (i == 0 && comparable.CompareTo(svC) != 0) ||
			(i == 1 && comparable.CompareTo(svB) != 0) {

			t.Fatal("Incorrect tree structure")
		}

		return true
	})
}

func TestAvlTreeRemove7(t *testing.T) {
	svA := ComparableString("A")
	svB := ComparableString("B")
	svC := ComparableString("C")
	svD := ComparableString("D")
	svE := ComparableString("E")
	svF := ComparableString("F")
	svG := ComparableString("G")

	tree := NewAvlTree()
	tree.Insert(svA)
	tree.Insert(svB)
	tree.Insert(svC)
	tree.Insert(svD)
	tree.Insert(svE)
	tree.Insert(svF)
	tree.Insert(svG)

	// tree looks like this:
	//       D        level 2
	//    B    F      level 1
	//   A C  E G     level 0

	tree.Remove(svF)

	// tree should look like this:
	//       D        level 2
	//    B    G      level 1
	//   A C  E       level 0
	if tree.Height() != 2 {
		t.Fatal("Tree height is incorrect")
	}
	tree.Traverse(func(comparable Comparable, i int) bool {
		if (i == 0 && !(comparable.CompareTo(svA) == 0 || comparable.CompareTo(svC) == 0 || comparable.CompareTo(svE) == 0)) ||
			(i == 1 && !(comparable.CompareTo(svB) == 0 || comparable.CompareTo(svG) == 0)) ||
			(i == 2 && comparable.CompareTo(svD) != 0) {

			t.Fatal("Incorrect tree structure")
		}

		return true
	})
}

func TestAvlTreeTraverse(t *testing.T) {
	tree := NewAvlTree()
	tree.Insert(ComparableString("A"))
	tree.Insert(ComparableString("G"))
	tree.Insert(ComparableString("B"))
	tree.Insert(ComparableString("F"))
	tree.Insert(ComparableString("D"))
	tree.Insert(ComparableString("C"))
	tree.Insert(ComparableString("E"))

	actual := []string{}

	tree.Traverse(func(comparable Comparable, i int) bool {
		cs := comparable.(ComparableString)
		actual = append(actual, string(cs))

		return true
	})

	expected := []string{"A", "B", "C", "D", "E", "F", "G"}

	if len(actual) != len(expected) {
		t.Fatal("Traverse failed")
	}

	for idx, v := range expected {
		if v != actual[idx] {
			t.Fatal("Traverse failed")
		}
	}
}

func TestAvlTreeTraverseInterrupt(t *testing.T) {
	tree := NewAvlTree()
	tree.Insert(ComparableString("A"))
	tree.Insert(ComparableString("G"))
	tree.Insert(ComparableString("B"))
	tree.Insert(ComparableString("F"))
	tree.Insert(ComparableString("D"))
	tree.Insert(ComparableString("C"))
	tree.Insert(ComparableString("E"))

	actual := []string{}

	tree.Traverse(func(comparable Comparable, i int) bool {
		cs := comparable.(ComparableString)
		actual = append(actual, string(cs))

		// stop in the middle
		return string(cs) != "D"
	})

	expected := []string{"A", "B", "C", "D"}

	if len(actual) != len(expected) {
		t.Fatal("Traverse failed")
	}

	for idx, v := range expected {
		if v != actual[idx] {
			t.Fatal("Traverse failed")
		}
	}
}
