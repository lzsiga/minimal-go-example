// test04_pkg/src/test04.go
// fixed size array

package main

import ("fmt")

const N = 5

func printVect (vect [N]int) {
    var j int
    for j=0; j<N; j++ {
        fmt.Printf ("%3d", vect[j])
        if j+1==N { fmt.Printf ("\n");
        } else    { fmt.Printf (" "); }
    }
}

func printArr (arr [N][N]int) {
    var i int
    var j int
    for i=0; i<N; i++ {
        for j=0; j<N; j++ {
            fmt.Printf ("%3d ", arr[i][j])
            if j+1==N { fmt.Printf ("\n");
            } else    { fmt.Printf (" ");  }
        }
    }
}

// parameters are passed by value,
// so here we use pointer

func fill (arr *[N][N]int) {
    var i int
    var j int
    for i=0; i<N; i++ {
        for j=0; j<N; j++ {
            arr[i][j] = 2*(i+1)+5*(j+1)
        }
    }
}

func prod (a [N][N]int, b [N][N]int) [N][N]int {
    var i int
    var j int
    var k int
    var c [N][N]int
    for i=0; i<N; i++ {
        for j=0; j<N; j++ {
            var sum= 0
            for k=0; k<N; k++ {
                sum += a[i][k]*b[k][j]
            }
            c[i][j]= sum
        }
    }
    return c;
}

func main () {
    var arr [N][N]int

    fill (&arr)
    fmt.Printf ("After 'fill':\n")
    printArr (arr)

// 'arr2' is an independent array copied from 'arr'
    var arr2= arr;
    arr2[0][0]= 0;
    arr2[1][1]= 0;
    arr2[2][2]= 0;

    fmt.Printf ("After modification 'arr':\n")
    printArr (arr)

    fmt.Printf ("And 'arr2':\n")
    printArr (arr2)

    var arrS = prod (arr,arr)
    fmt.Printf ("Square of 'arr':\n")
    printArr (arrS)
}
