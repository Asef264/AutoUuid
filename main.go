package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/atotto/clipboard"
	"github.com/google/uuid"
)

func main() {
	for {
		id := uuid.New().String()
		uuidString := fmt.Sprintln(id)
		err := clipboard.WriteAll(uuidString)
		if err != nil {
			panic("error on coping ")
		}
		fmt.Println("copied :) you can paste it now")
		bufio.NewReader(os.Stdin).ReadByte()
	}

}
