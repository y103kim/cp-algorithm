package array

// Inversion Count with merge sort ================================================================

type InversionCount struct {
	arr   []int32
	bak   []int32
	count int
}

func (iv *InversionCount) init(a []int32) {
	iv.arr = a
	iv.bak = make([]int32, len(a))
	iv.count = 0
}

func (iv *InversionCount) mergeSort(l, r int) {
	if l >= r {
		return
	}
	m := (l + r) / 2
	iv.mergeSort(l, m)
	iv.mergeSort(m+1, r)

	lp, up, wp := l, m+1, l
	for wp <= r {
		if up == r+1 || (lp <= m && iv.arr[lp] <= iv.arr[up]) {
			iv.bak[wp] = iv.arr[lp]
			lp, wp = lp+1, wp+1
		} else {
			iv.bak[wp] = iv.arr[up]
			up, wp = up+1, wp+1
			iv.count += m + 1 - lp
		}
	}
	for i := l; i <= r; i++ {
		iv.arr[i] = iv.bak[i]
	}
}
