/* ptrtest.c */

package main

import "fmt"

/* parameter passed by value */
func f1(invalue int) {
    invalue= 1968
}

/* parameter is a pointer-to-int */
func f2(byptr *int, newval int) {
    *byptr= newval
}

/* parameter is a pointer-pointer-to-int */
func f3(byptr **int, newval int) {
    **byptr= newval
}

/* parameter is a pointer-to-int-array */
func arrf2(byptr *[10]int) {
    var i int
    for i=0; i<10; i++ {   byptr [i]= 1970+i }    /* both syntax works */
    for i=0; i<10; i++ { (*byptr)[i]= 1980+i } /* I'd prefer this one */
}

/* parameter is a pointer-to-pointer-to-int-array */
func arrf3(byptr **[10]int) {
    var i int
    for i=0; i<10; i++ { (*byptr)[i]=  2010+i } /* both syntax works */
    for i=0; i<10; i++ { (**byptr)[i]= 2020+i } /* I'd prefer this one */
}

func main() {
    var num int = 1
    var expect int
    fmt.Printf("At start num=%d\n", num)
    f1 (num)
    fmt.Printf("After f1 num=%d\n", num)

    expect= 127
    f2 (&num, expect)
    fmt.Printf("After f2 num=%d (expected=%d)\n", num, expect)

    var pnum *int = &num
    expect= 255
    f3 (&pnum, expect)
    fmt.Printf("After f3 num=%d (expected=%d)\n", num, expect)

    var arr [10]int
    for i:=0; i<10; i++ { arr[i]= 100+i }

    fmt.Printf("Array before 'arr2f':")
    for i:=0; i<10; i++ { fmt.Printf (" %d", arr[i]) }
    fmt.Printf("\n")

    arrf2(&arr)

    fmt.Printf("Array after 'arrf2':")
    for i:=0; i<10; i++ { fmt.Printf (" %d", arr[i]) }
    fmt.Printf("\n")

    var arrPtr *[10]int = &arr
    arrf3(&arrPtr)
    fmt.Printf("Array after 'arrf3':")
    for i:=0; i<10; i++ { fmt.Printf (" %d", arr[i]) }
    fmt.Printf("\n")
}
