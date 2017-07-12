package main

import (
	"fmt"
	"github.com/abvdasker/randhash/lib"
)

func main() {
	hash := lib.NewRandHash()
	hash.Put("test", "first")
	hash.Put("test", "second")
	hash.Put("test2", "third")
	hash.Delete("test2")
	printString(hash, "test")
	printString(hash, "test2")
}

func printString(hash *lib.RandHash, key interface{}) {
	raw := hash.Get(key)
	value, _ := raw.(string)
	fmt.Println(value)
}
