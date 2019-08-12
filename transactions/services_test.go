package transactions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConstructor(t *testing.T) {
	service := NewTransactionServiceMemory()
	// assert for not nil (good when you expect something)
	if assert.NotNil(t, service) {
		// now we know that object isn't nil, we are safe to make
		// further assertions without causing any errors
		assert.Equal(t, 0.0, service.Balance())
	}
}

func TestAddCreditService(t *testing.T) {
	service := NewTransactionServiceMemory()
	credit := NewCredit(100, 1)
	err := service.AddTransaction(credit)
}

func TestCredit(t *testing.T) {
	c := NewCredit(100, 1)
	if assert.NotNil(t, c) {
		// now we know that object isn't nil, we are safe to make
		// further assertions without causing any errors
		assert.Equal(t, int64(1), c.Id())
		assert.Equal(t, float64(100), c.Amount())
		res, err := c.Logic()(1000)
		assert.Equal(t, float64(1100), res)
		assert.Nil(t, err)
	}
}

func TestDebit(t *testing.T) {
	d := NewDebit(100, 1)
	if assert.NotNil(t, d) {
		// now we know that object isn't nil, we are safe to make
		// further assertions without causing any errors
		assert.Equal(t, int64(1), d.Id())
		assert.Equal(t, float64(100), d.Amount())
		res, err := d.Logic()(1000)
		assert.Equal(t, float64(900), res)
		assert.Nil(t, err)
	}
}

func TestInvalidDebit(t *testing.T) {
	d := NewDebit(100, 1)
	if assert.NotNil(t, d) {
		// now we know that object isn't nil, we are safe to make
		// further assertions without causing any errors
		assert.Equal(t, int64(1), d.Id())
		assert.Equal(t, float64(100), d.Amount())
		res, err := d.Logic()(50)
		assert.Equal(t, float64(0), res)
		if assert.Error(t, err) {
			expectedError := NewInvalidTransaction(50, 100)
			assert.Equal(t, expectedError, err)
		}
	}
}
