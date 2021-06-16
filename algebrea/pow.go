package algebra

// Integer Power ==================================================================================

var MOD int64 = 1000000007

func Pow(a, n int64) int64 {
	var ret int64 = 1
	for n > 0 {
		if n&1 != 0 {
			ret = (ret * a) % MOD
		}
		a = (a * a) % MOD
		n /= 2
	}
	return ret
}

func Inverse(a int64) int64 {
	return Pow(a, MOD-2)
}

// Matrix Power ==================================================================================

const ORDER int = 8

type Matrix [ORDER][ORDER]int64

func Matmul(a, b *Matrix) *Matrix {
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

func GetIdentity() *Matrix {
	ret := &Matrix{}
	for i := 0; i < ORDER; i++ {
		ret[i][i] = 1
	}
	return ret
}

func Matpow(a *Matrix, n uint64) *Matrix {
	ret := GetIdentity()
	for n > 0 {
		if n&1 != 0 {
			ret = Matmul(ret, a)
		}
		a = Matmul(a, a)
		n /= 2
	}
	return ret
}
