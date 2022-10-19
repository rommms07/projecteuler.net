package p1

var mul3or5_sum_s1_mem = make(map[int64]int64)

func mul3or5_sum_s1(n int64) (res int64) {
	i, i2 := 3, 5

	added := make(map[int64]bool)

	sumf := func(N ...int) {
		for _, z := range N {
			for j := int64(z); j < n; j += int64(z) {
				if added[j] {
					continue
				}

				res += j
				added[j] = true
			}
		}
	}

	if mul3or5_sum_s1_mem[n] != 0 {
		res = mul3or5_sum_s1_mem[n]
	} else {
		sumf(i, i2)
		mul3or5_sum_s1_mem[n] = res
	}

	return
}

func mul3or5_sum_s2(n int64) (res int64) {
	for j := int64(1); j < n; j++ {
		if j%3 == 0 || j%5 == 0 {
			res += j
		}
	}

	return
}
