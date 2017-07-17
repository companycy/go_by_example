package main

import (
	"fmt"
)

// 1.9
// rsa
func rsa() {
	// pick 2 large n-bit primes p/q
	// public key is (N = p*q, e) where gcd(e, (p-1)*(q-1)) = 1)
	// private key is d, d = e^(-1) % (p-1)*(q-1)

	// recv msg: y = x^e % N
	// x = y ^ d % N
}

func main() {
	
}
