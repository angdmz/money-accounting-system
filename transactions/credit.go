package transactions

type Credit struct {
	amount uint64
}

func (c *Credit) TypeString() string {
	return "credit"
}

func NewCredit(amount uint64) *Credit {
	return &Credit{amount: amount}
}

func (c *Credit) Amount() uint64 {
	return c.amount
}

func (c *Credit) Logic() func(amount uint64) (uint64, error) {
	return func(amount uint64) (uint64, error) {
		return amount + c.amount, nil
	}
}
