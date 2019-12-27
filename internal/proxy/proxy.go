package proxy

import (
	"golang.org/x/xerrors"
)

// Proxy is domain server of proxy server.
type Proxy struct {
	host     Host
	port     Port
	username Username
	password Password
}

// NewProxy return new Proxy.
func NewProxy(
	host string,
	port string,
	username string,
	password string,
) (*Proxy, error) {
	var err error

	h := new(Host)
	h, err = NewHost(host)
	if err != nil {
		return nil, xerrors.Errorf("Host is wrong: %w", err)
	}

	p := new(Port)
	p, err = NewPort(port)
	if err != nil {
		return nil, xerrors.Errorf("Port is wrong: %w", err)
	}

	u := new(Username)
	u, err = NewUsername(username)
	if err != nil {
		return nil, xerrors.Errorf("Username is wrong: %w", err)
	}

	pw := new(Password)
	pw, err = NewPassword(password)
	if err != nil {
		return nil, xerrors.Errorf("Password is wrong: %w", err)
	}

	return &Proxy{
		host:     *h,
		port:     *p,
		username: *u,
		password: *pw,
	}, nil
}
