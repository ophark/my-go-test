package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func get_url(url string) []uint8 {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "打开URL错误：%s\n", err)
		os.Exit(1)
	}
	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "获取body错误：%s\n", err)
		os.Exit(1)
	}
	//fmt.Printf("%T\n", body)
	return body
}
