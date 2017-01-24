package main

import (
	"fmt"
	"os"
)

func main() {
	//chengfa99_GO()
	fmt.Println("#########")
	//fmt.Println(equal(3, 3))
	//fmt.Printf("%c\n", test_if(50))
	//echo()
	//echo1()
	//echo2()
	/*********/
	for _, arg := range os.Args[1:] {
		//get_url(arg)
		fmt.Printf("%s\n", get_url(arg))
	}
}
