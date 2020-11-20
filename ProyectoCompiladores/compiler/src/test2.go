package main

import (
	"fmt"
)
func main() {
	M["Bell Labs"] =  "ooooo"
	if M["Bell Labs"] == ""{
		fmt.Println("NO ESTA")	
	}else{
		fmt.Println(M["Bell Labs"])
	}
}
