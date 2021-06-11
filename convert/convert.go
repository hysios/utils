package convert

import (
	"fmt"
	"strconv"
	"time"
)

// Int 转成整型
func Int(v interface{}) (int, bool) {
	switch i := v.(type) {
	case int:
		return i, true
	case uint:
		return int(i), true
	case int8:
		return int(i), true
	case int16:
		return int(i), true
	case int32:
		return int(i), true
	case int64:
		return int(i), true
	case uint8:
		return int(i), true
	case uint16:
		return int(i), true
	case uint32:
		return int(i), true
	case uint64:
		return int(i), true
	case float32:
		return int(i), true
	case float64:
		return int(i), true
	case bool:
		if i {
			return 1, true
		} else {
			return 0, true
		}
	case string:
		f, err := strconv.ParseFloat(i, 64)
		if err != nil {
			return 0, false
		}
		return int(f), true
	default:
		return 0, false
	}
}

func Int32(v interface{}) (int32, bool) {
	switch i := v.(type) {
	case int:
		return int32(i), true
	case uint:
		return int32(i), true
	case int8:
		return int32(i), true
	case int16:
		return int32(i), true
	case int32:
		return int32(i), true
	case int64:
		return int32(i), true
	case uint8:
		return int32(i), true
	case uint16:
		return int32(i), true
	case uint32:
		return int32(i), true
	case uint64:
		return int32(i), true
	case float32:
		return int32(i), true
	case float64:
		return int32(i), true
	case bool:
		if i {
			return 1, true
		} else {
			return 0, true
		}
	case string:
		f, err := strconv.ParseFloat(i, 64)
		if err != nil {
			return 0, false
		}
		return int32(f), true
	default:
		return 0, false
	}
}

func Int64(v interface{}) (int64, bool) {
	switch i := v.(type) {
	case int:
		return int64(i), true
	case uint:
		return int64(i), true
	case int8:
		return int64(i), true
	case int16:
		return int64(i), true
	case int32:
		return int64(i), true
	case int64:
		return int64(i), true
	case uint8:
		return int64(i), true
	case uint16:
		return int64(i), true
	case uint32:
		return int64(i), true
	case uint64:
		return int64(i), true
	case float32:
		return int64(i), true
	case float64:
		return int64(i), true
	case bool:
		if i {
			return 1, true
		} else {
			return 0, true
		}
	case string:
		f, err := strconv.ParseFloat(i, 64)
		if err != nil {
			return 0, false
		}
		return int64(f), true
	default:
		return 0, false
	}
}

func Uint(v interface{}) (uint, bool) {
	switch i := v.(type) {
	case int:
		return uint(i), true
	case uint:
		return uint(i), true
	case int8:
		return uint(i), true
	case int16:
		return uint(i), true
	case int32:
		return uint(i), true
	case int64:
		return uint(i), true
	case uint8:
		return uint(i), true
	case uint16:
		return uint(i), true
	case uint32:
		return uint(i), true
	case uint64:
		return uint(i), true
	case float32:
		return uint(i), true
	case float64:
		return uint(i), true
	case bool:
		if i {
			return 1, true
		} else {
			return 0, true
		}
	case string:
		f, err := strconv.ParseFloat(i, 64)
		if err != nil {
			return 0, false
		}
		return uint(f), true
	default:
		return 0, false
	}
}

func Uint32(v interface{}) (uint32, bool) {
	switch i := v.(type) {
	case int:
		return uint32(i), true
	case uint:
		return uint32(i), true
	case int8:
		return uint32(i), true
	case int16:
		return uint32(i), true
	case int32:
		return uint32(i), true
	case int64:
		return uint32(i), true
	case uint8:
		return uint32(i), true
	case uint16:
		return uint32(i), true
	case uint32:
		return uint32(i), true
	case uint64:
		return uint32(i), true
	case float32:
		return uint32(i), true
	case float64:
		return uint32(i), true
	case bool:
		if i {
			return 1, true
		} else {
			return 0, true
		}
	case string:
		f, err := strconv.ParseFloat(i, 64)
		if err != nil {
			return 0, false
		}
		return uint32(f), true
	default:
		return 0, false
	}
}

func Uint64(v interface{}) (uint64, bool) {
	switch i := v.(type) {
	case int:
		return uint64(i), true
	case uint:
		return uint64(i), true
	case int8:
		return uint64(i), true
	case int16:
		return uint64(i), true
	case int32:
		return uint64(i), true
	case int64:
		return uint64(i), true
	case uint8:
		return uint64(i), true
	case uint16:
		return uint64(i), true
	case uint32:
		return uint64(i), true
	case uint64:
		return uint64(i), true
	case float32:
		return uint64(i), true
	case float64:
		return uint64(i), true
	case bool:
		if i {
			return 1, true
		} else {
			return 0, true
		}
	case string:
		f, err := strconv.ParseFloat(i, 64)
		if err != nil {
			return 0, false
		}
		return uint64(f), true
	default:
		return 0, false
	}
}

// Float 转成浮点
func Float(v interface{}) (float64, bool) {
	switch i := v.(type) {
	case int:
		return float64(i), true
	case uint:
		return float64(i), true
	case int8:
		return float64(i), true
	case int16:
		return float64(i), true
	case int32:
		return float64(i), true
	case int64:
		return float64(i), true
	case uint8:
		return float64(i), true
	case uint16:
		return float64(i), true
	case uint32:
		return float64(i), true
	case uint64:
		return float64(i), true
	case float32:
		return float64(i), true
	case float64:
		return float64(i), true
	case bool:
		if i {
			return 1, true
		} else {
			return 0, true
		}
	case string:
		f, err := strconv.ParseFloat(i, 64)
		if err != nil {
			return 0, false
		}
		return f, true
	default:
		return 0, false
	}
}

