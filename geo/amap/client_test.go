/**
 * @Author: wg2019
 * @Date: 2020/6/12
 * @Desc: amap_test.go，
 */
package amap

import "testing"

const TestAppKey = "test_app_key"

func TestClient_DistanceUrl(t *testing.T) {
	param := new(DistanceParam)
	param.dataType = DistanceDataTypeLine
	param.output = OutputJson
	param.origins = []string{"116.481028,39.98964", "114.481028,39.989643", "115.481028,39.989643"}
	param.destination = "114.465302,40.004717"
	client := NewClient(TestAppKey)
	distance, err := client.Distance(param)
	if err != nil {
		t.Errorf("err: %+v", err)
		return
	}
	t.Logf("distance: %+v", distance)
}
