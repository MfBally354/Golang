package main

import "fmt"

func main() {
    // Buat slice
    slice := []int{1, 2, 3}
    
    // Append
    slice = append(slice, 4, 5)
    
    // Akses elemen
    fmt.Println(slice[0]) // 1
    
    // Length
    fmt.Println(len(slice)) // 5
    
    // Loop slice
    for i, val := range slice {
        fmt.Printf("Index %d: %d\n", i, val)
    }
}
