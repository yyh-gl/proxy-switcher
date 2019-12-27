package proxy

import "golang.org/x/xerrors"

// Password is value object that represent Password of login information.
type Password struct {
	value string
}

// NewPassword create new Password.
func NewPassword(val string) (*Password, error) {
	if val == "" {
		return nil, xerrors.New("Password is required")
	}

	return &Password{value: val}, nil
}
