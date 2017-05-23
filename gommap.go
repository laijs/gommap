package main

import (
	"fmt"
	"os"

	"golang.org/x/exp/mmap"
)

func main() {
	if len(os.Args) < 2 {
		return
	}

	m, err := mmap.Open(os.Args[1])
	if err != nil {
		fmt.Printf("error :%v", err)
	}
	for {
		for i := 0; i < m.Len(); i = i + 4096 {
			m.At(i)
		}
	}
}
