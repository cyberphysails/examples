package main

import (
    "fmt"
    "time"
    "net"
    "strings"
)


func createLog(msg, quit chan string) {
    lastlog, newlog := "starting", ""
    pattern, index := []string{"-", "\\", "|", "/"}, 0 
    for {
        select {
        case newlog = <- msg:
            fmt.Print("\033[2K\r+ ", lastlog)
            fmt.Print("\n")
            lastlog = strings.Trim(newlog, "\n")
        case <- quit:
            fmt.Print("\033[2K\r+ ", lastlog)
            fmt.Print("\n+ end\n")
            return
        default:
            fmt.Printf("\033[2K\r%v %v", pattern[index], lastlog)
            index += 1
            if index >= len(pattern) {
                index = 0
            }
            time.Sleep(500 * time.Millisecond)
        }
    }
}



func load() {
    // clear screen
    //fmt.Print("\033c")
    // print loading...
    var pattern = [3]string{".", "..", "..."}
    //var pattern = [4]string{"-", "\\", "|", "/"}
    for {
        for _, val := range pattern {
            // \033[2K clear screen from current cursor to the beginning of line
            // \r go to the beginning of line 
            fmt.Print("\033[2K\rloading", val)
            time.Sleep(1 * time.Second)
        }
    }
}
func listenS(output chan string) {
    listener, err := net.Listen("unix", "slog.socket")
    //listener, err := net.Listen("tcp4", "0.0.0.0:8080")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer listener.Close()

    for {
        conn, err := listener.Accept()
        if err != nil {
            fmt.Println(err)
            return
        }
        msg := make([]byte, 255)
        n, err := conn.Read(msg)
        if err != nil {
            fmt.Println(err)
            return
        }
        conn.Close()

        output <- string(msg[:n])
    }
}
func main() {
    //load()
    msg := make(chan string)
    quit := make(chan string)
    go createLog(msg, quit)

    //listenS(msg) 
    msg <- "task1"
    time.Sleep(5 * time.Second)
    msg <- "task2"
    time.Sleep(5 * time.Second)
    close(quit)
    //quit <- "end"
    
}
