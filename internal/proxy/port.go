package proxy

import "golang.org/x/xerrors"

// Port is value object that represent host of proxy server.
type Port struct {
	value string
}

// NewPort create new Host.
func NewPort(val string) (*Port, error) {
	if val == "" {
		return nil, xerrors.New("Port is required")
	}

	return &Port{value: val}, nil
}
