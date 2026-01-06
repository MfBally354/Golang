package main

import "fmt"

func main() {
    x := 10
    
    // IF statement
    if x > 5 {
        fmt.Println("Besar")
    } else {
        fmt.Println("Kecil")
    }
    
    // FOR loop (Go tidak punya while!)
    for i := 0; i < 5; i++ {
        fmt.Println(i)
    }
    
    // FOR sebagai while
    count := 0
    for count < 3 {
        fmt.Println("Count:", count)
        count++
    }
    
    // Infinite loop
    // for {
    //     fmt.Println("Forever")
    //     break // gunakan break untuk keluar
    // }
}
