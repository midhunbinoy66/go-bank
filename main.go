package main

import "fmt"

func main(){
	server := NewApiServer(":3000")
	server.Run()
	fmt.Print("server is created")

}
