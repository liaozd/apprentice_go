package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	for idx, args := range os.Args {
		fmt.Println("Args: "+strconv.Itoa(idx)+":", args)
	}
}
