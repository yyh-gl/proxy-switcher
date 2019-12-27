package proxy

import "golang.org/x/xerrors"

// Host is value object that represent host of proxy server.
type Host struct {
	value string
}

// NewHost create new Host.
func NewHost(val string) (*Host, error) {
	if val == "" {
		return nil, xerrors.New("Host is required")
	}

	return &Host{value: val}, nil
}
