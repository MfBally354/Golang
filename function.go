package main

import "fmt"

// Function dengan parameter dan return
func tambah(a int, b int) int {
    return a + b
}

// Multiple return values (CIRI KHAS GO)
func bagi(a, b int) (int, error) {
    if b == 0 {
        return 0, fmt.Errorf("tidak bisa dibagi nol")
    }
    return a / b, nil
}

func main() {
    hasil := tambah(5, 3)
    fmt.Println("Hasil:", hasil)
    
    hasil2, err := bagi(10, 2)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Hasil bagi:", hasil2)
    }
}
