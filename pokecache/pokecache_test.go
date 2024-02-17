package pokecache

import (
	"testing"
	"time"
)

func TestCreateCache(t *testing.T) {
	cache := NewCache(time.Millisecond)
	if cache.cache == nil {
		t.Error("Cache not created")
	}
}

func TestAddGetCache(t *testing.T) {
	cache := NewCache(time.Microsecond)

	cases := []struct {
		inputKey string
		inputVal []byte
	}{
		{
			inputKey: "key1",
			inputVal: []byte("val1"),
		},
		{
			inputKey: "key2",
			inputVal: []byte("val2"),
		},
		{
			inputKey: "",
			inputVal: []byte("val3"),
		},
	}

	for _, c := range cases {
		cache.Add(c.inputKey, c.inputVal)
		actual, ok := cache.Get(c.inputKey)
		if !ok {
			t.Errorf("%s not found", c.inputKey)
			continue
		}
		if string(actual) != string(c.inputVal) {
			t.Errorf("Expected %s but got %s", c.inputVal, actual)
		}
	}
}

func TestReapCache(t *testing.T) {
	interval := time.Millisecond * 10
	cache := NewCache(interval)

	key := "key1"
	cache.Add(key, []byte("val1"))

	time.Sleep(interval + time.Millisecond)

	_, ok := cache.Get(key)
	if ok {
		t.Errorf("%s not reaped", key)
	}

	key = "key2"
	cache.Add(key, []byte("val2"))
	time.Sleep(interval / 2)

	_, ok = cache.Get(key)
	if !ok {
		t.Errorf("%s reaped too soon", key)
	}
}
