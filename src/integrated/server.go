package main

import (
	"github.com/hidez8891/shm"
    "fmt"
    "net"
    "os"
	"bufio"
	"syscall"
	"time"
)

var cs chan string
var cs1 chan string
var cs2 chan string
var cs3 chan string
var pipeFile = "../pipe.log"

func main() {
	cs = make(chan string)
	cs1 = make(chan string)
	cs2 = make(chan string)
	cs3 = make(chan string)
	go writeDataToSocket()
	go writeDataToPipe()
	go writeDataToMemory()
	time.Sleep(1 * time.Second)
	go getInput()
	for {
		number := <-cs
		cs1 <- number
		cs2 <- number
		cs3 <- number
	}
}

func writeDataToMemory() {
	initMem()

	fmt.Println("Start to write data to Client using `shared memory`")
	w, _ := shm.Create("m1", 256)

	for {
		number := <-cs3
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

func writeDataToPipe() {
    os.Remove(pipeFile)
    err := syscall.Mkfifo(pipeFile, 0666)
	checkError(err)
    fmt.Println("Start to write data to Client using `pipe`")
    f, err := os.OpenFile(pipeFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	checkError(err)
    for {
		number := <-cs2
        f.WriteString(number)
    }
}

func writeDataToSocket() {
    service := "127.0.0.1:30001"
    tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
    checkError(err)
    fmt.Println("Start to write data to Client using `socket`")
    listener, err := net.ListenTCP("tcp", tcpAddr)
    checkError(err)
    for {
        conn, err := listener.Accept()
        if err != nil {
            continue
        }
        go handleClient(conn)
    }
}

func handleClient(conn net.Conn) {
    defer conn.Close()
	for {
		number := <-cs1
		conn.Write([]byte(number))
	}
}

func getInput() {
    reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter number: ")
		text, _ := reader.ReadString('\n')
		cs <- text
	}
}

func checkError(err error) {
    if err != nil {
        fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
        os.Exit(1)
    }
}