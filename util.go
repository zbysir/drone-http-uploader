package main

import (
	"regexp"
)

// 注入变量
// 如将 ?secret={{SECRET}} 中的${SECRET}替换为SECRET变量的真实值.
// tip: 不要使用${}格式, 这是drone内置的变量替换逻辑,
func injectEnv(src string, data map[string]string) string {
	reg := regexp.MustCompile(`\{\{.+?\}\}`)
	src = reg.ReplaceAllStringFunc(src, func(s string) string {
		key := s[2 : len(s)-2]
		if d, ok := data[key]; ok {
			return d
		}
		return s
	})
	return src
}
