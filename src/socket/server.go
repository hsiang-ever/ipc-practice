package main

import (
    "fmt"
    "net"
    "os"
	"bufio"
)

func main() {
    writeDataToSocket()
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
	cs := make(chan string)
	go getInput(cs)
	for {
		number := <-cs
		conn.Write([]byte(number))
	}
}

func getInput(c chan string) {
    reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter number: ")
		text, _ := reader.ReadString('\n')
		c <- text
	}
}

func checkError(err error) {
    if err != nil {
        fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
        os.Exit(1)
    }
}