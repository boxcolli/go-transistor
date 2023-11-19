package tools

import "math/big"

var (
	intOne = big.NewInt(1)
	intZero = big.NewInt(0)
	floatOne = big.NewFloat(1)
)

type Counter struct {
	I 	*big.Int
	F	*big.Float
	Mod	*big.Int
}

func NewCounter(mod int64) *Counter {
	return &Counter{
		I:		big.NewInt(0),
		F:		big.NewFloat(0),
		Mod:	big.NewInt(mod),
	}
}
func (c *Counter) Reset() {
	c.I = big.NewInt(0)
	c.F = big.NewFloat(0)
}
func (c *Counter) AddOne() {
	c.I.Add(c.I, intOne)
	c.F.Add(c.F, floatOne)
}
func (c *Counter) Add(x int64) {
	c.I.Add(c.I, big.NewInt(x))
	c.F.Add(c.F, big.NewFloat(float64(x)))
}
func (c *Counter) Quo(x float64) *big.Float {
	d := big.NewFloat(x)
	return d.Quo(c.F, d)
}
func (c *Counter) IsModZero() bool {
	return big.NewInt(0).Mod(c.I, c.Mod).Cmp(intZero) == 0
}
func (c *Counter) String() string {
	return c.I.String()
}
