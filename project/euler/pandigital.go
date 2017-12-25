package euler

// IsPandigital10 1-9までのパンデジタル数？
//
// https://ja.wikipedia.org/wiki/%E3%83%91%E3%83%B3%E3%83%87%E3%82%B8%E3%82%BF%E3%83%AB%E6%95%B0
func IsPandigital10(m int) bool {
	if m < 123456789 || 987654321 < m {
		return false
	}

	flags := uint16(0)

	for ; m > 0; m = m / 10 {
		f := uint16(1) << uint32(m%10-1)
		if flags&f > 0 {
			return false
		}
		flags |= f
	}

	return (flags == 512-1) // 2^9-1
}
