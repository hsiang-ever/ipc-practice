package main

import (
    "github.com/hidez8891/shm"
    "fmt"
    "bytes"
    "os"
    "strings"
	"strconv"
    "github.com/montanaflynn/stats"
    "time"
)

func main() {
    readDataFromMemory()
}

func readDataFromMemory() {
    rbuf := make([]byte, 256)
    empty := make([]byte, 256)
    fmt.Println("Start to read data from server using `shared memory`")
    for {
        r, _ := shm.Open("m1", 256)
        i, _ := r.Read(rbuf)
        if bytes.Equal(rbuf, empty)==false {
            result := string(rbuf[:i])
            fmt.Print("Data read from `shared memory`: " + result)
            fmt.Println("Mode Value: ", getMode(result))
            time.Sleep(2 * time.Second)
        }
        r.Close()
    }
}

func getMode(serialNum string) []float64 {
    serialNum = strings.TrimSuffix(serialNum, "\n")
    nums := strings.Fields(serialNum)
	var data []float64
	for _, v := range nums {
		v, err := strconv.ParseFloat(v, 64)
		if err != nil {
            continue
        }
        data = append(data, v)
	}
    mode, _ := stats.Mode(data)
    return mode
}

func checkError(err error) {
    if err != nil {
        fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
        os.Exit(1)
    }
}