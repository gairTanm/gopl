package funcs

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

var choice = flag.Int("c", 256, "use 384 or 512 for SHA384 or SHA512 respectively, defaults to SHA256")

//usage go run cryptCli.go -c <flag value> <some string>

func CliHash() {
	flag.Parse()
	b := *choice
	fmt.Println("Choice:", b)
	switch b {
	case 384:
		c := sha512.Sum384([]byte(os.Args[3]))
		fmt.Printf("SHA384: %x\n", c)
	case 512:
		c := sha512.Sum512([]byte(os.Args[3]))
		fmt.Printf("SHA512: %x\n", c)
	default:
		c := sha256.Sum256([]byte(os.Args[1]))
		fmt.Printf("SHA256: %x\n", c)
	}
}
