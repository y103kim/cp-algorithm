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

// MillerRabin ====================================================================================

func IsComposition(n, a, exp, cnt uint64) bool {
	x := Upow(a, exp, n)
	if x == 1 || x == n-1 {
		return false
	}
	for i := uint64(1); i < cnt; i++ {
		x = x * x % n
		if x == n-1 {
			return false
		}
	}
	return true
}

func MillerRabin(n uint64) bool {
	if n < 2 {
		return false
	}

	var exp, cnt uint64 = n - 1, 0
	for exp%2 == 0 {
		exp /= 2
		cnt++
	}

	candidates := []uint64{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37}
	for _, a := range candidates {
		if a == n {
			return true
		}
		if IsComposition(n, a, exp, cnt) {
			return false
		}
	}
	return true
}
