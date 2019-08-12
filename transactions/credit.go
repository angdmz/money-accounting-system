package transactions

type Credit struct {
	amount float64
	id     int64
}

func NewCredit(amount float64, id int64) *Credit {
	return &Credit{amount: amount, id: id}
}

func (c *Credit) Amount() float64 {
	return c.amount
}

func (c *Credit) Id() int64 {
	return c.id
}

func (c *Credit) Logic() func(amount float64) (float64, error) {
	return func(amount float64) (float64, error) {
		return amount + c.amount, nil
	}
}
