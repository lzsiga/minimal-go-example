// test01_pkg/src/main.go

package main

import ("fmt"; "os"; "echo_args"; "minimath")

func sample() {
    fmt.Printf("Now a function is called from the same source\n")
}

func main() {
    fmt.Printf("Hello World, I am %s argc=%d\n",
        os.Args[0], len(os.Args))
    echo_args.EchoArgs()
    sample()
    sample2()
    fmt.Printf("gcd(56,24)=%d, lcm(56,24)=%d\n",
        minimath.Gcd(56, 24), minimath.Lcm(56,24))
}
