/**
 * @Author: wg2019
 * @Date: 2020/6/12
 * @Desc: 创建高德地图客户端及对应配置
 */
package amap

import (
	"net/http"
)

// 高德地图客户端
type Client struct {
	// http客户端
	*http.Client
	// 高德开发者申请的应用key
	appKey string
}

// 创建高德地图客户端
func NewClient(appKey string) *Client {
	c := new(Client)
	c.appKey = appKey
	c.Client = new(http.Client)
	return c
}

// 设置http客户端
func (c *Client) SetHttpClient(httpClient *http.Client) {
	c.Client = httpClient
}
