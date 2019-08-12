package transactions

type Transaction interface {
	Amount() float64
	Id() int64
	Logic() func(amount float64) (float64, error)
}

type TransactionService interface {
	Balance() float64
	Transact(t *Transaction) error
}

type TransactionServiceMemory struct {
	transactions map[int64]*Transaction
	balance      float64
}

func (tsm *TransactionServiceMemory) Transact(t *Transaction) error {
	res, err := t.Logic()(tsm.Balance())
}

func (t *TransactionServiceMemory) Balance() float64 {
	return t.balance
}

func NewTransactionServiceMemory() *TransactionServiceMemory {
	return &TransactionServiceMemory{balance: 0}
}

func NewTransactionService() *TransactionService {
	return NewTransactionServiceMemory()
}
