package main

import (
	"bloom-filter/bloom"	// your path may very
	"fmt"
)


// Driver program
func main() {
	usernameStore := bloom.Init(1000, 0.01)

	usernameStore.Add("foo")
	usernameStore.Add("bar")
	usernameStore.Add("baz")

	fmt.Printf("%t\n", usernameStore.Contains("foo"))				// true
	fmt.Printf("%t\n", usernameStore.Contains("ishouldbefalse"))	// false
	fmt.Printf("%t\n", usernameStore.Contains("metoo"))				// false
	fmt.Printf("%t\n", usernameStore.Contains("bar"))				// true
	fmt.Printf("%t\n", usernameStore.Contains("idon'tbelonghere"))	// false
	fmt.Printf("%t\n", usernameStore.Contains("foo"))				// true

}
