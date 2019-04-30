package main

import (
	"fmt"

	ascii "github.com/Brofo/ica02/ascii/asciiPack"
)

func main() {

	ascii.IterateOverASCIIStringLiteral(ascii.GetASCIIStringLiteral())

	fmt.Println(ascii.GreetingASCII())

}
