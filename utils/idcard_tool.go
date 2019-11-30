package utils

import (
	"regexp"
	"strconv"
	"strings"
)

type IDCardTool struct {
}

func NewTool() *IDCardTool {
	return new(IDCardTool)
}

const (
	MOD_FACTOR = 11
	MAX_LEN    = 18
	MIN_LEN    = 15
	VALID_LEN  = 17
)

// 检查身份证号是否符合规则
func (t *IDCardTool) CheckIDCard(no string) bool {
	str := strings.Split(no, "")
	if len(str) == MAX_LEN { // 针对18位的
		pattern := "^([1-8][0-9]{5})([12][0-9])[0-9]{2}([01][0-9][0-3][0-9])[0-9]{3}([0-9X])$" // 检查基本正则
		reg, _ := regexp.Compile(pattern)
		if reg.MatchString(no) {
			// 检查校验位是否正确
			factor := []int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}
			mapFactor := []string{"1", "0", "X", "9", "8", "7", "6", "5", "4", "3", "2", "1"}
			sum := 0
			for i, n := range str[0:VALID_LEN] {
				t, _ := strconv.Atoi(n)
				sum += t * factor[i]
			}
			return mapFactor[sum%MOD_FACTOR] == str[VALID_LEN]
		}
	} else if len(str) == MIN_LEN { // 针对15位的
		reg, _ := regexp.Compile("^([1-8][0-9]{5})[0-9]{2}([01][0-9][0-3][0-9])[0-9]{3}$")
		return reg.MatchString(no)
	}
	return false
}

// 根据身份证号分析户籍住址,出生日期,性别
func (t *IDCardTool) AnalyzeIDCard(no string) map[string]string {
	res := make(map[string]string)
	if t.CheckIDCard(no) {
		noArray := strings.Split(no, "")
		if len(noArray) == MAX_LEN {
			res["birthday"] = strings.Join(noArray[6:14], "")
			i, _ := strconv.Atoi(noArray[MAX_LEN-2])
			res["gender"] = strconv.Itoa(i % 2)
		} else if len(noArray) == MIN_LEN {
			res["birthday"] = "19" + strings.Join(noArray[6:12], "")
			i, _ := strconv.Atoi(noArray[MIN_LEN-1])
			res["gender"] = strconv.Itoa(i % 2)
		}
		res["address"] = strings.Join(noArray[0:6], "")
		res["error"] = "0"

	} else {
		res["error"] = "1"
	}
	return res
}
