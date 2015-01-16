package main

import (
	. "fmt"
	"runtime"
	"time"
)

var i int = 0

func Inkrementer() {
	for x := 0; x<1000000; x++{
		i++	
	}
}

func Dekrementer() {
	for x := 0; x<1000000; x++{
		i--	
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) 
	go Inkrementer()
	go Dekrementer()
	time.Sleep(100*time.Millisecond)
	Println(i)
}
