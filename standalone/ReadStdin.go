/* ReadStdin.go */

package main

import "fmt"

func main() {
    var num, num1, num2, n int
    var err error
    var str string

    fmt.Printf("Enter two numbers ('Scanf'): ")
    n, err = fmt.Scanf("%d %d", &num1, &num2)
    if err!=nil {
        fmt.Printf("Error, I'm about to panic\n")
        panic(err)
    }
    fmt.Printf("Got %d values: num1=%d num2=%d\n", n, num1, num2)

    fmt.Printf("Enter some string ('Scanln'): ")
    fmt.Scanln(&str)
    fmt.Sscanf(str, "%d", &num);
    fmt.Printf("Got \"%s\"\nAs int: %d\n", str, num);
}
