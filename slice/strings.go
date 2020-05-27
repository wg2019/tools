/**
 * @Author: wg
 * @Date: 2020/5/27 12:57 下午
 * @Desc:
 */
package slice

// @param rows 切片
// @param pointer 位置
// @return value 当前位置的值
func GetValue(rows []string, pointer int) (value string) {
	if len(rows) > pointer {
		return rows[pointer]
	}
	return ""
}
