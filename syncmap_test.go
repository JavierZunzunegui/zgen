package zgen_test

import (
	"sort"
	"testing"

	"github.com/JavierZunzunegui/zgen"
)

func TestSyncMap_StoreLoad(t *testing.T) {
	m := zgen.NewSyncMap[string, int]()

	m.Store("a", 42)
	v, ok := m.Load("a")
	if !ok || v != 42 {
		t.Errorf("Load hit: got (%v, %v), want (42, true)", v, ok)
	}

	v, ok = m.Load("missing")
	if ok || v != 0 {
		t.Errorf("Load miss: got (%v, %v), want (0, false)", v, ok)
	}
}

func TestSyncMap_Delete(t *testing.T) {
	m := zgen.NewSyncMap[string, int]()
	m.Store("a", 1)
	m.Delete("a")
	_, ok := m.Load("a")
	if ok {
		t.Error("expected key to be deleted")
	}
}

func TestSyncMap_LoadOrStore(t *testing.T) {
	m := zgen.NewSyncMap[string, int]()

	v, loaded := m.LoadOrStore("k", 10)
	if loaded || v != 10 {
		t.Errorf("absent key: got (%v, %v), want (10, false)", v, loaded)
	}

	v, loaded = m.LoadOrStore("k", 99)
	if !loaded || v != 10 {
		t.Errorf("present key: got (%v, %v), want (10, true)", v, loaded)
	}
}

func TestSyncMap_Range(t *testing.T) {
	m := zgen.NewSyncMap[string, int]()
	m.Store("a", 1)
	m.Store("b", 2)
	m.Store("c", 3)

	t.Run("collects all", func(t *testing.T) {
		var keys []string
		m.Range(func(k string, _ int) bool {
			keys = append(keys, k)
			return true
		})
		sort.Strings(keys)
		if len(keys) != 3 || keys[0] != "a" || keys[1] != "b" || keys[2] != "c" {
			t.Errorf("unexpected keys: %v", keys)
		}
	})

	t.Run("early exit", func(t *testing.T) {
		count := 0
		m.Range(func(_ string, _ int) bool {
			count++
			return false
		})
		if count != 1 {
			t.Errorf("expected 1 iteration, got %d", count)
		}
	})
}
