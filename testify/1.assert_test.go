package testify

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// reference: https://darjun.github.io/2021/08/11/godailylib/testify/

func TestEqual(t *testing.T) {
	var (
		a = 100
		b = 200
	)
	// test, expected, actual, message
	assert.Equal(t, a, b, "测试AB数值")
}

func TestContains(t *testing.T) {
	var (
		a = []int{1, 2, 3}
		b = 4
	)
	// test 原串 子串 msg
	assert.Contains(t, a, b, "测试Contains")
}

func TestDirExists(t *testing.T) {
	var (
		dirPath = "/"
	)
	assert.DirExists(t, dirPath, "测试目录是否存在")
}

/*
常用函数
	1. ElementsMatch(t TestingT, listA, listB interface{}, msgAndArgs ...interface{}) bool
		断言listA和listB包含相同的元素，忽略元素出现的顺序
	2. Empty(t TestingT, object interface{}, msgAndArgs ...interface{}) bool
		断言object是空 (对应类型的空值)
	3. EqualError(t TestingT, theError error, errString string, msgAndArgs ...interface{}) bool
		断言theError.Error()的返回值与errString相等。
	4. EqualValues(t TestingT, expected, actual interface{}, msgAndArgs ...interface{}) bool
		断言expected与actual相等，或者可以转换为相同的类型，并且相等。这个条件比Equal更宽
	5. Error(t TestingT, err error, msgAndArgs ...interface{}) bool
		断言err不为nil
	6. ErrorAs(t TestingT, err error, target interface{}, msgAndArgs ...interface{}) bool
		ErrorAs断言err表示的 error 链中至少有一个和target匹配。这个函数是对标准库中errors.As的包装
	7. ErrorIs(t TestingT, err, target error, msgAndArgs ...interface{}) bool
		ErrorIs断言err的 error 为对应target。

	逆断言带上Not即可
*/

type MyInt int

func TestEqualValues(t *testing.T) {

	var a = 100
	var b MyInt = 100
	assert.EqualValues(t, a, b, "测试EqualValues 转换成相同类型相等即可")
}

func TestErrorAs(t *testing.T) {

	var (
		a = errors.New("test")
		b = fmt.Errorf("包装 %v", a)
	)
	// 此处坑点: 注意包装error target 指定对应指针
	assert.ErrorAs(t, a, &b, "测试ErrorAs")
}
func TestErrorIs(t *testing.T) {
	var (
		a = errors.New("test")
		b = a
	)
	assert.ErrorIs(t, a, b, "测试ErrorIs")
}
