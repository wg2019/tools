/**
 * @Author: wg
 * @Date: 2020/5/29 6:37 下午
 * @Desc:
 */
package slice

import (
	"testing"
)

func TestGetStringValue(t *testing.T) {
	var testStringSlice = []string{
		"A1",
		"A2",
		"A3",
		"A4",
	}
	testCases := []struct {
		pointer int
		expect  string
	}{
		{-1, ""},
		{0, "A1"},
		{1, "A2"},
		{2, "A3"},
		{3, "A4"},
		{4, ""},
	}
	for _, test := range testCases {
		get := GetStringValue(testStringSlice, test.pointer)
		if get != test.expect {
			t.Errorf("slice: %+v, pointer: %d, expect: %s, but get %s", testStringSlice, test.pointer, test.expect, get)
		}
	}
}
