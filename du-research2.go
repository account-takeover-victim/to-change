package main
import "github.com/ventu-io/go-shortid"
import "fmt"

func main() {
    // Get a greeting message and print it.
    fmt.Println("Your short ID is:")
    fmt.Println(shortid.Generate())
}

