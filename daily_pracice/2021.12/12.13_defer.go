package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	contents, err := Contents("/tmp/glances-root.log")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(contents)
}

// Contents returns the file's contents as a string
func Contents(filename string) (string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer f.Close() // f.Close will run when we're finished.

	var result []byte
	buf := make([]byte, 100)
	for {
		n, err := f.Read(buf[0:])
		result = append(result, buf[0:n]...)
		if err != nil {
			if err == io.EOF {
				break
			}
			return "", err // f will be closed if we return here.
		}
	}
	return string(result), nil // f will be closed if we return here.
}
