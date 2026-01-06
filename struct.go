package main

import "fmt"

// Definisi struct
type User struct {
    Name string
    Age  int
    Email string
}

func main() {
    // Buat instance
    user1 := User{
        Name: "Budi",
        Age: 25,
        Email: "budi@example.com",
    }
    
    // Akses field
    fmt.Println(user1.Name)
    
    // Update field
    user1.Age = 26
}
