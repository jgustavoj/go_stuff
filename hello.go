package main

import (
	"fmt"

	greetings "github.com/jgustavoj/go_greetings"
	// "rsc.io/quote"
)

// func main() {
//     fmt.Println("Hello, World!")
// }

// func main() {
// 	fmt.Println(quote.Go())
// }

func main() {
	// Get a greeting message and print it.
	message := greetings.Hello("Gladys")
	fmt.Println(message)
}
