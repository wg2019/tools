/**
 * @Author: wg2019
 * @Date: 2020/6/12
 * @Desc: types.go，
 */
package amap

const (
	// 计算的方式和方法
	DataType = "type"
	// 请求服务权限标识
	AppKey = "key"
	// 出发点(批量)
	Origins = "origins"
	// 目的地
	Destination = "destination"
	// 返回数据格式类型
	Output = "output"
)

// 逗号分隔符
const CommaSeparator = ","

// 竖线分隔符
const CommaVerticalBar = "|"

// 空字符串
const EmptyString = ""

const (
	// 返回数据格式类型: JSON
	OutputJson = "JSON"
	// 返回数据格式类型: XML
	OutputXml = "XML"
)

// 高德地图请求返回数据格式
// 缺少字段result，调用房补充
type ResponseBaseBody struct {
	Status   string `json:"status"`
	Info     string `json:"info"`
	InfoCode string `json:"infocode"`
}
