package webuery

import (
	"strconv"
	"strings"
)

// StringToInt64s parse "1,2,3,4" to [1 2 3 4] and ignore ""
// so it can parse "1,2,3,,,4," to [1 2 3 4]
func StringToInt64s(str string, sep string) ([]int64, error) {
	splits := strings.Split(str, sep)

	splitsLen := len(splits)
	ints := []int64{}

	for i := 0; i < splitsLen; i++ {
		if len(splits[i]) > 0 {
			num, err := strconv.ParseInt(splits[i], 10, 64)
			if err != nil {
				return nil, err
			}
			ints = append(ints, num)
		}
	}

	return ints, nil
}
