/**
 * @Author: wg
 * @Date: 2020/5/27 12:57 下午
 * @Desc:
 */
package slice

// strings slice
type Strings []string

// rows 切片
// pointer 位置
// value 当前位置的值
func (s Strings) GetValue(pointer int) (value string) {
	if pointer < 0 {
		return ""
	}
	if len(s) > pointer {
		return s[pointer]
	}
	return ""
}

// remote duplicates, return new strings slice
func (s Strings) RemoteDuplicates() (newS Strings) {
	stringsMap := make(map[string]int)
	newS = make([]string, 0)
	for pointer, value := range s {
		if _, ok := stringsMap[value]; !ok {
			newS = append(newS, value)
		}
		stringsMap[value] = pointer
	}
	return newS
}
