package pokecache

import (
	"testing"
	"time"
)

// TestCreateCache verifies the cache initialization.
func TestCreateCache(t *testing.T) {
	// Initialize a new cache with a very short interval for the test.
	cache := NewCache(time.Millisecond)
	// Ensure the cache map is not nil after initialization.
	if cache.cache == nil {
		t.Error("cache is nil")
	}
}

// TestAddGetCache verifies adding and retrieving items from the cache.
func TestAddGetCache(t *testing.T) {
	cache := NewCache(time.Millisecond)

	// Define test cases with keys and values to add to the cache.
	cases := []struct {
		inputKey string
		inputVal []byte
	}{
		{inputKey: "key1", inputVal: []byte("val1")},
		{inputKey: "key2", inputVal: []byte("val2")},
		{inputKey: "", inputVal: []byte("val3")}, // Test case with empty key.
	}

	// Iterate over each test case, adding items to the cache and verifying they can be retrieved.
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

// TestReap verifies that items expire from the cache as expected.
func TestReap(t *testing.T) {
	interval := time.Millisecond * 10 // 10 Milliseconds interval for item expiry.
	cache := NewCache(interval)

	// Add an item to the cache.
	keyOne := "key1"
	valueOne := "value1"
	cache.Add(keyOne, []byte(valueOne))

	// Wait for longer than the cache interval to allow the item to expire.
	time.Sleep(interval + time.Millisecond)

	// Verify the item has been expired and removed from the cache.
	_, ok := cache.Get(keyOne)
	if ok {
		t.Errorf("The key %s should have been reaped", keyOne)
	}
}

// TestReapFail verifies that items do not expire prematurely.
func TestReapFail(t *testing.T) {
	interval := time.Millisecond * 10 // 10 Milliseconds interval for item expiry.
	cache := NewCache(interval)

	// Add an item to the cache.
	keyOne := "key1"
	valueOne := "value1"
	cache.Add(keyOne, []byte(valueOne))

	// Wait for less than the cache interval to ensure the item should still be present.
	time.Sleep(interval / 2)

	// Verify the item has not been expired and is still retrievable.
	_, ok := cache.Get(keyOne)
	if !ok {
		t.Errorf("The key %s should not have been reaped", keyOne)
	}
}
