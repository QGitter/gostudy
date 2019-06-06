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



