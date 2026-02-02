/* randtest.go */

package main

import ("fmt"; "math/rand")

func main() {
    var i int
    for i=0; i<10; i++ {
        var r = rand.Intn(6)    /* range: [0,5] */
        fmt.Printf ("%d\n", r)
    }
}
