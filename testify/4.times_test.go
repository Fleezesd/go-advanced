package testify

import (
	"testing"

	"github.com/stretchr/testify/mock"
)

// Times(n int) 可以精确断言某方法以特定参数的调用次数
type MockExample struct {
	mock.Mock
}

func (e *MockExample) Hello(n int) int {
	args := e.Mock.Called(n)
	return args.Int(0)
}

func TestExample(t *testing.T) {
	e := new(MockExample)

	e.On("Hello", 1).Return(1).Times(1)
	e.On("Hello", 2).Return(2).Times(2)
	e.On("Hello", 3).Return(3).Times(3)

	ExampleFunc(e)

	e.AssertExpectations(t)
}
