package main
 
import (
    "fmt"
    "os"
	"bufio"
    "syscall"
)
 
var pipeFile = "pipe.log"
 
func main() {
    writeDataToPipe()
}
 
func writeDataToPipe() {
    os.Remove(pipeFile)
    err := syscall.Mkfifo(pipeFile, 0666)
	checkError(err)
    fmt.Println("Start to write data to Client using `pipe`")
    f, err := os.OpenFile(pipeFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	checkError(err)
	cs := make(chan string)
	go getInput(cs)
    for {
		number := <-cs
        f.WriteString(number)
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