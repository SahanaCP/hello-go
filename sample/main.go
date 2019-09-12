package main

import (
	"fmt"

	wasm "github.com/wasmerio/go-ext-wasm/wasmer"
)

func main() {
	// Reads the WebAssembly module as bytes.
	bytes, _ := wasm.ReadBytes("./eval.wasm")

	// Instantiates the WebAssembly module.
	instance, _ := wasm.NewInstance(bytes)
	defer instance.Close()

	// Gets the `sum` exported function from the WebAssembly instance.
	eval := instance.Exports["eval"]

	// Calls that exported function with Go standard values. The WebAssembly
	// types are inferred and values are casted automatically.

	// result, _ := eval(5, '+', 10)

	fmt.Println(eval(5, '+', 10))
	fmt.Println(eval(5, '-', 10))
	fmt.Println(eval(5, '*', 10))
	fmt.Println(eval(15, '/', 4))
	fmt.Println(eval(50, '%', 9))
	fmt.Println(eval(2, '^', 5))
}
