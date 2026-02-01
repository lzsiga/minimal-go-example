/* gcd.go */

package main;

import "fmt";

func gcd(a int, b int) int {
    if a<0 { a= -a }
    if b<0 { b= -b }
    if a<b { a, b= b, a }
    for ;b!=0; {
        var remainder= a%b
        a, b= b, remainder
    }
    return a
}

func Test_gcd(a int, b int) {
    var g= gcd(a, b)
    fmt.Printf("gcd(%d,%d)=%d\n", a, b, g)
}

func main() {
    Test_gcd(0,0)
    Test_gcd(0,1)
    Test_gcd(1,96)
    Test_gcd(-64,96)
    Test_gcd(96,-64)
    Test_gcd(6,4)
}
