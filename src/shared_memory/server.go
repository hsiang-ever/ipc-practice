package main

import (
    "github.com/hidez8891/shm"
    "fmt"
    "os"
	"bufio"
	"time"
)

func main() {
	writeDataToMemory()
}

func writeDataToMemory() {
	initMem()

	fmt.Println("Start to write data to Client using `shared memory`")
	w, _ := shm.Create("m1", 256)
    cs := make(chan string)
	go getInput(cs)

	for {
		number := <-cs
		w.Write([]byte(number))
		time.Sleep(1 * time.Second)
		w.Close()
		w, _ = shm.Create("m1", 256)
	}
}

func initMem() {
    w, _ := shm.Create("m1", 256)
    w.Close()
}

func getInput(c chan string) {
    reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter number: ")
		text, _ := reader.ReadString('\n')
		c <- text
	}
}