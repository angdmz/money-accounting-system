package transactions

type Debit struct {
	amount float64
}

func NewDebit(amount float64) *Debit {
	return &Debit{amount: amount}
}

func (c *Debit) Amount() float64 {
	return c.amount
}

func (c *Debit) Logic() func(amount float64) (float64, error) {
	return func(amount float64) (float64, error) {
		if c.amount > amount {
			return 0.0, NewInvalidTransaction(amount, c.amount)
		}
		return amount - c.amount, nil
	}
}
