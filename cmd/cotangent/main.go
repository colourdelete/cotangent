package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"nyiyui.ca/cotangent/lib"
)

func main() {
	err := main2()
	if err != nil {
		log.Fatal(err)
	}
}

func main2() error {
	r := lib.NewRenderer()
	source, err := io.ReadAll(os.Stdin)
	if err != nil {
		return fmt.Errorf("stdin read: %v", err)
	}
	out, err := r.RenderAndSanitize(source)
	if err != nil {
		return fmt.Errorf("render and sanitize: %v", err)
	}
	os.Stdout.Write(out)
	return nil
}
