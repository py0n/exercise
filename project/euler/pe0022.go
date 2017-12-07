package euler

import (
	"encoding/csv"
	"io"
	"os"
	"sort"
	"strings"
)

// PE0022 リストに在る単語の点数を計算し総和を求める
func PE0022(fn string) (int, error) {
	fh, err := os.Open(fn)
	if err != nil {
		return 0, err
	}
	defer fh.Close()

	reader := csv.NewReader(fh)
	rows, err := reader.ReadAll()
	if err != io.EOF && err != nil {
		return 0, err
	}

	names := []string{}
	for _, row := range rows {
		names = append(names, row...)
	}
	sort.SliceStable(names, func(i, j int) bool {
		return strings.Compare(names[i], names[j]) < 1
	})

	sum := 0
	for i, name := range names {
		sum += (i + 1) * wordWorth(name)
	}

	return sum, nil
}

func wordWorth(word string) int {
	worth := 0
	for _, r := range word {
		worth += int(r-'A') + 1
	}
	return worth
}
