package transactions

import (
	"math/rand"
	"time"
)

type Transaction interface {
	Amount() float64
	Logic() func(amount float64) (float64, error)
	TypeString() string
}

type IdGenerator interface {
	Generate() int64
}

type RandomGenerator struct {
	min int64
	max int64
}

func NewRandomGenerator() *RandomGenerator {
	return &RandomGenerator{}
}

func (u *RandomGenerator) Generate() int64 {
	rand.Seed(time.Now().UnixNano())
	return rand.Int63()
}

type TransactionDTO struct {
	Id           int64  `json:"id,omitempty"`
	Type         string `json:"type,omitempty"`
	EmissionDate string `json:"emissionDate,omitempty"`
	Amount       uint64 `json:"amount,omitempty"`
}
