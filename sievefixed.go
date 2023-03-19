package main

import "fmt"

func main() {
	fmt.Println(SieveOfEratosthenes(101))
	fmt.Println(TwinPrimes(101))
}

// working off of wikipedia's definition of the sieve https://en.wikipedia.org/wiki/Sieve_of_Eratosthenes
func SieveOfEratosthenes(n uint) []uint {
	n++ // assuming the input wants to be inclusive for finding primes, so its easiest to increment it by 1
	isPrime := make([]bool, n)

	for i := 0; uint(i) < n; i++ {
		isPrime[i] = true // it makes the most sense to me to start with them as true for the algorithm to work
	}

	isPrime[0] = false //  I dont think 0 or 1 are prime?
	if n > 1 {         // in case of 0 input
		isPrime[1] = false
	}

	//fmt.Println("primes created")
	//fmt.Println(isPrime)

	for i := 2; uint(i) < n; i++ { //seemed easier to just start at 2, also allows for the optimization of only going to the sqrt(n)
		//fmt.Println(isPrime[i])
		if isPrime[i] == true { //if we've gotten to a value that is true, we know it is prime because no previous value is a factor of it, furthermore we can then elimante its multiples
			for k := uint(2); uint(i)*k < n; k++ { //this could also start at i^2 and then iterate i^2+i, i^2+2i, but I left it as is.
				isPrime[uint(i)*k] = false
			}
		}
	}

	//fmt.Println("primes found")
	//fmt.Println(isPrime)

	var primes []uint
	for p, _ := range isPrime {
		if isPrime[p] == true {
			primes = append(primes, uint(p))
		}
	}

	return primes
}

func TwinPrimes(n uint) [][]uint {
	primes := SieveOfEratosthenes(n)     // first finds the primes up to the given number
	var TwinPrimes [][]uint              // will store the primes in pairs
	for i := 0; i < len(primes)-1; i++ { // for each prime check if the next prime is a twin, if so add both to the list
		if primes[i]+2 == primes[i+1] {
			var temp []uint
			temp = append(temp, uint(primes[i]), uint(primes[i+1]))
			TwinPrimes = append(TwinPrimes, temp)

		}
	}
	return TwinPrimes
}
