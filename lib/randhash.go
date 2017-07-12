package lib

import (
	"bytes"
	"encoding/gob"
	"hash/fnv"
	"math"
)

type RandHash struct {
	backing [][]*hashValue
	head    *cell
}

// backing doubly linked list of pointers to hashValues
// every insert adds to list and stores pointer to new value in list
// every delete snips entry from linked list
// get is noop

type hashValue struct {
	key   interface{}
	value interface{}
	ref   *cell
}

type cell struct {
	hVal *hashValue
	prev *cell
	next *cell
}

func NewRandHash() *RandHash {
	return &RandHash{
		backing: buildBacking(),
		head:    new(cell),
	}
}

func (h *RandHash) Put(key interface{}, value interface{}) {
	idx := h.backingIndex(key)
	row := h.backing[idx]
	currentHashValue := getValueFromRow(row, key)
	if currentHashValue == nil {
		backingValue := buildBackingValue(key, value)
		h.backing[idx] = append(row, backingValue)
		h.addCell(backingValue)
	} else {
		currentHashValue.key = key
		currentHashValue.value = value
	}
}

func (h *RandHash) Get(key interface{}) interface{} {
	row := h.backing[h.backingIndex(key)]
	hVal := getValueFromRow(row, key)
	if hVal == nil {
		return nil
	}
	return hVal.value
}

func (h *RandHash) Delete(key interface{}) interface{} {
	backingIdx := h.backingIndex(key)
	row := h.backing[backingIdx]
	for index, hVal := range row {
		if hVal.key == key {
			if hVal.ref.prev != nil {
				hVal.ref.prev.next = nil
			}
			if hVal.ref.next != nil {
				hVal.ref.next.prev = nil
			}
			updatedRow := deleteFromRow(row, index)
			h.backing[backingIdx] = updatedRow
			return hVal.value
		}
	}
	return nil
}

func deleteFromRow(row []*hashValue, index int) []*hashValue {
	newRow := make([]*hashValue, 0)
	for i, hVal := range row {
		if i != index {
			newRow = append(newRow, hVal)
		}
	}
	return newRow
}

func (h *RandHash) backingIndex(key interface{}) int {
	hashedKey := hashKey(key)
	normalizedKey := math.Mod(float64(hashedKey), float64(h.backingSize()))
	return int(normalizedKey)
}

func (h *RandHash) backingSize() uint64 {
	return uint64(len(h.backing))
}

func (h *RandHash) addCell(hVal *hashValue) {
	newCell := &cell{
		hVal: hVal,
	}
	newCell.next = h.head.next
	newCell.prev = h.head
	if h.head.next != nil {
		h.head.next.prev = newCell
	}
	h.head.next = newCell
	hVal.ref = newCell
}

func getValueFromRow(row []*hashValue, key interface{}) *hashValue {
	for _, hVal := range row {
		if hVal.key == key {
			return hVal
		}
	}
	return nil
}

func hashKey(key interface{}) uint64 {
	keyBytes := serializeKey(key)
	h := fnv.New64()
	h.Write(keyBytes)
	return h.Sum64()
}

func serializeKey(key interface{}) []byte {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	enc.Encode(key)
	return buf.Bytes()
}

func buildBacking() [][]*hashValue {
	backing := make([][]*hashValue, 10, 10)
	for i, _ := range backing {
		backing[i] = make([]*hashValue, 0)
	}
	return backing
}

func buildBackingValue(key interface{}, value interface{}) *hashValue {
	return &hashValue{
		key:   key,
		value: value,
	}
}
