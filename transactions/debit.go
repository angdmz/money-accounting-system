package transactions

type Debit struct {
	amount float64
	id     int64
}

func NewDebit(amount float64, id int64) *Debit {
	return &Debit{amount: amount, id: id}
}

func (c *Debit) Amount() float64 {
	return c.amount
}

func (c *Debit) Id() int64 {
	return c.id
}

func (c *Debit) Logic() func(amount float64) (float64, error) {
	return func(amount float64) (float64, error) {
		if c.amount > amount {
			return 0.0, NewInvalidTransaction(amount, c.amount)
		}
		return amount - c.amount, nil
	}
}
