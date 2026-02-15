/* q8.go */

package main

import "fmt"

type ChessTable struct {
    nset    int     /* filled values in 'row'; init with zero */
    row  [9]int     /* row-positions of 'nset' queens (values: 1..8) */
                    /* column-positions are not stored, queen#n stands in column#n */
}

/* Is the specfied field attacked by any of the 'nset' queens? */

func isFieldAttacked (table *ChessTable, row int, col int) bool {
    if row<1 || row>8 || col<1 || col>8 {
        panic("invalid parameter(s) in 'isFieldAttacked'\n")
    }
    var fAttack bool = false
    var i int
    for i=1; i<=table.nset && !fAttack; i++ {
        if col==i || row==table.row[i] ||
           col-i == row-table.row[i] ||
           col-i == table.row[i]-row {
            fAttack= true
        }
    }
    return fAttack
}

func printTable (table *ChessTable) {
    fmt.Printf (" +--------+\n")
    for r:=8; r>=1; r-- {
        fmt.Printf (" |")
        for c:=1; c<=8; c++ {
            if table.nset>=c && table.row[c]==r {
                fmt.Printf("Q")
            } else {
                fmt.Printf(" ")
            }
        }
        fmt.Printf ("|\n")
    }
    fmt.Printf (" +--------+\n\n")
}

func Test1 () {
    var table ChessTable
    table.nset= 2
    table.row[1]= 1
    table.row[2]= 8
    printTable (&table)

    fmt.Printf (" +--------+\n")
    for r:=8; r>=1; r-- {
        fmt.Printf (" |")
        for c:=1; c<=8; c++ {
            if isFieldAttacked (&table, r, c) {
                fmt.Printf("*")
            } else {
                fmt.Printf(" ")
            }
        }
        fmt.Printf ("|\n")
    }
    fmt.Printf (" +--------+\n\n")
}

func main () {
    var table ChessTable
    table.nset= 0
    var col int = 1
    var row int = 0
    var allDone bool = false

    for ;!allDone; {
        var found bool = false
        for row= table.row[col]; row<=8; row++ {
            if !isFieldAttacked (table, row, col) { found= true }
        }
        if found {
            table.nset= col
            table.row[col]= row
            if col==8 {
                printTable (table)
            } else {
                col++
                table.row[col]= 0
            }
        } else {
            if col==1 {
                allDone= true
            } else {
                col--
                table.nset= col
            }
        }
    }
}
