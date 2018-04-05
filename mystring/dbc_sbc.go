package mystring

import "bytes"

// 该工具类用两种方式实现了字符串全角转半角，其中 DBCtoSBCNew 性能更好，详细的 bench test 数据见测试文件

//全角转换半角
func DBCtoSBC(s string) string {
	retstr := ""
	for _, i := range s {
		inside_code := i
		if inside_code == 12288 {
			inside_code = 32
		} else {
			inside_code -= 65248
		}
		if inside_code < 32 || inside_code > 126 {
			retstr += string(i)
		} else {
			retstr += string(inside_code)
		}
	}
	return retstr
}

//全角转换半角
func DBCtoSBCNew(s string) string {
	buffer := bytes.Buffer{}
	for _, i := range s {
		inside_code := i
		if inside_code == 12288 {
			inside_code = 32
		} else {
			inside_code -= 65248
		}
		if inside_code < 32 || inside_code > 126 {
			buffer.WriteRune(i)
		} else {
			buffer.WriteRune(inside_code)
		}
	}
	return buffer.String()
}