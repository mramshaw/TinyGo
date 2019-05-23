package main

import (
	"syscall/js"
)

func main() {
}

//go:export reverseBytewise
func reverseBytewise(s string) string {
	r := []byte(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

//go:export update
func update() {
	document := js.Global().Get("document")
	inStr := document.Call("getElementById", "inString").Get("value").String()
	outStr := reverseBytewise(inStr)
	document.Call("getElementById", "outString").Set("value", outStr)
}
