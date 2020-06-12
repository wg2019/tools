/**
 * @Author: wg2019
 * @Date: 2020/5/29 6:37 下午
 * @Desc:
 */
package slice

import (
	"github.com/goinaction/code/chapter2/sample/search"
	"testing"
)

func TestGetStringValue(t *testing.T) {
	var testStringSlice = Strings{
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
		get := testStringSlice.GetValue(test.pointer)
		if get != test.expect {
			t.Errorf("slice: %+v, pointer: %d, expect: %s, but get %s", testStringSlice, test.pointer, test.expect, get)
		}
	}
}

func TestStrings_RemoteDuplicates(t *testing.T) {
	testCases := []struct {
		oldSlice Strings
		newSlice Strings
	}{
		{Strings{}, Strings{}},
		{Strings{"a"}, Strings{"a"}},
		{Strings{"a", "b"}, Strings{"a", "b"}},
		{Strings{"a", "b", "a"}, Strings{"a", "b"}},
		{Strings{"a", "b", "a", "c", "c"}, Strings{"a", "b", "c"}},
	}
	t.Logf("testCases: %+v", testCases)
	search.Run("aaa")

}
