package main

import (
    "fmt"
    "os"
)

func main() {
    // Contoh 1: Buka file
    file, err := os.Open("data.txt")
    if err != nil {
        fmt.Println("Error cuy, ini penjelasanya:", err)
        return
    }
    defer file.Close() // defer = jalankan di akhir function
    
    fmt.Println("File berhasil dibuka")
}
