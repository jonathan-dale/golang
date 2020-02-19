package main

import "fmt"



// the func intSeq returns another function 
// which we define anonymously in the body
// the returned func closes over the var i to form closure.
func intSeq() func() int {
    i := 0
    return func() int {
        i++
        return i
    }

}

func main() {


    // we call intSeq and assign the return value (a function) 
    // to nextInt. Its value i will be updated each time we call nextInt like with print statement.
    nextInt := intSeq()


    // see the effects by calling nextInt a few times
    fmt.Println(nextInt())
    fmt.Println(nextInt())
    fmt.Println(nextInt())
    fmt.Println(nextInt())
    fmt.Println(nextInt())


    // The state is unique to a particular finction,
    // we call intSeq and set its value to a new variable here
    // to confirm it is reset.
    newInts := intSeq()
    fmt.Println("newInts value is now ", newInts())
    fmt.Println("newInts value is now ", newInts())

}
