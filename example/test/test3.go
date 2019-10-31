package main

import (
	"google.golang.org/grpc"
	"log"
)

func main() {
	//c := make(chan bool)
	//m := make(map[string]string)
	//go func() {
	//	m["1"] = "a" // First conflicting access.
	//	c <- true
	//}()
	//m["2"] = "b" // Second conflicting access.
	//<-c
	//for k, v := range m {
	//	fmt.Println(k, v)
	//}

	cc,err := grpc.Dial("192.168.1.1:8080",grpc.WithInsecure())
	if err == nil {
		log.Println(cc,err)
	}
}
