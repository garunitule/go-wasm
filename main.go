package main

import (
	"syscall/js"
)

func main() {
	c := make(chan struct{}, 0)
	js.Global().Set("hello", js.FuncOf(hello))
	<-c
}

func hello(this js.Value, p []js.Value) interface{} {
	js.Global().Get("document").Call("getElementById", "output").Set("innerText", "Hello, WebAssembly!")
	// js console.log p
	js.Global().Get("console").Call("log", p)

	return nil
}
