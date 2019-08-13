package transactions

type Debit struct {
	amount uint64
}

func (c *Debit) TypeString() string {
	return "debit"
}

func NewDebit(amount uint64) *Debit {
	return &Debit{amount: amount}
}

func (c *Debit) Amount() uint64 {
	return c.amount
}

func (c *Debit) Logic() func(amount uint64) (uint64, error) {
	return func(amount uint64) (uint64, error) {
		if c.amount > amount {
			return 0.0, NewInvalidTransaction(amount, c.amount)
		}
		return amount - c.amount, nil
	}
}
