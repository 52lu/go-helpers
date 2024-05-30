package strutil

import (
	"strings"
	"unicode"
)

/*
* @Description: 将字符串中每个单词的首字母转换为小写
* @Author: LiuQHui
* @Param s
* @Date 2024-05-30 11:10:51
 */
func ToLowerFirstEachWord(s string) string {
	words := strings.Fields(s)
	for i, word := range words {
		if word == "" {
			continue
		}
		runes := []rune(word)
		runes[0] = unicode.ToLower(runes[0])
		words[i] = string(runes)
	}
	return strings.Join(words, " ")
}
