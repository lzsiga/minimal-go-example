/* switchtest.go */

package main

import "fmt"

func main() {
    var i int
    for i=1; i<=15; i++ {
        switch {
        case i%15 == 0:
            fmt.Printf("%d - FizzBuzz\n", i);
        case i%5 == 0:
            fmt.Printf("%d - Buzz\n", i);
        case i%3 == 0:
            fmt.Printf("%d - Fizz\n", i);
        default:
            fmt.Printf("%d\n", i);
        }
    }
}
