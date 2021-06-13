package main

// Integer Power ==================================================================================

func pow(a, n int64) int64 {
	var ret, MOD int64 = 1, 1000000007
	for n > 0 {
		if n&1 != 0 {
			ret = (ret * a) % MOD
		}
		a = (a * a) % MOD
		n /= 2
	}
	return ret
}
