package transactions

type Transaction interface {
	Amount() float64
	Id() int64
	Logic()
}

type TransactionService interface {
	Balance() float64
}

type TransactionServiceMemory struct {
	transactions map[int64]*Transaction
}
