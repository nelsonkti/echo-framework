package helper

// Diff
// @Description: 差集
// @param a
// @param b
// @return []T
func Diff[T OrderedType](a, b []T) []T {
	if len(a) == 0 {
		return a
	}
	mb := make(map[T]struct{}, len(b))
	for _, x := range b {
		mb[x] = struct{}{}
	}
	var diff []T
	for _, x := range a {
		if _, found := mb[x]; !found {
			diff = append(diff, x)
		}
	}
	return diff
}

// Unique
// @Description: 数组去重
// @param a
// @return []T
func Unique[T OrderedType](a []T) []T {
	temp := map[T]struct{}{}
	result := make([]T, 0, len(a))
	for _, item := range a {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}

	return result
}

// Union
// @Description: 求并集
// @param slice1
// @param slice2
// @return []string
func Union[T OrderedType](slice1, slice2 []T) []T {
	m := make(map[T]int)
	for _, v := range slice1 {
		m[v]++
	}

	for _, v := range slice2 {
		times, _ := m[v]
		if times == 0 {
			slice1 = append(slice1, v)
		}
	}
	return slice1
}

// Intersect
// @Description: 求交集
// @param slice1
// @param slice2
// @return []string
func Intersect[T OrderedType](slice1, slice2 []T) []T {
	m := make(map[T]int)
	nn := make([]T, 0)
	for _, v := range slice1 {
		m[v]++
	}

	for _, v := range slice2 {
		times, _ := m[v]
		if times == 1 {
			nn = append(nn, v)
		}
	}
	return nn
}

// RemoveSlice
// @Description: 删除数组的某个元素
// @param a
// @param b
// @return []T
func RemoveSlice[T OrderedType](a []T, b T) []T {
	var newData []T
	for _, value := range a {
		if value != b {
			newData = append(newData, value)
		}
	}
	return newData
}
