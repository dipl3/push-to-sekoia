package main

import (
	"bufio"
	"bytes"
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

var in *bufio.Reader

func receivedLog(stdin io.Reader) (string, error) {
	if in == nil {
		in = bufio.NewReader(stdin)
	}
	logLine, err := in.ReadString('\n')
	return logLine, err
}

func main() {

	// Function to put the STDIN in a var called str
	loglign, err := receivedLog(os.Stdin)

	if err != nil {
		panic(err)
	}
	//fmt.Println(strings.TrimSuffix(str, "\n"))

	// Fuction to make a POST request
	url := "https://intake.sekoia.io/plain"

	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(loglign)))
	req.Header.Set("X-SEKOIAIO-INTAKE-KEY", os.Args[1])
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept-Encoding", "*")

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("panic")
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		fmt.Println(resp.StatusCode)
		os.Exit(resp.StatusCode)
	}
	body, _ := io.ReadAll(resp.Body)

	fmt.Printf("%s: %s\n", time.Now(), string(body))

	// TODO proper logging
	os.Exit(0)
}
