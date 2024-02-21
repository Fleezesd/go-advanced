package testify

import (
	"testing"

	"github.com/stretchr/testify/mock"
)

// mock User返回数据
type MockCrawler struct {
	mock.Mock
}

var (
	MockUsers []*User
)

func init() {
	MockUsers = append(MockUsers, &User{"DJ", 18}, &User{"ZhangSan", 20})
}

func (m *MockCrawler) GetUserList() ([]*User, error) {
	args := m.Called()
	return args.Get(0).([]*User), args.Error(1) // 返回值指定
}

func TestGetAndPrintUsers(t *testing.T) {
	crawler := new(MockCrawler)
	crawler.On("GetUserList").Return(MockUsers, nil) // 指示调用GetUserList()方法的返回值分别为MockUsers和nil

	GetAndPrintUsers(crawler)
	crawler.AssertExpectations(t) // 确保了模拟对象上的方法按照预期被调用，
}
