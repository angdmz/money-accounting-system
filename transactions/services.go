package transactions

import (
	"errors"
	"github.com/satori/go.uuid"
)

type Transaction interface {
	Amount() float64
	Logic() func(amount float64) (float64, error)
}

type TransactionService interface {
	Balance() float64
	Transact(t Transaction) (string, error)
}

type TransactionServiceMemory struct {
	transactions map[int64]*Transaction
	balance      float64
	generator    IdGenerator
}

func (tsm *TransactionServiceMemory) Transact(t Transaction) (string, error) {
	res, err := t.Logic()(tsm.Balance())
	if err == nil {
		tsm.balance = res
	}
	return tsm.generator.Generate(), err
}

func (tsm *TransactionServiceMemory) Balance() float64 {
	return tsm.balance
}

func NewTransactionServiceMemory(ig IdGenerator) *TransactionServiceMemory {
	return &TransactionServiceMemory{balance: 0, generator: ig}
}

func NewTransactionService() TransactionService {
	return NewTransactionServiceMemory(NewUuidGenerator())
}

type IdGenerator interface {
	Generate() string
}

type UuidGenerator struct {
}

func (u *UuidGenerator) Generate() string {
	gen := uuid.Must(uuid.NewV4())
	return gen.String()
}

func NewUuidGenerator() *UuidGenerator {
	return &UuidGenerator{}
}

func ProcessTransactionType(tx TransactionDTO) (Transaction, error) {
	if tx.Type == "debit" {
		return NewDebit(tx.Amount), nil
	} else if tx.Type == "credit" {
		return NewCredit(tx.Amount), nil
	} else {
		return nil, errors.New("Invalid type of transaction")
	}
}

type TransactionDTO struct {
	Type   string  `json:"type,omitempty"`
	Amount float64 `json:"amount,omitempty"`
}
