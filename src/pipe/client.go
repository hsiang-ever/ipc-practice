package main
 
import (
    "bufio"
    "fmt"
    "os"
	"strings"
	"strconv"
	"github.com/montanaflynn/stats"
)
 
var pipeFile = "pipe.log"
 
func main() {
    readDataFromPipe()
}

func readDataFromPipe() {
    fmt.Println("Start to read data from Server using `pipe`")
    file, err := os.OpenFile(pipeFile, os.O_CREATE, os.ModeNamedPipe)
    checkError(err)
 
    reader := bufio.NewReader(file)
 
    for {
        line, err := reader.ReadBytes('\n')
        if err == nil {
            fmt.Print("Data read from `pipe`: " + string(line))
			fmt.Println("Median Value: ", getMedian(string(line)))
        }
    }
}

func getMedian(serialNum string) float64 {
    nums := strings.Fields(serialNum)
	var data []float64
	for _, v := range nums {
		v, err := strconv.ParseFloat(v, 64)
		checkError(err)
        data = append(data, v)
	}
    median, _ := stats.Median(data)
    return median
}

func checkError(err error) {
    if err != nil {
        fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
        os.Exit(1)
    }
}