package ascii

import "fmt"

const ascii = "\x00\x01\x02\x03\x04\x05\x06\x07\x08\x09\x0a\x0b\x0c\x0d\x0e\x0f" +
	"\x10\x11\x12\x13\x14\x15\x16\x17\x18\x19\x1a\x1b\x1c\x1d\x1e\x1f" +
	` !"#$%&'()*+,-./0123456789:;<=>?` +
	`@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\]^_` +
	"`abcdefghijklmnopqrstuvwxyz{|}~\x7f"

func GetASCIIStringLiteral() string {
	return ascii
}

func IterateOverASCIIStringLiteral(stringLiteral string) {
	for i := 0; i < len(ascii); i++ {
		fmt.Printf("%X ", ascii[i])
		fmt.Printf("%+q ", ascii[i])
		fmt.Printf("%b ", ascii[i])
		fmt.Println(" ")

	}
}

func GreetingASCII() string {
	hex := "\x22\x48\x65\x6c\x6c\x6f\x20\x3a\x2d\x29\x22"
	return hex
}