// Bool 转成 Boolean 类型
func Bool(v interface{}) (bool, bool) {
	switch i := v.(type) {
	case int:
		if i > 0 {
			return true, true
		} else {
			return false, true
		}
	case uint:
		if i > 0 {
			return true, true
		} else {
			return false, true
		}
	case int8:
		if i > 0 {
			return true, true
		} else {
			return false, true
		}
	case int16:
		if i > 0 {
			return true, true
		} else {
			return false, true
		}
	case int32:
		if i > 0 {
			return true, true
		} else {
			return false, true
		}
	case int64:
		if i > 0 {
			return true, true
		} else {
			return false, true
		}
	case uint8:
		if i > 0 {
			return true, true
		} else {
			return false, true
		}
	case uint16:
		if i > 0 {
			return true, true
		} else {
			return false, true
		}
	case uint32:
		if i > 0 {
			return true, true
		} else {
			return false, true
		}
	case uint64:
		if i > 0 {
			return true, true
		} else {
			return false, true
		}
	case float32:
		if i > 0 {
			return true, true
		} else {
			return false, true
		}
	case float64:
		if i > 0 {
			return true, true
		} else {
			return false, true
		}
	case bool:
		return i, true
	case string:
		b, err := strconv.ParseBool(i)
		if err != nil {
			return false, false
		}
		return b, true
	default:
		return false, false
	}
}

// Byte 转成 Byte 类型
func Byte(v interface{}) (byte, bool) {
	switch i := v.(type) {
	case int:
		return byte(i), true
	case uint:
		return byte(i), true
	case int8:
		return byte(i), true
	case int16:
		return byte(i), true
	case int32:
		return byte(i), true
	case int64:
		return byte(i), true
	case uint8:
		return byte(i), true
	case uint16:
		return byte(i), true
	case uint32:
		return byte(i), true
	case uint64:
		return byte(i), true
	case float32:
		return byte(i), true
	case float64:
		return byte(i), true
	case bool:
		if i {
			return 1, true
		} else {
			return 0, true
		}
	case string:
		if len(i) > 0 {
			return i[0], true
		}
		return 0, false
	default:
		return 0, false
	}
}

// Bytes 转成 Bytes 类型
func Bytes(v interface{}) ([]byte, bool) {
	switch bytes := v.(type) {
	case []byte:
		return bytes, true
	case string:
		return []byte(bytes), true
	default:
		return nil, false
	}
}

// String 转成 String 类型
func String(v interface{}) (string, bool) {
	switch s := v.(type) {
	case []byte:
		return string(s), true
	case string:
		return s, true
	default:
		return fmt.Sprint(s), true
	}
}

// Time 转成 Time 类型
func Time(v interface{}) (time.Time, bool) {
	switch t := v.(type) {
	case time.Time:
		return t, true
	case *time.Time:
		return *t, true
	case string:
		if tt, err := time.Parse(time.RFC3339, t); err == nil {
			return tt, true
		}
	default:
		return time.Time{}, false
	}
	return time.Time{}, false
}

// Duration 转成 Duration 类型
func Duration(v interface{}) (time.Duration, bool) {
	switch t := v.(type) {
	case time.Duration:
		return t, true
	case int:
		return time.Duration(t), true
	case float64:
		return time.Duration(int(t)), true
	default:
		return 0, false
	}
}

// SliceInt 转成 []int 类型
func SliceInt(v interface{}) ([]int, bool) {
	switch t := v.(type) {
	case []int:
		return t, true
	case []interface{}:
		ints := make([]int, 0, len(t))
		for _, v := range t {
			if n, ok := Int(v); ok {
				ints = append(ints, n)
			}
		}
		return ints, true
	default:
		return nil, false
	}
}

// SliceFloat 转成 []float64 类型
func SliceFloat(v interface{}) ([]float64, bool) {
	switch t := v.(type) {
	case []float64:
		return t, true
	case []float32:
		floats := make([]float64, 0, len(t))
		for _, v := range t {
			floats = append(floats, float64(v))
		}
		return floats, true
	case []interface{}:
		floats := make([]float64, 0, len(t))
		for _, v := range t {
			if f, ok := Float(v); ok {
				floats = append(floats, f)
			}
		}
		return floats, true
	default:
		return nil, false
	}
}

// SliceString 转成 []string 类型
func SliceString(v interface{}) ([]string, bool) {
	switch t := v.(type) {
	case []string:
		return t, true
	case []interface{}:
		ints := make([]string, 0, len(t))
		for _, v := range t {
			if n, ok := String(v); ok {
				ints = append(ints, n)
			}
		}
		return ints, true
	default:
		return nil, false
	}
}

// Map 转成 map 类型
func Map(v interface{}) (map[string]interface{}, bool) {
	switch m := v.(type) {
	case map[string]interface{}:
		return m, true
	case map[interface{}]interface{}:
		var mm = make(map[string]interface{})
		for k, v := range m {
			mm[fmt.Sprint(k)] = v
		}
		return mm, true
	default:
		return nil, false
	}
}

// MapString 转成 map[string]string 类型
func MapString(v interface{}) (map[string]string, bool) {
	switch m := v.(type) {
	case map[string]string:
		return m, true
	case map[interface{}]interface{}:
		var mm = make(map[string]string)
		for k, v := range m {
			s, _ := String(v)
			mm[fmt.Sprint(k)] = s
		}
		return mm, true
	default:
		return nil, false
	}
}
