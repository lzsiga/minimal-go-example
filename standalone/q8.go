/* q8.go */

package main

import "fmt"

type ChessTable struct {
    nset    int     /* filled values in 'row'; init with zero */
    row  [9]int     /* row-positions of 'nset' queens (values: 1..8) */
                    /* column-positions are not stored, queen#n stands in column#n */
}

/* Is the specfied field attacked by any of the 'nset' queens? */

func isFieldAttacked (table *ChessTable, col int, row int) bool {
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

/* find the first available field in a row,
   return value: true/false = found/not found
   retrow: the found row (1..8) or -1
 */

func findAvailableField (table *ChessTable, col int, minrow int, retrow *int) bool {
    var foundrow int = -1
    for r:=minrow; foundrow==-1 && r<=8; r++ {
        if !isFieldAttacked(table, col, r) {
            foundrow = r
        }
    }
    var found= (foundrow!=-1)
    *retrow= foundrow
    return found
}

func printTable (table *ChessTable) {
    fmt.Printf (" +--------+\n")
    for r:=8; r>=1; r-- {
        fmt.Printf (" |")
        for c:=1; c<=8; c++ {
            if table.nset>=c && table.row[c]==r {
                fmt.Printf("Q")
            } else {
                if (c-r)%2==0 {
                    fmt.Printf("▒")
                } else {
                    fmt.Printf(" ")
                }
            }
        }
        fmt.Printf ("|\n")
    }
    fmt.Printf (" +--------+\n\n")
}

func main () {
    var table ChessTable
    table.nset= 0
    table.row[1]= 0
    var col int = 1
    var row int = 0
    var nsolutions = 0
    var allDone bool = false

/* This algorithm is called 'backtrack';
   it means sometimes we enter dead-ends and have to turn back.
   For details see this: https://en.wikipedia.org/wiki/Backtracking
 */

    for ;!allDone; {
        var minrow int = table.row[col]+1
        var found bool = findAvailableField (&table, col, minrow, &row)
/* fmt.Printf("findAvailableField(col=%d,min=%d) returned %d; nset=%d\n", col, minrow, row, table.nset) */
        if found {                    /* ok, there is an available field in this column */
            table.nset= col
            table.row[col]= row
            if col<8 {                /* either we move right */
                col++
                table.row[col]= 0
            } else {                  /* or we have found a solution */
                nsolutions++
                fmt.Printf("Solution #%d\n", nsolutions)
                printTable (&table)   /* we go on, there might be other solution(s) */
            }
        } else {                      /* sadly, there is no available field in this column */
            if col>1 {                /* either we move left */
                col--
                table.nset= col-1
            } else {                  /* or there is no (more) solutions */
                allDone= true
            }
        }
    }
}
