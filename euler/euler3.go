//The prime factors of 13195 are 5, 7, 13 and 29.

//What is the largest prime factor of the number 600851475143 ?

package main

import (
  "fmt"
  "math"
)

// Only primes less than or equal to N will be generated
//const N = 100
//const inN = 100

func get_primes() [] int {
  var x, y, n int
  nsqrt := math.Sqrt(inN)

  is_prime := [inN]bool{}

  for x = 1; float64(x) <= nsqrt; x++ {
    for y = 1; float64(y) <= nsqrt; y++ {
      n = 4*(x*x) + y*y
      if n <= inN && (n%12 == 1 || n%12 == 5) {
        is_prime[n] = !is_prime[n]
      }
      n = 3*(x*x) + y*y
      if n <= inN && n%12 == 7 {
        is_prime[n] = !is_prime[n]
      }
      n = 3*(x*x) - y*y
      if x > y && n <= inN && n%12 == 11 {
        is_prime[n] = !is_prime[n]
      }
    }
  }

  for n = 5; float64(n) <= nsqrt; n++ {
    if is_prime[n] {
      for y = n * n; y < inN; y += n * n {
        is_prime[y] = false
      }
    }
  }

  is_prime[2] = true
  is_prime[3] = true

  primes := make([]int, 0, 1270606)
  for x = 0; x < len(is_prime)-1; x++ {
    if is_prime[x] {
      primes = append(primes, x)
    }
  }

  // primes is now a slice that contains all the
  // primes numbers up to inN

  // let's print them
  return primes
  //for _, x := range primes {
   //fmt.Println(x)
  //
}


const find = 600851475143
const inN = 1000000
//13195/2

func main() {
  //fmt.Println("get primes")
  primes := make([]int, 0, 1270606)
  primes = get_primes()
  //fmt.Println("done geting primes")
  for _, x := range primes {
    if find % x == 0 {
      fmt.Println(x)
    }
  }
}

