package main

import "fmt"

// Interface mendefinisikan behavior
type Speaker interface {
    Speak() string
}

// Struct 1
type Dog struct {
    Name string
}

func (d Dog) Speak() string {
    return "Woof!"
}

// Struct 2
type Cat struct {
    Name string
}

func (c Cat) Speak() string {
    return "Meow!"
}

// Function yang menerima interface
func makeSound(s Speaker) {
    fmt.Println(s.Speak())
}

func main() {
    dog := Dog{Name: "Buddy"}
    cat := Cat{Name: "Whiskers"}
    
    makeSound(dog) // Woof!
    makeSound(cat) // Meow!
}
