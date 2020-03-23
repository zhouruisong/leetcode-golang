package main

import (
	"runtime/debug"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			debug.PrintStack()
		}
	}()

	value := 111
	zero := 0
	value = value / zero
}
