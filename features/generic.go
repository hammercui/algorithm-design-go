package features

/**
找到值对应的下标,T约束为可比较类型
*/
func findIndex[T comparable](array []T, x T) int {
	for i, val := range array {
		if val == x {
			return i
		}
	}
	return - 1
}

/**
自定义约束
 */
type ConcatConstraint interface {
	int | int32 | int64 | float32
}

func concatenate[T ConcatConstraint](a []T,b []T) []T{
	// 创建一个长度为a+b的切片
	result := make([]T, len(a)+len(b))

	// 复制a和b的元素到新的切片
	copy(result,a)
	copy(result[len(a):],b)
	return result
}

