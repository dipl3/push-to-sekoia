package main

import (
        "bytes"
        "fmt"
        "io"
        "net/http"
        "os"
        "strings"
)

func main() {
        // Function to put the STDIN in a var called str
        stdin, err := io.ReadAll(os.Stdin)

        if err != nil {
                panic(err) 
        }
        str := string(stdin) 
        fmt.Println(strings.TrimSuffix(str, "\n")) 
        
        // Fuction to make a POST request 
        url := "https://intake.sekoia.io/plain" 

        var data = []byte(str) 
        req, err := http.NewRequest("POST", url, bytes.NewBuffer(data)) 
        req.Header.Set("X-SEKOIAIO-INTAKE-KEY", os.Args[1]) 
        req.Header.Set("Content-Type", "application/json") 

        client := &http.Client{} 
        resp, err := client.Do(req)
        if err != nil {
                panic(err) 
        }
        defer resp.Body.Close()

        fmt.Println("response Status:", resp.Status) 
        fmt.Println("response Headers:", resp.Header)
        body, _ := io.ReadAll(resp.Body)
        fmt.Println("response Body:", string(body))
}
