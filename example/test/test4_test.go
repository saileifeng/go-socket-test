package main

import (
	"sync"
	"testing"
)

type myint int64

type Inccer interface {
	inc()
}

func (i *myint) inc() {
	*i = *i + 1
}

func BenchmarkIntmethod(b *testing.B) {
	i := new(myint)
	incnIntmethod(i, b.N)
}

func BenchmarkInterface(b *testing.B) {
	i := new(myint)
	incnInterface(i, b.N)
}

func BenchmarkTypeSwitch(b *testing.B) {
	i := new(myint)
	incnSwitch(i, b.N)
}

func BenchmarkTypeAssertion(b *testing.B) {
	i := new(myint)
	incnAssertion(i, b.N)
}

func incnIntmethod(i *myint, n int) {
	for k := 0; k < n; k++ {
		i.inc()
	}
}

func incnInterface(any Inccer, n int) {
	for k := 0; k < n; k++ {
		any.inc()
	}
}

func incnSwitch(any Inccer, n int) {
	for k := 0; k < n; k++ {
		switch v := any.(type) {
		case *myint:
			v.inc()
		}
	}
}

func incnAssertion(any Inccer, n int) {
	for k := 0; k < n; k++ {
		if newint, ok := any.(*myint); ok {
			newint.inc()
		}
	}
}

func BenchmarkMux(b *testing.B) {
	mux := sync.Mutex{}
	flag := 0
	for i := 0; i < b.N; i++ {
		go func() {
			mux.Lock()
			flag ++
			mux.Unlock()
		}()
	}
}

//func BenchmarkAtom(b *testing.B) {
//	mux := sync.Mutex{}
//	flag := 0
//	for i := 0; i < b.N; i++ {
//		mux.Lock()
//		flag ++
//		mux.Unlock()
//	}
//}

func BenchmarkChan(b *testing.B) {
	ch := make(chan int,100000000)
	//flag := 0
	for i := 0; i < b.N; i++ {
		go func(v int) {
			ch <- v
		}(i)
	}
	//close(ch)
}
