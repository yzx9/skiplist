package skiplist_test

import (
	"math/rand"
	"testing"

	"github.com/yzx9/skiplist"
)

func TestSkipList(t *testing.T) {
	t.Parallel()

	list := skiplist.New[int, int]()
	for i := 0; i < 10; i++ {
		v := i
		list.Insert(i, &v)
	}

	for i := 0; i < 10; i++ {
		v, err := list.Get(i)
		if err != nil {
			t.Errorf("skiplist.Get(%v) got error %s; excepted %v", i, err, v)
		} else if *v != i {
			t.Errorf("skiplist.Get(%v)=%v; excepted %v", i, v, i)
		}
	}

	if err := list.Delete(10); err != skiplist.KeyNotExist {
		t.Errorf("skiplist.Get(%v) got nothing; excepted error", 10)
	}

	for i := 0; i < 10; i += 2 {
		if err := list.Delete(i); err != nil {
			t.Errorf("skiplist.Get(%v) got error %s; excepted nothing", i, err)
		}
	}

	for i := 0; i < 10; i += 2 {
		if _, err := list.Get(i); err != skiplist.KeyNotExist {
			t.Errorf("skiplist.Get(%v) got %s; excepted error", i, err)
		}
	}

	for i := 1; i < 10; i += 2 {
		v, err := list.Get(i)
		if err != nil {
			t.Errorf("skiplist.Get(%v) got error %s; excepted %v", i, err, i)
		} else if *v != i {
			t.Errorf("skiplist.Get(%v) got %v; excepted %v", i, *v, i)
		}
	}

	for i := 8; i >= 0; i -= 2 {
		v := i
		list.Insert(i, &v)
	}

	for i := 0; i < 10; i++ {
		v, err := list.Get(i)
		if err != nil {
			t.Errorf("skiplist.Get(%v) got error %s; excepted %v", i, err, v)
		} else if *v != i {
			t.Errorf("skiplist.Get(%v)=%v; excepted %v", i, v, i)
		}
	}
}

func BenchmarkSkipListGet100(b *testing.B)    { benchmarkSkipListGet(b, 100) }
func BenchmarkSkipListGet1000(b *testing.B)   { benchmarkSkipListGet(b, 1000) }
func BenchmarkSkipListGet10000(b *testing.B)  { benchmarkSkipListGet(b, 10000) }
func BenchmarkSkipListGet100000(b *testing.B) { benchmarkSkipListGet(b, 100000) }

func benchmarkSkipListGet(b *testing.B, n int) {
	list := skiplist.New[int, int]()
	for i := 0; i < n; i++ {
		k := rand.Intn(i + 1)
		v := k
		list.Insert(k, &v)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for i := 0; i < n; i++ {
			_, _ = list.Get(n)
		}
	}
}
