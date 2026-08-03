package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main()  {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter source file path:")
	scanner.Scan()
	src := scanner.Text()

	fmt.Println("Enter destination file path:")
	scanner.Scan()
	dst := scanner.Text()

	r, err := os.Open(src)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	
	defer r.Close()

	w, err := os.Create(dst)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	defer w.Close()
	
	_, err = io.Copy(w, r)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Copied successfully")
}