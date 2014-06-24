package math

// Taken from go Playground (play.golang.org/p/9U22NfrXeq)
// Send the sequence 2, 3, 4, ... to channel 'ch'.
func generate(ch chan<- int) {
	for i := 2; ; i++ {
		ch <- i // Send 'i' to channel 'ch'.
	}
}

// Copy the values from channel 'in' to channel 'out',
// removing those divisible by 'prime'.
func filterPrimes(in <-chan int, out chan<- int, prime int) {
	for {
		i := <-in // Receive value from 'in'.
		if i%prime != 0 {
			out <- i // Send 'i' to 'out'.
		}
	}
}

// The prime sieve: Daisy-chain Filter processes.
func NPrimes(n int) []int {
	s := make([]int, n)
	ch := make(chan int) // Create a new channel.
	go generate(ch)      // Launch Generate goroutine.
	for i := 0; i < n; i++ {
		prime := <-ch
		s[i] = prime
		ch1 := make(chan int)
		go filterPrimes(ch, ch1, prime)
		ch = ch1
	}

	return s
}
