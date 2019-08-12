package transactions

import "fmt"

func (it *invalidTransaction) Error() string {
	return fmt.Sprintf("Invalid transaction current: %0.2f; debit amount %f", it.currentBalance, it.creditAmount)
}

type invalidTransaction struct {
	currentBalance float64
	creditAmount   float64
}

func NewInvalidTransaction(currentBalance float64, creditAmount float64) *invalidTransaction {
	return &invalidTransaction{currentBalance: currentBalance, creditAmount: creditAmount}
}
