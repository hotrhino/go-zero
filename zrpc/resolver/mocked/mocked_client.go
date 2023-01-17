package mocked

import (
	"google.golang.org/grpc/resolver"
	"google.golang.org/grpc/serviceconfig"
)

type ClientConn struct {
	State resolver.State
	Err   error
}

func New() *ClientConn {
	return new(ClientConn)
}

func (m *ClientConn) UpdateState(state resolver.State) error {
	m.State = state
	return m.Err
}

func (m *ClientConn) ReportError(err error) {
}

func (m *ClientConn) NewAddress(addresses []resolver.Address) {
}

func (m *ClientConn) NewServiceConfig(serviceConfig string) {
}

func (m *ClientConn) ParseServiceConfig(serviceConfigJSON string) *serviceconfig.ParseResult {
	return nil
}
