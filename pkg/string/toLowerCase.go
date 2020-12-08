package string

/*
ToLowerCase exported function
https://leetcode-cn.com/problems/to-lower-case/
将字符串中的大写转化为小写
*/
func ToLowerCase(str string) string {
	b := []byte(str)
	for i := 0; i < len(b); i++ {
		if str[i] >= 'A' && str[i] <= 'Z' {
			b[i] = str[i] + 32
			continue
		}
		b[i] = str[i]
	}
	return string(b)
}
