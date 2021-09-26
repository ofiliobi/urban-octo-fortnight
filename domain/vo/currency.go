package vo

import "errors"

const (
	//Currency types
	NGN TypeCurrency = "NGN"
	USD TypeCurrency = "USD"
	GBP TypeCurrency = "GBP"
)

type (
	// TypeCurrency define currency type
	TypeCurrency string
)

func (tc TypeCurrency) String() string {
	return string(tc)
}

var (
	ErrInvalidCurrency = errors.New("invalid currency")
)

type (
	// Currency structure
	Currency struct {
		value TypeCurrency
	}
)

// NewCurrency creates new currency
func NewCurrency(value string) (Currency, error) {
	var c = Currency{value: TypeCurrency(value)}

	if !c.validate() {
		return Currency{}, ErrInvalidCurrency
	}
	return c, nil
}

func (c Currency) validate() bool {
	switch c.value {
	case NGN, USD, GBP:
		return true
	}
	return false
}

// Value return value Currency
func (c Currency) Value() TypeCurrency {
	return c.value
}

// String returns string representation of the Currency
func (c Currency) String() string {
	return string(c.value)
}

// Equals check that two Currency are the same
func (c Currency) Equals(value Value) bool {
	o, ok := value.(Currency)
	return ok && c.value == o.value
}