package main

import "github.com/ventu-io/go-shortid"
import "fmt"


func Shalom(name string) string {
    // Return a greeting that embeds the name in a message.
    message := fmt.Sprintf("Shalom, %v. Welcome!", name)
    return message
}


func main() {
    // Get a greeting message and print it.
    fmt.Println("Your short ID is:")
    fmt.Println(shortid.Generate())
}

