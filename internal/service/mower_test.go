package service

import (
	"jellysmack-test/internal/mock"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Run(t *testing.T) {
	tests := []struct {
		name        string
		expect      mock.YardServiceExpectations
		x           int
		y           int
		o           string
		instruction string
		want        string
	}{
		{
			name: "normal use case",
			expect: func(m *mock.YardService) {
				m.On("TakePlace", uint(1), uint(2))
				m.On("CheckForwardPlace", uint(1), uint(2), "W").Return(false)
				m.On("ClearPlace", uint(1), uint(2))
				m.On("TakePlace", uint(0), uint(2))
			},
			x:           1,
			y:           2,
			o:           "N",
			instruction: "LF",
			want:        "0 2 W",
		},
		{
			name: "long use case",
			expect: func(m *mock.YardService) {
				m.On("TakePlace", uint(1), uint(2))
				m.On("CheckForwardPlace", uint(1), uint(2), "W").Return(false)
				m.On("ClearPlace", uint(1), uint(2))
				m.On("TakePlace", uint(0), uint(2))
				m.On("CheckForwardPlace", uint(0), uint(2), "S").Return(false)
				m.On("ClearPlace", uint(0), uint(2))
				m.On("TakePlace", uint(0), uint(1))
				m.On("CheckForwardPlace", uint(0), uint(1), "E").Return(false)
				m.On("ClearPlace", uint(0), uint(1))
				m.On("TakePlace", uint(1), uint(1))
				m.On("CheckForwardPlace", uint(1), uint(1), "N").Return(false)
				m.On("ClearPlace", uint(1), uint(1))
				m.On("TakePlace", uint(1), uint(3))
				m.On("CheckForwardPlace", uint(1), uint(2), "N").Return(false)
			},
			x:           1,
			y:           2,
			o:           "N",
			instruction: "LFLFLFLFF",
			want:        "1 3 N",
		},
		{
			name: "cannot move",
			expect: func(m *mock.YardService) {
				m.On("TakePlace", uint(1), uint(2))
				m.On("CheckForwardPlace", uint(1), uint(2), "W").Return(true)
			},
			x:           1,
			y:           2,
			o:           "N",
			instruction: "LF",
			want:        "1 2 W",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			yard := mock.NewYardService(tt.expect)
			s, err := NewMowerService(tt.x, tt.y, tt.o, tt.instruction, yard)
			assert.Nil(t, err)
			result := s.Run()
			assert.Equal(t, tt.want, result)
		})
	}
}
