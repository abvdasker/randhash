package lib

import "testing"

func TestPutGet(t *testing.T) {
	h := NewRandHash()
	h.Put("key", "value")
	rawVal := h.Get("key")
	value := rawVal.(string)
	assertEqual(t, "value", value)
}

func TestPutPutGet(t *testing.T) {
	h := NewRandHash()
	h.Put("key", "value")
	h.Put("key", "value2")
	rawVal := h.Get("key")
	value := rawVal.(string)
	assertEqual(t, "value2", value)
}

func TestPutDeleteGet(t *testing.T) {
	h := NewRandHash()
	h.Put("key", "value")
	h.Delete("key")
	rawVal := h.Get("key")
	assertEqual(t, nil, rawVal)
}

func TestPutPutSize(t *testing.T) {
	h := NewRandHash()
	h.Put("key", "value")
	h.Put("key", "value2")
	h.Put("key2", "value3")
	assertEqual(t, 2, h.Size())
}

func TestPutPutDeleteSize(t *testing.T) {
	h := NewRandHash()
	h.Put("key", "value")
	h.Put("key2", "value2")
	h.Delete("key2")
	assertEqual(t, 1, h.Size())
}

func TestCollisions(t *testing.T) {
	h := NewRandHash()
	for i := 0; i < 50; i++ {
		h.Put(i, i)
		rawVal := h.Get(i)
		val := rawVal.(int)
		assertEqual(t, i, val)
	}
}

func pass(t *testing.T) {
	t.Log("pass")
}

func assertEqual(t *testing.T, first interface{}, second interface{}) {
	if first == second {
		pass(t)
	} else {
		t.Error("values not equal")
	}
}
