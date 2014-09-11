package omap

import (
	"container/list"
	"fmt"
	"testing"
)

func traverseByLevel(m *Map, info string) {
	qcur := list.New()
	qcur.PushBack(m.root)
	level := 0
	cur_cnt, next_cnt := 1, 0

	fmt.Println(info, "traverse by level:")

	for qcur.Len() > 0 {
		n := qcur.Remove(qcur.Front()).(*node)
		cur_cnt--

		var left, right string = "/", "\\"
		if n.left == nil {
			left = ""
		}
		if n.right == nil {
			right = ""
		}
		color := "black"
		if n.red {
			color = "red"
		}
		fmt.Printf("%q=>%v%s%s(%q)", n.key, n.value,
			left, right, color)
		if cur_cnt > 0 {
			fmt.Print(", ")
		}

		if n.left != nil {
			qcur.PushBack(n.left)
			next_cnt++
		}
		if n.right != nil {
			qcur.PushBack(n.right)
			next_cnt++
		}
		if cur_cnt == 0 {
			fmt.Println()
			cur_cnt = next_cnt
			next_cnt = 0
			level++
		}
	}
	fmt.Println("level =", level, ", count=", m.length)
	fmt.Println()
}

func validate(root *node, t *testing.T) {
	if root == nil {
		return
	}
	if isRed(root) {
		t.Error("root is red")
	}
	var depth func(root *node) int
	depth = func(root *node) int {
		if root == nil {
			return 1
		}
		left := depth(root.left)
		righ := depth(root.right)
		if left != righ {
			t.Error("unbalanced tree:", left, righ)
		}
		if isRed(root.left) {
			return left
		} else {
			return left + 1
		}
	}
	if depth(root.left) != depth(root.right) {
		t.Error("rb tree unbalanced")
	}
}

func TestStringOrderedMap(t *testing.T) {
	m := NewStringOrderedMap()
	m.Insert("oa", 3)
	fmt.Println(m)
	if m.length != 1 {
		t.Error("failed insert")
	}
	validate(m.root, t)
}

func ExampleOmap() {
	m := NewStringOrderedMap()
	m.Insert("oa", 3)
	m.Insert("ob", 3)
	m.Insert("ab", 4)
	traverseByLevel(m, "new map with 3 nodes:")

	// m.Delete("obbc")
	// traverseByLevel(m, "after delete non-exist:")
	m.Insert("cab", 4)
	traverseByLevel(m, "insert one more:")

	key := "abc"
	val, found := m.Find(key)
	fmt.Printf("looking for %q => %v, %t\n", key, val, found)

	for k, v := range map[string]int{
		"zz":  5,
		"abc": 4, "cde": 5, "aef": 6, "oa": 4} {
		fmt.Printf("insert %q=>%v: %t\n", k, v, m.Insert(k, v))
	}
	m.Do(func(a, b interface{}) {
		fmt.Printf("%v=>%v,", a, b)
	})
	fmt.Println()
	fmt.Println(m)
	traverseByLevel(m, "last:")
	m.Delete("ab")
	traverseByLevel(m, "deleteMin:")
	// Output:
}

func BenchmarkOmap(b *testing.B) {
	fmt.Println("Hello")
}
