package main

import (
	"syscall/js"
)

// The empty main function is required (http://github.com/tinygo-org/tinygo/issues/186)
func main() {
}

func reverseBytewise(s string) string {
	r := []byte(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

// The next line looks like a Go pragma but is not; it's a TinyGo directive for WASM
//go:export update
func update() {
	document := js.Global().Get("document")
	inStr := document.Call("getElementById", "inString").Get("value").String()
	outStr := reverseBytewise(inStr)
	document.Call("getElementById", "outString").Set("value", outStr)
}
