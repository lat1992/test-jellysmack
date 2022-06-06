package mock

import "github.com/stretchr/testify/mock"

type YardServiceExpectations func(*YardService)

type YardService struct {
	mock.Mock
}

func (m *YardService) TakePlace(x, y uint) {
	_ = m.Called(x, y)
}

func (m *YardService) ClearPlace(x, y uint) {
	_ = m.Called(x, y)
}

func (m *YardService) CheckForwardPlace(x, y uint, o string) bool {
	args := m.Called(x, y, o)
	return args.Bool(0)
}

func NewYardService(expect YardServiceExpectations) *YardService {
	s := &YardService{}
	if expect != nil {
		expect(s)
	}
	return s
}
