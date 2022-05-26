package main

// Fibonacci sequence terms
func fibonnacci(n int) {
	if n <=1 {
	return n
	}

	return fibonacci(n-1) + fibonacci(n-2)
}

