package testify

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

// require提供了和assert同样的接口，但是遇到错误时，require直接终止测试，而assert返回false。
func TestRequireErrorIs(t *testing.T) {
	var (
		a = errors.New("test")
		b = errors.New("test2")
	)
	require.ErrorIs(t, a, b, "测试ErrorIs")
	TestContains(t)
}
