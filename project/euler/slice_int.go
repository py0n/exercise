package euler

// ProductIntSlice スライスの各要素の積を計算
func ProductIntSlice(ns []int) int {
	if len(ns) == 0 {
		return 0
	}

	product := 1
	for _, n := range ns {
		product *= n
	}
	return product
}

// SumIntSlice スライスの各要素の和を計算
func SumIntSlice(ns []int) int {
	if len(ns) == 0 {
		return 0
	}

	sum := 0
	for _, n := range ns {
		sum += n
	}
	return sum
}
