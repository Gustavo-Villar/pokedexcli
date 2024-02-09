package pokecache

import (
	"testing"
	"time"
)

func TestCreateCache(t *testing.T) {
	cache := NewCache(time.Millisecond)
	if cache.cache == nil {
		t.Error("cache is nil")
	}
}

func TestAddGetCache(t *testing.T) {
	cache := NewCache(time.Millisecond)

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

	for _, testCase := range cases {
		cache.Add(testCase.inputKey, []byte(testCase.inputVal))
		actual, ok := cache.Get(testCase.inputKey)
		if !ok {
			t.Errorf("%s not found\n", string(testCase.inputKey))
			continue
		}
		if string(actual) != string(testCase.inputVal) {
			t.Errorf("values %s and %s does not match", string(actual), string(testCase.inputVal))
			continue
		}
	}
}

func TestReap(t *testing.T) {
	interval := time.Millisecond * 10 //10 Milliseconds
	cache := NewCache(interval)

	keyOne := "key1"
	valueOne := "value1"

	cache.Add(keyOne, []byte(valueOne))

	time.Sleep(interval + time.Millisecond)

	_, ok := cache.Get(keyOne)
	if ok {
		t.Errorf("The key %s should have been reaped", keyOne)
	}

}

func TestReapFail(t *testing.T) {
	interval := time.Millisecond * 10 //10 Milliseconds
	cache := NewCache(interval)

	keyOne := "key1"
	valueOne := "value1"

	cache.Add(keyOne, []byte(valueOne))

	time.Sleep(interval / 2)

	_, ok := cache.Get(keyOne)
	if !ok {
		t.Errorf("The key %s should not have been reaped", keyOne)
	}

}
