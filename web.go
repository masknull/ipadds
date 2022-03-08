package main

import (
	"fmt"
	ip2region "ipadds"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	if len(os.Args) == 3 {
		http.HandleFunc("/", handler)
		log.Fatal(http.ListenAndServe("0.0.0.0:"+os.Args[2], nil))
	} else if len(os.Args) == 2 {
		http.HandleFunc("/", handler)
		log.Fatal(http.ListenAndServe("0.0.0.0:80", nil))
	} else {
		fmt.Println(fmt.Sprintf("\x1b[0;31m%s\x1b[0m", "格式错误：./ipaddweb db文件 端口\n如：./ipaddweb db 80"))
	}
}

func handler(writer http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" && strings.Contains(request.Header.Values("Accept")[0], "html") && request.URL.RawQuery != "" {
		var StrUrl []string
		var data string

		Str := strings.Split(request.URL.RawQuery, "&")
		if len(Str) != 0 {
			StrUrl = strings.Split(Str[0], "=")
		} else {
			StrUrl = strings.Split(request.URL.RawQuery, "=")
		}
		Ipi := StrUrl[1]
		Ipis := strings.Split(Ipi, ".")
		if len(Ipis) != 4 || Ipis[0] == "" || Ipis[1] == "" || Ipis[2] == "" || Ipis[3] == "" {
			data = "请输入正确的ip"
		} else {
			FileName := os.Args[1]
			region, _ := ip2region.New(FileName)
			ip, err := region.MemorySearch(Ipi)
			data = Ipi + " " + ip.String()
			if err != nil {
				data = "获取失败"
			}
		}
		t := time.Now()
		fmt.Printf("%d-%d-%d %d:%d:%d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
		print(" - ip:" + Ipi + "\n")
		fmt.Fprint(writer, data)
	} else if request.URL.RawQuery == "" {
		data := "Hello World"
		fmt.Fprint(writer, data)
	} else if request.Method != "GET" {
		data := "GET"
		fmt.Fprint(writer, data)
	}
}
