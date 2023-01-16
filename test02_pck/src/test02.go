/* test02_pkg/src/test02.go */

package main

import ("fmt"; "os"; "unsafe")

/*
void dlsym_test(const char *name);
void free(void *);
#cgo LDFLAGS: -ldl
*/
import "C"

func dlsymTest(symname string) {
    cstr := C.CString(symname)
    C.dlsym_test(cstr)
    C.free(unsafe.Pointer(cstr))
}

func main() {
    fmt.Printf("Hello World, I am %s\n", os.Args[0])
    dlsymTest("printf");
    dlsymTest("malloc");
    dlsymTest("atexit");
    dlsymTest("__cxa_atexit");
}
