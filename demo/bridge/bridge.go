package bridge

import "fmt"

// # TODO: can the compiler and linker take a relative path?
// #cgo CFLAGS: -IC:/code/go-static-linking/include
// #cgo LDFLAGS: C:/code/go-static-linking/build/libgb.a
// #include <junk.h>
import "C"

func Run() {
	fmt.Printf("Invoking c library...\n")
	C.x(10)
	fmt.Printf("Done\n")
}
