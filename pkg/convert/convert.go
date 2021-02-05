package convert

import "strconv"

// 处理接口返回的响应

type StrTo string

// 类型转换
// 字符串
func (s StrTo) String() string {
	return string(s)
}

// int
func (s StrTo) Int() (int, error) {
	i, err := strconv.Atoi(s.String())
	return i, err
}

func (s StrTo) MustInt() int {
	i, _ := s.Int()
	return i
}

// UInt32
func (s StrTo) UInt32() (uint32, error) {
	i, err := strconv.Atoi(s.String())
	return uint32(i), err
}

func (s StrTo) MustUInt32() uint32 {
	i, _ := s.UInt32()
	return i
}
