package mystring

import (
	"strings"
	"bytes"
	"unicode"
)

// 去除输出字符串中的所有空格，换行和制表符
func Trim(s string) string {
	// 去除空格
	s = strings.Replace(s, " ", "", -1)
	// 去除换行符
	s = strings.Replace(s, "\n", "", -1)
	// 去除制表符
	s = strings.Replace(s, "\t", "", -1)
	// 去除换行符
	s = strings.Replace(s, "\r", "", -1)

	return s
}

// 遍历版本去除输出字符串中的所有空格，换行和制表符
func TrimNew(s string) string {
	var buffer bytes.Buffer
	for _, char := range s {
		if !unicode.IsSpace(char) {
			buffer.WriteRune(char)
		}
	}
	return buffer.String()
}

// 另外一个版本去除输出字符串中的所有空格，换行和制表符
func TrimNew2(s string)  string {
	return strings.Join(strings.FieldsFunc(s, unicode.IsSpace), "")
}
