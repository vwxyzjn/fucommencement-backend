package backend

import "fmt"

// A better println debugging method
func Describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
