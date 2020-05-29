/**
 * @Author: wg
 * @Date: 2020/5/27 12:57 下午
 * @Desc:
 */
package slice

// @param rows 切片
// @param pointer 位置
// @return value 当前位置的值
func GetStringValue(rows []string, pointer int) (value string) {
	if pointer < 0 {
		return ""
	}
	if len(rows) > pointer {
		return rows[pointer]
	}
	return ""
}
