package calc

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

func TestDivideTableDriven(t *testing.T) {
	cases := map[string]struct{ n, d, want float64 }{
		"whole":    {6, 3, 2},
		"negative": {-6, 3, -2},
		"fraction": {1, 8, 0.125},
	}
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			got, err := Divide(c.n, c.d)
			require.NoError(t, err)
			assert.InDelta(t, c.want, got, 1e-9)
		})
	}
}

func TestDivideByZero(t *testing.T) {
	_, err := Divide(1, 0)
	require.ErrorIs(t, err, ErrDivideByZero)
}

// A suite, for the case where several tests share setup.
type CalcSuite struct {
	suite.Suite
	values []float64
}

func (s *CalcSuite) SetupTest() { s.values = []float64{1, 2, 3.5} }

func (s *CalcSuite) TestRunningTotal() {
	s.Equal([]float64{1, 3, 6.5}, RunningTotal(s.values))
}

func TestCalcSuite(t *testing.T) { suite.Run(t, new(CalcSuite)) }
