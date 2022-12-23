package utils

import (
	"strconv"
	"strings"
	"time"
)

func ParseDuration(d string) (time.Duration, error) {
	// 调用 strings.TrimSpace(d) 来删除字符串 d 中的前导和尾随空格。
	d = strings.TrimSpace(d)
	// 调用 time.ParseDuration 函数来尝试将字符串解析为 time.Duration 类型的值
	dr, err := time.ParseDuration(d)
	if err == nil {
		return dr, nil
	}
	// 检查字符串中是否包含字符 "d"
	if strings.Contains(d, "d") {
		index := strings.Index(d, "d")
		// 使用 strconv.Atoi 函数将字符串的前半部分转换为整数
		hour, _ := strconv.Atoi(d[:index])
		dr = time.Hour * 24 * time.Duration(hour)
		// 计算出字符 "d" 后面的字符串所代表的时间间隔
		ndr, err := time.ParseDuration(d[index+1:])
		if err != nil {
			return dr, nil
		}
		// 将两个时间间隔相加。
		return dr + ndr, nil
	}

	dv, err := strconv.ParseInt(d, 10, 64)
	return time.Duration(dv), err
}
