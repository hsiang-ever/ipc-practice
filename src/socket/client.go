package main

import (
    "fmt"
    "net"
    "os"
    "bufio"
    "strings"
    "strconv"
)

func main() {
    readDataFromSocket()
}

func readDataFromSocket() {
    service := "127.0.0.1:30001"
    tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
    checkError(err)
    conn, err := net.DialTCP("tcp", nil, tcpAddr)
    checkError(err)
    fmt.Println("Start to read data from server using `socket`")
    readFromConn(conn)
    os.Exit(0)
}

func readFromConn(conn net.Conn) {
    defer conn.Close()
    scanner := bufio.NewScanner(conn)
    for scanner.Scan() {
        ln := scanner.Text()
        fmt.Println("Data read from `socket`: " + ln)
        fmt.Println("Mean Value: ", getMean(ln))
    }
}

func getMean(serialNum string) float64 {
    nums := strings.Fields(serialNum)
    total := 0.0
    for _, textVal := range nums {
        v, err := strconv.ParseFloat(textVal, 64)
        checkError(err)
        total += v
    }
    return total/float64(len(nums))
}

func checkError(err error) {
    if err != nil {
        fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
        os.Exit(1)
    }
}