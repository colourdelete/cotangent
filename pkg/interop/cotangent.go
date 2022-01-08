package main

import (
	"C"

	"nyiyui.ca/cotangent/lib"
)

var r = lib.NewRenderer()

//export RenderAndSanitize
func RenderAndSanitize(src *C.char, srcLen C.int) (res *C.char, resLen C.int) {
	out, err := r.RenderAndSanitize([]byte(C.GoStringN(src, srcLen)))
	if err != nil {
		msg := err.Error()
		res = C.CString(msg)
		resLen = C.int(len(msg))
		return
	}
	res = C.CString(string(out))
	resLen = C.int(len(out))
	return
}

func main() {}
