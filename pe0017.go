package project_euler

import (
	"sort"
	"strings"

	"github.com/pkg/errors"
)

var pe0017WordLength = map[int]int{
	1000: len(strings.Replace("one thousand", " ", "", -1)),
	900:  len(strings.Replace("nine hundred ", " ", "", -1)),
	800:  len(strings.Replace("eight hundred", " ", "", -1)),
	700:  len(strings.Replace("seven hundred", " ", "", -1)),
	600:  len(strings.Replace("six hundred", " ", "", -1)),
	500:  len(strings.Replace("five hundred", " ", "", -1)),
	400:  len(strings.Replace("four hundred", " ", "", -1)),
	300:  len(strings.Replace("three hundred", " ", "", -1)),
	200:  len(strings.Replace("two hundred", " ", "", -1)),
	100:  len(strings.Replace("one hundred", " ", "", -1)),
	90:   len("ninety"),
	80:   len("eighty"),
	70:   len("seventy"),
	60:   len("sixty"),
	50:   len("fifty"),
	40:   len("forty"),
	30:   len("thirty"),
	20:   len("twenty"),
	19:   len("nineteen"),
	18:   len("eighteen"),
	17:   len("seventeen"),
	16:   len("sixteen"),
	15:   len("fifteen"),
	14:   len("fourteen"),
	13:   len("thirteen"),
	12:   len("twelve"),
	11:   len("eleven"),
	10:   len("ten"),
	9:    len("nine"),
	8:    len("eight"),
	7:    len("seven"),
	6:    len("six"),
	5:    len("five"),
	4:    len("four"),
	3:    len("three"),
	2:    len("two"),
	1:    len("one"),
}

/*
	https://projecteuler.net/problem=16
*/
func PE0017(n int) (int, error) {
	if n < 1 || 1000 < n {
		return 0, errors.New("n must be between 1 and 1000")
	}

	// keys (降順)
	ks := []int{}
	for k, _ := range pe0017WordLength {
		ks = append(ks, k)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(ks)))

	sumWc := 0 // 総文字数
	for m := 1; m <= n; m++ {
		l := m
		for _, k := range ks {
			if l/k > 0 {
				sumWc += pe0017WordLength[k]
				if l = l % k; l > 0 && 100 <= k && k <= 900 {
					sumWc += 3 // "and" 分
				}
			}
		}
	}
	return sumWc, nil
}
