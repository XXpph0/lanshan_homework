package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	f, err := os.OpenFile("test", os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	start := time.Now()
	bufWriter := bufio.NewWriter(f)
	for i := 0; i < 9999; i++ {
		bufWriter.Write([]byte("hello"))
	}
	bufWriter.Flush()
	cost := time.Since(start)
	fmt.Printf("有缓冲：%v ns\n", cost)

	start = time.Now()
	for i := 0; i < 9999; i++ {
		f.Write([]byte("hello"))
	}
	cost = time.Since(start)
	fmt.Printf("无缓冲：%v ns\n", cost)
}
