package vo

// Money structure
type Money struct {
	currency 	Currency
	amount		Amount
}

// NewMoney creates new Money
func NewMoney(currency Currency, amount Amount) Money {
	return Money{
		currency: currency,
		amount: amount,
	}
}

// NewMoneyNGN creates new Money with currency NGN
func NewMoneyNGN(amount Amount) Money {
	return Money{
		currency: Currency{value: NGN},
		amount: amount,
	}
}

// NewMoneyUSD creates new Money with currency USD
func NewMoneyUSD(amount Amount) Money {
	return Money{
		currency: Currency{value: USD},
		amount: amount,
	}
}

// Amount return value Amount
func (m Money) Amount() Amount {
	return m.amount
}

// Currency return value Currency
func (m Money) Currency() Currency {
	return m.currency
}

// Add adds value in Amount
func (m Money) Add(amount Amount) Money {
	return Money{
		currency: m.currency,
		amount:		Amount{value: m.amount.Value() + amount.Value()},
	}
}

// Sub subtracts value from Amount
func (m Money) Sub(amount Amount) Money {
	return Money{
		currency: m.currency,
		amount:	Amount{value: m.amount.Value() - amount.Value()},
	}
}

// Equals check that two Money are the same
func (m Money) Equals(value Value) bool {
	o, ok := value.(Money)
	return ok && m.amount == o.amount && m.currency == o.currency
}