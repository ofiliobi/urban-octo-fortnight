package vo

// Password structure
type Password struct {
	value string
}

// NewPassword creates a new password
func NewPassword(value string) Password {
	return Password{value: value}
}

// Value returns the value  password
func (p Password) Value() string {
	return p.value
}

// Equals check that two password are equal
func (p Password) Equals(value Value) bool {
	o, ok := value.(Password)
	return ok && p.value == o.value
}