package ascii

import (
	"fmt"
	"testing"
	"unicode"
)

const asciiConst = "\x00\x01\x02\x03\x04\x05\x06\x07\x08\x09\x0a\x0b\x0c\x0d\x0e\x0f" +
	"\x10\x11\x12\x13\x14\x15\x16\x17\x18\x19\x1a\x1b\x1c\x1d\x1e\x1f" +
	` !"#$%&'()*+,-./0123456789:;<=>?` +
	`@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\]^_` +
	"`abcdefghijklmnopqrstuvwxyz{|}~\x7f"

func TestAscii(t *testing.T) {

	//Jeg ser om det er noen char som er høyere enn 127. Hvis det er det,
	//inneholder ikke stringen kun ascii-tegn. Til dette bruker jeg unicode.MaxASCII.
	for i := 0; i < len(asciiConst); i++ {
		if asciiConst[i] > unicode.MaxASCII {
			t.Errorf("%q is not ASCII.", asciiConst[i])
		}
	}
	fmt.Println("")
}
