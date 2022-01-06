//go:build js && wasm

package main

import (
	"fmt"
	"syscall/js"

	"nyiyui.ca/cotangent/lib"
)

func main() {
	js.Global().Set("cotangent", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		handler := js.FuncOf(func(this js.Value, args2 []js.Value) interface{} {
			resolve := args2[0]
			reject := args2[1]

			go func() {
				errorConstructor := js.Global().Get("Error")
				if len(args) != 1 {
					fmt.Println("args:", args)
					errorObject := errorConstructor.New("must only have one string argument")
					reject.Invoke(errorObject)
				}
				r := lib.NewRenderer()
				sanitized, err := r.RenderAndSanitize([]byte(args[0].String()))
				if err != nil {
					errorObject := errorConstructor.New(err.Error())
					reject.Invoke(errorObject)
				} else {
					resolve.Invoke(js.ValueOf(string(sanitized)))
				}
			}()

			return nil
		})

		promiseConstructor := js.Global().Get("Promise")
		return promiseConstructor.New(handler)
	}))
	select {}
}
