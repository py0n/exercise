package euler

// IsPandigital10 1-9までのパンデジタル数？
//
// https://ja.wikipedia.org/wiki/%E3%83%91%E3%83%B3%E3%83%87%E3%82%B8%E3%82%BF%E3%83%AB%E6%95%B0
func IsPandigital10(m int) bool {
	if m < 123456789 {
		return false
	}
	if m > 987654321 {
		return false
	}

	flags := int16(0)

	for m > 0 {
		r := m % 10
		flags = flags | (1 << uint32(r-1))
		m = m / 10
	}

	return (flags == 511)
}
