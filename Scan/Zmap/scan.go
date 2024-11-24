package main

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"sync"
	"time"
)

var syncWait sync.WaitGroup

var statok, fall, all int
var port int
var filename string

func checkProxy(ip string) {
	defer syncWait.Done()

	proxy := fmt.Sprintf("http://%s:%d", ip, port)

	parsedProxy, err := url.Parse(proxy)
	if err != nil {
		fmt.Println("Lỗi", err)
		return
	}

	transport := &http.Transport{
		Proxy: http.ProxyURL(parsedProxy),

		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{
		Transport: transport,
		Timeout:   5 * time.Second,
	}

	resp, err := client.Get("https://httpbin.org/get")
	if err != nil {
		all++
		fall++
		return
	}
	defer resp.Body.Close() // Đảm bảo đóng body

	if resp.StatusCode == http.StatusOK {
		statok++
		luufile(ip)
	} else {
		//	all++
		fall++
	}
	all++
}
func luufile(ip string) {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Lỗi khi mở tệp:", err)
		return
	}
	defer file.Close()

	content := ip + "\n"

	_, err = file.WriteString(content)
	if err != nil {
		fmt.Println("Lỗi khi ghi vào tệp:", err)
		return
	}

}

func main() {

	const (
		Reset   = "\033[0m"
		Red     = "\033[31m"
		Green   = "\033[32m"
		Yellow  = "\033[33m"
		Blue    = "\033[34m"
		Magenta = "\033[35m"
		Cyan    = "\033[36m"
		White   = "\033[37m"
	)

	if len(os.Args) >= 2 {
		portArg := os.Args[1]
		var err error
		port, err = strconv.Atoi(portArg)
		if err != nil {

			port = 8080
		}
	} else {
		port = 8080
	}
	filename = fmt.Sprintf("proxy_%d.txt", port)
	var timest int

	go func() {
		for {
			fmt.Printf(Yellow+" SCAN PROXY : OKE : %d  | "+Red+"FALSE : %d | "+Cyan+"ALL : %d | Time : %d"+Reset+"\n", statok, fall, all, timest)
			time.Sleep(1 * time.Second)
			timest++

		}
	}()

	for {
		r := bufio.NewReader(os.Stdin)
		scan := bufio.NewScanner(r)
		for scan.Scan() {
			go checkProxy(scan.Text())

			syncWait.Add(1)
		}
	}

}
