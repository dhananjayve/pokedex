package pokecache

import (
	"testing"
	"time"
)

func TestCreateCache(t *testing.T) {
	ttime := time.Millisecond * 10
	cache := NewCache(ttime)
	if cache.cache == nil {
		t.Error("cache is nil")
	}
}

func TestAddGetCache(t *testing.T) {
	ttime := time.Millisecond * 10
	cache := NewCache(ttime)

	cases := []struct {
		inputKey string
		inputVal []byte
	}{
		{
			inputKey: "key1",
			inputVal: []byte("value1"),
		},
		{
			inputKey: "key2",
			inputVal: []byte("value2"),
		},
	}
	for _, c := range cases {
		cache.Add(c.inputKey, c.inputVal)
		actual, ok := cache.Get(c.inputKey)

		if !ok {
			t.Error("key not found")
			continue
		}
		if string(actual) != string(c.inputVal) {
			t.Error("value do not match")
			continue
		}
	}

}

func TestPurge(t *testing.T) {
	ttime := time.Millisecond * 10
	cache := NewCache(ttime)

	keyOne := "key1"
	cache.Add(keyOne, []byte("value1"))
	time.Sleep(ttime + time.Millisecond)
	_, ok := cache.Get(keyOne)
	if ok {
		t.Error("key should be purged")
	}
}

func TestPurgeFail(t *testing.T) {
	ttime := time.Millisecond * 10
	cache := NewCache(ttime)

	keyOne := "key1"
	cache.Add(keyOne, []byte("value1"))
	time.Sleep(ttime / 2)
	_, ok := cache.Get(keyOne)
	if !ok {
		t.Error("key should exists")
	}
}
