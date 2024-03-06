package tool

import (
	"regexp"
	"time"
)

// CheckStrSliDup 返回字符串slice中所有重复的字符串
func CheckStrSliDup(ss []string) []string {
	var ds []string
	set := make(map[string]struct{})
	for _, s := range ss {
		if _, ok := set[s]; ok {
			ds = append(ds, s)
		} else {
			set[s] = struct{}{}
		}
	}
	return ds
}

// CheckDateStr 判断dateStr是否是YYYYMMDD格式，是否合法，是否不早于今天
func CheckDateStr(dateStr string) bool {
	t, err := time.Parse("20060102", dateStr)
	if err != nil || t.After(time.Now()) {
		return false
	}

	return true
}

// CheckEmail 判断email是否满足邮箱格式
func CheckEmail(email string) bool {
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	matched, err := regexp.MatchString(pattern, email)
	if err != nil {
		return false
	}
	return matched
}
