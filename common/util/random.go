package util

import (
	"math/rand"
	"time"
)

const (
	RK_NUM   = 0 // 纯数字
	RK_LOWER = 1 // 小写字母
	RK_UPPER = 2 // 大写字母
	RK_ALL   = 3 // 数字、大小写字母
)

// Random 随机字符串
func Random(size int, kind int) string {
	ikind, kinds, result := kind, [][]int{[]int{10, 48}, []int{26, 97}, []int{26, 65}}, make([]byte, size)
	isAll := kind > 2 || kind < 0
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		if isAll { // random ikind
			ikind = rand.Intn(3)
		}
		scope, base := kinds[ikind][0], kinds[ikind][1]
		result[i] = uint8(base + rand.Intn(scope))
	}
	return string(result)
}
