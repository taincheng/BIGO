package utils

func InterfaceToInt(i interface{}) int {
	switch i.(type) {
	case int:
		return i.(int)
	case int8:
		return int(i.(int8))
	case int16:
		return int(i.(int16))
	case int32:
		return int(i.(int32))
	case int64:
		return int(i.(int64))
	case uint:
		return int(i.(uint))
	case uint8:
		return int(i.(uint8))
	case uint16:
		return int(i.(uint16))
	case uint32:
		return int(i.(uint32))
	case uint64:
		return int(i.(uint64))
	default:
		return 0
	}
}

// Ptr 获取指针
func Ptr[T any](v T) *T {
	return &v
}

// ToPtrSlice 将值切片转换为指针切片
func ToPtrSlice[T any](slice []T) []*T {
	result := make([]*T, len(slice))
	for i := range slice {
		result[i] = &slice[i]
	}
	return result
}
