package euler

// ProductIntSlice スライスの各要素の積を計算
func ProductIntSlice(ns []int, convertor func(n int) int) int {
	product := 1
	for _, n := range ns {
		product *= convertor(n)
	}
	return product
}

// SumIntSlice スライスの各要素の和を計算
func SumIntSlice(ns []int, convertor func(n int) int) int {
	sum := 0
	for _, n := range ns {
		sum += convertor(n)
	}
	return sum
}
