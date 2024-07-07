package main

import "runtime"

func main() {
	// currentMaxProcs := runtime.GOMAXPROCS(0)
	// print(currentMaxProcs)
	runtime.GOMAXPROCS(2)
}
