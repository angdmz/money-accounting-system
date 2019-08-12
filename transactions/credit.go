package transactions

type Credit struct {
	amount float64
}

func NewCredit(amount float64) *Credit {
	return &Credit{amount: amount}
}

func (c *Credit) Amount() float64 {
	return c.amount
}

func (c *Credit) Logic() func(amount float64) (float64, error) {
	return func(amount float64) (float64, error) {
		return amount + c.amount, nil
	}
}
