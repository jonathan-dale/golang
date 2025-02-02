package main

import "fmt"

func sum(nums ...int) {

    fmt.Print(nums, " ")
    total := 0
    for _, num := range nums {
        total += num
    }
    fmt.Println(total)

}


func main() {

    sum(1, 2)
    sum(1, 2, 3)
    sum(1, 2, 3, 4)

    // if we have mult. args in a slice
    // apply them to variadic function using func(slice...) like this
    nums := []int{1, 2, 3, 4, 5, 6}
    sum(nums...)

}
