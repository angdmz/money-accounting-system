package transactions

import "fmt"

func (it *invalidTransaction) Error() string {
	return fmt.Sprintf("Invalid transaction current: %0.2f; debit amount %f", float64(it.currentBalance), float64(it.creditAmount))
}

type invalidTransaction struct {
	currentBalance uint64
	creditAmount   uint64
}

func NewInvalidTransaction(currentBalance uint64, creditAmount uint64) *invalidTransaction {
	return &invalidTransaction{currentBalance: currentBalance, creditAmount: creditAmount}
}
