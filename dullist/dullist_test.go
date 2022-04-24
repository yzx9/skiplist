package dullist_test

import (
	"testing"

	"github.com/yzx9/skiplist/dullist"
)

func TestDulList(t *testing.T) {
	t.Parallel()

	l1 := dullist.New[int, int]()
	for i := 0; i < 10; i++ {
		v := i
		l1.Insert(i, &v)
	}

	for i := 0; i < 10; i++ {
		v, err := l1.Get(i)
		if err != nil {
			t.Errorf("dullist.Get(%v) got error %s; excepted %v", i, err, v)
		} else if *v != i {
			t.Errorf("dullist.Get(%v)=%v; excepted %v", i, v, i)
		}
	}

	if err := l1.Delete(10); err != dullist.KeyNotExist {
		t.Errorf("dullist.Get(%v) got nothing; excepted error", 10)
	}

	for i := 0; i < 10; i += 2 {
		if err := l1.Delete(i); err != nil {
			t.Errorf("dullist.Get(%v) got error %s; excepted nothing", i, err)
		}
	}

	for i := 0; i < 10; i += 2 {
		if _, err := l1.Get(i); err != dullist.KeyNotExist {
			t.Errorf("dullist.Get(%v) got %s; excepted error", i, err)
		}
	}

	for i := 1; i < 10; i += 2 {
		v, err := l1.Get(i)
		if err != nil {
			t.Errorf("dullist.Get(%v) got error %s; excepted %v", i, err, i)
		} else if *v != i {
			t.Errorf("dullist.Get(%v) got %v; excepted %v", i, *v, i)
		}
	}

	for i := 8; i >= 0; i -= 2 {
		v := i
		l1.Insert(i, &v)
	}

	for i := 0; i < 10; i++ {
		v, err := l1.Get(i)
		if err != nil {
			t.Errorf("dullist.Get(%v) got error %s; excepted %v", i, err, v)
		} else if *v != i {
			t.Errorf("dullist.Get(%v)=%v; excepted %v", i, v, i)
		}
	}
}

func BenchmarkDullistGet100(b *testing.B)    { benchmarkDullistGet(b, 100) }
func BenchmarkDullistGet1000(b *testing.B)   { benchmarkDullistGet(b, 1000) }
func BenchmarkDullistGet10000(b *testing.B)  { benchmarkDullistGet(b, 10000) }
func BenchmarkDullistGet100000(b *testing.B) { benchmarkDullistGet(b, 100000) }

func benchmarkDullistGet(b *testing.B, n int) {
	list := dullist.New[int, int]()
	for i := 0; i < 100; i++ {
		v := i
		list.Insert(i, &v)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for i := 0; i < n; i++ {
			_, _ = list.Get(n)
		}
	}
}
