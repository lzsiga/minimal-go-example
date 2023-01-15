// src/echo_args/echo_args.go

package echo_args

import ("fmt"; "os")

// upper-case 'E' int EchoArgs mean exported symbol
func EchoArgs() {
    for i := range os.Args {
        fmt.Printf("%4d %s\n", i, os.Args[i])
    }
}
