package main

/***
Golang doesn’t have classical OOP like inheritance/classes (as in Java),
but you can implement all core OOP concepts using structs, methods, interfaces, and embedding.

Encapsulation
Inheritance (via embedding)
Polymorphism (via interfaces)

1. Define Base Struct (Encapsulation)
package main

import "fmt"

// Base struct
type Animal struct {
    Name string
}

// Method (behavior)
func (a Animal) Speak() string {
    return "Some generic sound"
}




2. Inheritance using Embedding

type Dog struct {
    Animal // embedding (like inheritance)
}

func (d Dog) Speak() string {
    return "Bark"
}

type Cat struct {
    Animal
}

func (c Cat) Speak() string {
    return "Meow"
}




3. Polymorphism using Interface
// Interface
type Speaker interface {
    Speak() string
}


4. Use Polymorphism (Dynamic Dispatch)
func printSound(s Speaker) {
    fmt.Println(s.Speak())
}


5. Main Function

func main() {
    dog := Dog{Animal{Name: "Buddy"}}
    cat := Cat{Animal{Name: "Whiskers"}}

    printSound(dog) // Bark
    printSound(cat) // Meow

    // Interface slice (polymorphism)
    animals := []Speaker{dog, cat}

    for _, a := range animals {
        fmt.Println(a.Speak())
    }
}
**/

// Approach
// Define structs to represent objects with fields as properties.
// Use methods with value or pointer receivers on structs to define behavior (methods).
// Implement interfaces by defining a set of method signatures and ensure structs define those methods.
// Use interface variables to achieve polymorphism and dynamic dispatch.
// Leverage embedding in structs to compose behaviors and support inheritance-like features.

// ****************************************************************************************************

// Q. How would you implement a RESTful API using Golang?
// Ans. Implement a RESTful API using Golang with routing and handlers for CRUD operations
// Approach
// Initialize a new Go module and import necessary packages like net/http and gorilla/mux for routing
// Define data structures to represent resources (e.g., structs for models)
// Create handler functions for each RESTful endpoint (GET, POST, PUT, DELETE) to manage resource operations
// Use a router (e.g., gorilla/mux) to map endpoints to handler functions
// Start the HTTP server listening on a specific port

// ***********************************************************************************************************



