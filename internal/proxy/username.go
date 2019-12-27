package proxy

import "golang.org/x/xerrors"

// Username is value object that represent username of login information.
type Username struct {
	value string
}

// NewUsername create new Username.
func NewUsername(val string) (*Username, error) {
	if val == "" {
		return nil, xerrors.New("Username is required")
	}

	return &Username{value: val}, nil
}
