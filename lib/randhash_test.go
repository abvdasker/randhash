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