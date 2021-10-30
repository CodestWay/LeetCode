package middle

import (
	"sort"
	"strconv"
)

var dict map[string]bool

func init() {
	dict = make(map[string]bool)
	for i := 0; i < 30; i++ {
		dict[toString(1<<i)] = true
	}
}

func toString(n int) string {
	ba := make([]byte, 0)
	for ; n != 0; n /= 10 {
		ba = append(ba, strconv.Itoa(n % 10)[0])
	}
	sort.Slice(
		ba,
		func(i, j int) bool {
			return ba[i] < ba[j]
		},
	)
	return string(ba)
}

func reorderedPowerOf2(n int) bool {
	return dict[toString(n)]
}
