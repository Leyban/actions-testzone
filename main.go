// Golang program to illustrate the usage of
// Time.Truncate() function

// Including main package
package main

// Importing fmt and time
import (
	"fmt"
	"time"
)

// Calling main
func main() {

	// Defining t for Truncate method
	t := time.Now().UTC()

	// Defining duration
	d := (60 * time.Minute)

	// Calling Truncate() method
	trunc := t.Truncate(d)

	// Prints output
	fmt.Println(t)

	fmt.Println(trunc)
}
