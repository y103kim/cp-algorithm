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

// Matrix Power ==================================================================================

const MOD uint64 = 1000000007
const ORDER int = 8

type Matrix [ORDER][ORDER]uint64

func mul(a, b *Matrix) *Matrix {
	z := &Matrix{}
	for i := 0; i < ORDER; i++ {
		for j := 0; j < ORDER; j++ {
			for k := 0; k < ORDER; k++ {
				z[i][j] = (z[i][j] + a[i][k]*b[k][j]) % MOD
			}
		}
	}
	return z
}

func getIdentity() *Matrix {
	ret := &Matrix{}
	for i := 0; i < ORDER; i++ {
		ret[i][i] = 1
	}
	return ret
}

func matpow(a *Matrix, n uint64) *Matrix {
	ret := getIdentity()
	for n > 0 {
		if n&1 != 0 {
			ret = mul(ret, a)
		}
		a = mul(a, a)
		n /= 2
	}
	return ret
}
