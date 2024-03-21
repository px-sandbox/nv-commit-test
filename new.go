package main

import "fmt"

func main() {
    number := 5
    result := factorial(number)
    fmt.Printf("The factorial of %d is %d\n", number, result)
}

func factorial(n int) int {
    if n == 0 {
        return 1
    }
    return n * factorial(n-1)
}
