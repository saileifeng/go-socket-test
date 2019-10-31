package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"sync"
	"time"
)

func main() {

	go func() {
		http.ListenAndServe(":9999", nil)
	}()
	mux := sync.Mutex{}
	for j := 0; j < 10; j ++  {
		go func(num int) {
			for  {
				mux.Lock()
				//defer mux.Unlock()
				log.Printf("number %v get lock",num)
				if num != 9 {
					mux.Unlock()
				}
				time.Sleep(time.Millisecond*100)
			}
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