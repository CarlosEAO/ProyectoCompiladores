package main

import "fmt"

var m map[string]string

func main() {
	m = make(map[string]string)
	m["Bell Labs"] =  "ooooo"
	if m["berll"] == ""{
		fmt.Println("NO ESTA")	
	}else{
		fmt.Println(m["Bell"])
	}
}
