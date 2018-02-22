package main

import (
	"net/http"
	"runtime"
	"fmt"
	"time"
)

var reqNum int64

func main() {
	runtime.GOMAXPROCS(4)

	for i := 0; i < 10000; i++ {
		//go get("http://www.xiaoxiao.work")
		go get("http://letsgo.tech")
		//go get("http://blog.xiaoxiao.work")
	}
	time.Sleep(time.Minute * 10)
}

func get(url string) {
	for {
		_, err := http.Get(url)

		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("get over")
		//defer resp.Body.Close()
		//body, err := ioutil.ReadAll(resp.Body)
		//if err != nil {
		//	// handle error
		//	fmt.Println(err)
		//}
		//fmt.Println(string(body))
	}
}
