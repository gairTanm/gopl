package funcs

import (
	"bufio"
	"fmt"
	"os"
)

func WordFreq()map[string] int{
	freq:= make(map[string] int)
	f, err := os.Open("./data/words.txt");
	if err!=nil{
		fmt.Println("Error", err.Error())
		os.Exit(1)
	}
	input:= bufio.NewScanner(f)
	input.Split(bufio.ScanWords)
	for input.Scan(){
		word := input.Text()
		freq[word]++
	}
	if err:=input.Err(); err!=nil{
		fmt.Fprintf(os.Stderr, "error")
		os.Exit(1)
	}
	return freq
}
