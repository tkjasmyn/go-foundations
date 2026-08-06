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

	stat, _ := os.Stat(r.Name())
	size := stat.Size()

	buf := make([]byte, 1024)
	var copied int64 = 0

	for {
		n, err := r.Read(buf)
		if n > 0 {
			w.Write(buf[:n])
			copied += int64(n)
			per := (copied * 100) / size
			fmt.Printf("Copying... %d%%\n", per)
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
	}

	fmt.Println("Copied successfully")
}