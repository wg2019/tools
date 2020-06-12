/**
 * @Author: wg2019
 * @Date: 2020/6/12
 * @Desc: amap，频率限制
 */
package amap

import "time"

// 漏桶映射，对应每一个请求
var frequencyMap map[string]frequencyQueue

// 漏桶，控制访问频率
type frequencyQueue <-chan time.Time

// 初始化漏桶映射
func init() {
	frequencyMap = make(map[string]frequencyQueue)
}

// 获取对应的漏桶
// 获取失败则不进行频率限制
func frequency(apiPath string) {
	if frequencyQueue, ok := frequencyMap[apiPath]; ok {
		<-frequencyQueue
	}
}

func AddFrequency(apiPath string, frequencyQueue <-chan time.Time) {
	frequencyMap[apiPath] = frequencyQueue
}
