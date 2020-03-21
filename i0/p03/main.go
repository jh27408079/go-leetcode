package main

import "strings"

func main() {
	str := "abcabcbb"
	l := lengthOfLongestSubstring(str)
	println(l)
}

func lengthOfLongestSubstring(s string) int {
	l := len(s)
	runes := []rune(s)
	temp := ""
	ret := 0
	for i := 0; i < l; i++ {
		if strings.ContainsRune(temp, runes[i]) {
			lt := len(temp)
			ret = max(ret, lt)
			runet := []rune(temp)
			for j := 0; j <= lt; j++ {
				if runet[j] == runes[i] {
					temp = string(runet[j+1:])
					break
				}
			}
		}
		temp += string(runes[i])
	}
	ret = max(ret, len(temp))
	return ret
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func lengthOfLongestSubstring2(s string) int {
	var Length int
	var sl string
	left := 0
	right := 0
	sl = s[left : right]
	for ; right < len(s); right++{
		if index := strings.IndexByte(sl, s[right]); index != -1{
			left += index + 1
		}
		sl = s[left : right + 1]
		if len(sl) > Length{
			Length = len(sl)
		}
	}
	return Length
}