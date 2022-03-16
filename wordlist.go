package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type wordlist []string

//reading wordlist

func readwordlist(filename string) wordlist {
	//bs refer to byte and err refer to error
	bs, err := ioutil.ReadFile(filename)
	//nil means null or zero
	if err != nil {
		fmt.Print("error:", err)
		os.Exit(1)
	}
	//spliting new lines from file
	// s := strings.Split(string(bs), "/n")
	// fmt.Println(reflect.TypeOf(s))
	// fmt.Println(s[0])
	// return wordlist(s)

	//reading lines
	read_line := strings.Trim(string(bs), "\r\n")
	//reflecting data_type of readline = fmt.Println("var2 = ", reflect.TypeOf(read_line))
	// converting string as slice
	read_line_slice := strings.Fields(read_line)
	return wordlist(read_line_slice)
}

// //print string after reading file
// func (w wordlist) print() {
// 	for _, words := range w {
// 		fmt.Println("hello/" + words)
// 	}
// }
// another way of printing
func print(w wordlist) {
	for _, word := range w {
		a := "http://hello.com/" + word
		c := make(chan string)

		go request(a, c)
		// fmt.Println(a)
		// resp, err := http.Get(a)
		// if err != nil {
		// 	fmt.Print("error:", err)
		// }
		// fmt.Println("HTTP Response Status:", resp.StatusCode, http.StatusText(resp.StatusCode))
		// if resp.StatusCode >= 200 && resp.StatusCode <= 500 {
		// 	fmt.Println("HTTP Status is in the 2xx range")
		// } else {
		// 	fmt.Println("Argh! Broken")
		// }
		fmt.Println(<-c)
	}

}
func request(a string, c chan string) {
	resp, err := http.Get(a)
	if err != nil {
		fmt.Print("error:", err)

		c <- "error"
	}
	//fmt.Println("Url:", a, "\nHTTP Response Status:", resp.StatusCode, http.StatusText(resp.StatusCode))
	fmt.Println(a)

	c <- resp.Status
	// if resp.StatusCode >= 200 && resp.StatusCode <= 500 {
	// 	fmt.Println("HTTP Status is in the 2xx range")
	// 	c <- "working"
	// } else {
	// 	fmt.Println("Argh! Broken")
	// 	c <- "Broken"
	// 	return

	// }
	return
}
