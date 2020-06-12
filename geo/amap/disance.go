/**
 * @Author: wg2019
 * @Date: 2020/6/12
 * @Desc:
 */
package amap

import (
	"encoding/json"
	"net/url"
	"strings"
)

const (
	// 0：直线距离
	DistanceDataTypeLine = "0"
	// 1：驾车导航距离（仅支持国内坐标）
	// 必须指出，当为1时会考虑路况，故在不同时间请求返回结果可能不同
	// 此策略和驾车路径规划接口的 strategy=4策略基本一致，策略为“ 躲避拥堵的路线，但是可能会存在绕路的情况，耗时可能较长
	// 若需要实现高德地图客户端效果，可以考虑使用驾车路径规划接口
	DistanceDataTypeDriving = "1"
	// 3：步行规划距离（仅支持5km之间的距离）
	DistanceDataTypeWalking = "2"
)

// 距离测量请求参数
type DistanceParam struct {
	origins     []string
	destination string
	dataType    string
	output      string
}

// 距离测量响应结果详情数据
type DistanceResponseResultDetail struct {
	OriginId string `json:"origin_id"`
	DestId   string `json:"dest_id"`
	Distance string `json:"distance"`
	Duration string `json:"duration"`
}

// 距离测量响应结果数据
type DistanceResponseBody struct {
	ResponseBaseBody
	Results []DistanceResponseResultDetail `json:"results"`
}

// 距离测量请求方法
func (c *Client) Distance(param *DistanceParam) (*DistanceResponseBody, error) {
	// 频率控制
	// 调用方法 AddFrequency(DistancePathV3,频率)
	frequency(DistancePathV3)
	// 调用高德地图距离测量API
	response, err := c.Get(c.distanceUrl(param))
	if err != nil {
		return nil, err
	}
	// 解析获得结果数据
	result := new(DistanceResponseBody)
	err = json.NewDecoder(response.Body).Decode(result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// 获取请求地址
func (c *Client) distanceUrl(param *DistanceParam) string {
	requestUrl := new(url.URL)
	requestUrl.Host = RestApiHost
	requestUrl.Scheme = RestApiScheme
	requestUrl.Path = DistancePathV3
	requestUrl.RawQuery = c.distanceParam(param).Encode()
	return requestUrl.String()
}

// 获取请求参数
func (c *Client) distanceParam(param *DistanceParam) url.Values {
	urlParam := make(url.Values)
	urlParam.Set(Origins, strings.Join(param.origins, CommaVerticalBar))
	urlParam.Set(Destination, param.destination)
	urlParam.Set(DataType, param.dataType)
	urlParam.Set(AppKey, c.appKey)
	if param.output == EmptyString {
		param.output = OutputJson
	}
	urlParam.Set(Output, param.output)
	return urlParam
}
