/**
 * @Author: wg
 * @Date: 2020/5/27 12:57 下午
 * @Desc:
 */
package slice

type Strings []string

// @param rows 切片
// @param pointer 位置
// @return value 当前位置的值
func (s Strings) GetValue(pointer int) (value string) {
	if pointer < 0 {
		return ""
	}
	if len(s) > pointer {
		return s[pointer]
	}
	return ""
}
