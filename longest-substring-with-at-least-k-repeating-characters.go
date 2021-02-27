package leetcode_goversion

import "strings"

func longestSubstring(s string, k int) int {
	var n = len(s)
	var res = 0
	if n < k {
		return 0
	}
	var countMap = make(map[uint8]int)
	for i := 0; i < n; i++ {
		c := s[i]
		if countMap[c] == 0 {
			countMap[c] = 1
		} else {
			countMap[c] = countMap[c] + 1
		}
	}
	for key, v := range countMap {
		if v < k {
			var split = strings.Split(s, string(key))
			for _, element := range split {
				res = max(res, longestSubstring(element, k))
			}
			return res
		}
	}
	return n
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
