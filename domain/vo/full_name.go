package vo

// FullName structure
type FullName struct {
	value string
}

// NewFullName creates a new FullName
func NewFullName(value string) FullName {
	return FullName{
		value: value,
	}
}

// Value returns the value FullName
func (p FullName) Value() string {
	return p.value
}

// Equals check that two FullName are the same
func (p FullName) Equals(value Value) bool {
	o, ok := value.(FullName)
	return ok && p.value == o.value
}

