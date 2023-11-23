package main

import (
	"fmt"
	"strings"
)

//func entityParser(text string) string {
//	m := map[string]string{
//		"&quot;":  "\"",
//		"&apos;":  "'",
//		"&gt;":    ">",
//		"&lt;":    "<",
//		"&frasl;": "/",
//	}
//	for k, v := range m {
//		text = strings.ReplaceAll(text, k, v)
//	}
//	return strings.ReplaceAll(text, "&amp;", "&")
//}

func entityParser(t string) string {
	m := map[string]string{
		"&quot;":  "\"",
		"&apos;":  "'",
		"&gt;":    ">",
		"&lt;":    "<",
		"&frasl;": "/",
		"&amp;":   "&",
	}
	n := len(t)
	i := 0
	var sb strings.Builder
	for i < n {
		for i < n && t[i] != '&' {
			sb.WriteRune(rune(t[i]))
			i++
		}
		b := i
		for ; i < n && t[i] != ';'; i++ {
			if t[i] == '&' { // 考虑多个&的情况 "&&gt;"
				sb.WriteString(t[b:i])
				b = i
			}
		}
		e := i
		if i >= n {
			sb.WriteString(t[b:]) // t["&":]
			break
		}
		sub := t[b : e+1]
		if tar, ok := m[sub]; ok {
			sb.WriteString(tar)
		} else {
			sb.WriteString(sub)
		}
		i++
	}
	return sb.String()
}

func main() {
	s := entityParser("&amp; is an HTML entity but &ambassador; is not.")
	fmt.Printf(s)
}
