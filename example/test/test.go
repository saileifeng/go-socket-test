package main

import (
	"log"
	"os"
	"os/signal"
	"sync"
)

func main() {


	mux := sync.Mutex{}
	for j := 0; j < 10; j ++  {
		go func(num int) {
			mux.Lock()
			log.Printf("number %v get lock",num)
			//mux.Unlock()
		}(j)
	}


	ShutDownHook()
}


func ShutDownHook() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Kill)
	a := <-quit
	log.Println("close ",a)
}