package funcs

import "fmt"

func AdjSpaces() {
	const nihongo = "日本語"
	for idx, runeValue := range nihongo {
		fmt.Printf("%#U at %d", runeValue, idx)
	}
}
