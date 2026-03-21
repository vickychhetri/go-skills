Q1. What are goroutines
Ans. Goroutines are lightweight threads of execution in Go.
Goroutines are functions that can run concurrently with other functions.

They are cheap to create and can be used to handle multiple tasks simultaneously.

They communicate with each other using channels.

Goroutines are managed by the Go runtime and can be scheduled on multiple processors.

Example: go func() { fmt.Println("Hello, world!") }()

Example: go func() { result := doSomeWork(); channel <- result }()

Example: go func() { for i := 0; i < 10; i++ { fmt.Println(i) } }()




Asked in Seven Tech Solutions
6d ago

Q. How would you implement object-oriented programming concepts in Golang?
Ans. Implementing object-oriented programming concepts in Golang using structs and interfaces
Approach
Define structs to represent objects with fields as properties.
Use methods with value or pointer receivers on structs to define behavior (methods).
Implement interfaces by defining a set of method signatures and ensure structs define those methods.
Use interface variables to achieve polymorphism and dynamic dispatch.
Leverage embedding in structs to compose behaviors and support inheritance-like features.