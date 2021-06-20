package algebra

type Sieve struct {
	isNotSieve []bool
}

func (sieve *Sieve) Init(n int) {
	sieve.isNotSieve = make([]bool, n+1)
	sieve.isNotSieve[1] = true
	for i := 2; i*i <= n; i++ {
		if sieve.isNotSieve[i] == false {
			for j := 2 * i; j <= n; j += i {
				sieve.isNotSieve[j] = true
			}
		}
	}
}
