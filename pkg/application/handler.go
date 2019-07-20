package application

import (
	"fmt"
	"net/http"
	"time"
)

const fibNum = 10

// high computation
func fib1(n int) int {
	if n < 2 {
		return n
	}
	// some workload
	time.Sleep(100 * time.Millisecond)
	return fib1(n-1) + fib1(n-2)
}

// low computation
func fib2(n int) int {
	if n < 2 {
		return n
	}
	p2 := 1
	p1 := 1
	for i := 3; i < n; i++ {
		tmp := p1
		p1 = p2 + p1
		p2 = tmp
		// some workload
		time.Sleep(100 * time.Millisecond)
	}
	return p2 + p1
}

func fib1Handler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	res := fib1(fibNum)
	elapsed := time.Since(start)
	fmt.Fprintf(w, "fib1: n=%d, result=%d, time=%s", fibNum, res, elapsed)
}

func fib2Handler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	res := fib2(fibNum)
	elapsed := time.Since(start)
	fmt.Fprintf(w, "fib2: n=%d, result=%d, time=%s", fibNum, res, elapsed)
}

// Mux provides application route handling
func Mux() *http.ServeMux {
	m := http.NewServeMux()
	m.HandleFunc("/fib1", fib1Handler)
	m.HandleFunc("/fib2", fib2Handler)
	return m
}
