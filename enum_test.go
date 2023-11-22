package enum_test

import (
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/vazquezger/enum"
)

func TestFluent(t *testing.T) {
	got := enum.
		Of([]int{1, 2, 3, 4, 5, 6, 7, 9, 10}).
		Filter(func(it int) bool { return it%2 == 0 }).
		Map(func(it int) int { return it * 2 }).
		ToSlice()

	want := []int{4, 8, 12, 20}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("want %v - got %v", want, got)
	}
}

func TestSlice(t *testing.T) {
	table := []struct {
		in   []int
		i, j int
		want []int
	}{
		{[]int{1, 2, 3, 4}, 0, 0, []int{1}},
		{[]int{1, 2, 3, 4}, 0, 1, []int{1, 2}},
		{[]int{1, 2, 3, 4}, 0, 2, []int{1, 2, 3}},
		{[]int{1, 2, 3, 4}, 0, 3, []int{1, 2, 3, 4}},
		{[]int{1, 2, 3, 4}, 1, 1, []int{2}},
		{[]int{1, 2, 3, 4}, 1, 2, []int{2, 3}},
		{[]int{1, 2, 3, 4}, 1, 3, []int{2, 3, 4}},
		{[]int{1, 2, 3, 4}, 3, 3, []int{4}},
	}

	for i := range table {
		got := enum.
			Of(table[i].in).
			Slice(table[i].i, table[i].j).
			ToSlice()

		want := table[i].want
		if !reflect.DeepEqual(got, want) {
			t.Errorf("%d: - want %v - got %v", i, want, got)
		}
	}
}
func TestFilter(t *testing.T) {
	table := []struct {
		in   []string
		f    func(it string) bool
		want []string
	}{
		{[]string{"Hello", "hellO", "hello", "heLO"}, func(it string) bool { return strings.ToLower(it) == it }, []string{"hello"}},
		{[]string{"Hello", "hellO", "hOello", "heLO"}, func(it string) bool { return strings.ToLower(it) == it }, []string{}},
	}

	for i := range table {
		got := enum.
			Of(table[i].in).
			Filter(table[i].f).
			ToSlice()

		want := table[i].want
		if !reflect.DeepEqual(got, want) {
			t.Errorf("%d: - want %v - got %v", i, want, got)
		}
	}
}

func TestFind(t *testing.T) {
	table := []struct {
		in    []int
		f     func(it int) bool
		want  int
		found bool
	}{
		{[]int{}, func(it int) bool { return it == 100 }, 0, false},
		{[]int{34, 6, 1, 0, 3, 3, 4}, func(it int) bool { return it == 3 }, 3, true},
		{[]int{34, 6, 1, 0, 3, 3, 4}, func(it int) bool { return it == 100 }, 0, false},
	}

	for i := range table {
		got, gotFound := enum.
			Of(table[i].in).
			Find(table[i].f)

		want := table[i].want
		wantFound := table[i].found
		if wantFound != gotFound {
			t.Errorf("%d: - want %v - got %v", i, wantFound, gotFound)
		}
		if wantFound && want != *got {
			t.Errorf("%d: - want %v - got %v", i, want, *got)
		}
	}
}

func TestMap(t *testing.T) {
	table := []struct {
		in   []int
		f    func(it int) int
		want []int
	}{
		{[]int{}, func(it int) int { return it * 2 }, []int{}},
		{[]int{9, 4, 0, 1}, func(it int) int { return it * 2 }, []int{18, 8, 0, 2}},
		{[]int{9, 4, 0, 1}, func(it int) int { return it * 4 }, []int{36, 16, 0, 4}},
	}

	for i := range table {
		got := enum.
			Of(table[i].in).
			Map(table[i].f).
			ToSlice()

		want := table[i].want
		if !reflect.DeepEqual(got, want) {
			t.Errorf("%d: - want %v - got %v", i, want, got)
		}
	}
}

func TestReduce(t *testing.T) {
	table := []struct {
		in    []int
		accum int
		f     func(it int, accum any) any
		want  int
	}{
		{[]int{}, 0, func(it int, accum any) any { return accum.(int) + it }, 0},
		{[]int{}, 10, func(it int, accum any) any { return accum.(int) + it }, 10},
		{[]int{9, 4, 0, 1}, 0, func(it int, accum any) any { return accum.(int) + it }, 14},
		{[]int{9, 4, 0, 1}, 100, func(it int, accum any) any { return accum.(int) + it }, 114},
	}

	for i := range table {
		got := enum.
			Of(table[i].in).
			Reduce(table[i].accum, table[i].f)

		want := table[i].want
		if want != got {
			t.Errorf("%d: - want %v - got %v", i, want, got)
		}
	}
}

func TestReduceWithIndex(t *testing.T) {
	table := []struct {
		in    []int
		accum int
		f     func(i int, it int, accum any) any
		want  int
	}{
		{[]int{}, 0, func(i int, it int, accum any) any { return accum.(int) + i + it }, 0},
		{[]int{}, 10, func(i int, it int, accum any) any { return accum.(int) + i + it }, 10},
		{[]int{9, 4, 0, 1}, 0, func(i int, it int, accum any) any { return accum.(int) + i + it }, 20},
		{[]int{9, 4, 0, 1}, 100, func(i int, it int, accum any) any { return accum.(int) + i + it }, 120},
	}

	for i := range table {
		got := enum.
			Of(table[i].in).
			ReduceWithIndex(table[i].accum, table[i].f)

		want := table[i].want
		if want != got {
			t.Errorf("%d: - want %v - got %v", i, want, got)
		}
	}

}

func TestEach(t *testing.T) {
	var str string

	table := []struct {
		in  []int
		f   func(elem int)
		out string
	}{
		{[]int{}, func(it int) { str += fmt.Sprintf("%d-", it) }, ""},
		{[]int{1}, func(it int) { str += fmt.Sprintf("%d-", it) }, "1-"},
		{[]int{2, 2, 8, 9}, func(it int) { str += fmt.Sprintf("%d-", it) }, "2-2-8-9-"},
		{[]int{50, 60}, func(it int) { str += fmt.Sprintf("%d-", it) }, "50-60-"},
	}

	for i := range table {
		str = ""
		enum.
			Of(table[i].in).
			Each(table[i].f)

		got := str
		want := table[i].out
		if want != got {
			t.Errorf("%d: - want %v - got %v", i, want, got)
		}
	}
}
