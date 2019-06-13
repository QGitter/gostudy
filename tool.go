//计算一个类型为float64的slice的平均值
func average(xs []float64) float64 {
	sum := 0.0
	if len(xs) == 0 {
		return 0
	}
	for _, v := range xs {
		sum = sum + v
	}
	return sum / float64(len(xs))
}

//字符串翻转算法
func reveser(s string) string {
	if len(s) == 0 {
		return ""
	}
	a := []rune(s)
	for i, j := 0, len(a)-1; i < j; i, j = j+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	return string(a)
}

//斐波那契数列
func fib() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}
//求sl切片最小值，最大值
func compare(sl []int) (ma, mi int) {
	ma, mi = sl[0], sl[0]
	for _, v := range sl {
		if v >= ma {
			ma = v
		}
		if v <= mi {
			mi = v
		}
	}
	return
}
//冒泡排序
func bubble(l []int) []int {
	if len(l) == 0 {
		return l
	}
	for i := 0; i < len(l); i++ {
		for j := i + 1; j < len(l); j++ {
			if l[i] > l[j] {
				l[i], l[j] = l[j], l[i]
			}
		}
	}
	return l
}

//整数转2进制
func convertToBin(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

//计算字符串中最大不重复的字串的长度
func substrNoRepeat(s string) int {
	lastOccurred := make(map[byte]int)
	start := 0
	maxLength := 0
	for i, ch := range []byte(s) {
		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return maxLength
}



