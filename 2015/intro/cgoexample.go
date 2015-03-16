// +build OMIT

package main

import "log"
import "unsafe"

// START OMIT

// #include "stdlib.h"
// #include "pwd.h"
// #include "unistd.h"
import "C"

func main() {
	s := C.CString("Enter your password: ")
	pwd := C.getpass(s)
	log.Println("pwd is", C.GoString(pwd))
	C.free(unsafe.Pointer(s))
}

// STOP OMIT
