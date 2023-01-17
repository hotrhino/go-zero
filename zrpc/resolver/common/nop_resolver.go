package common

import "google.golang.org/grpc/resolver"

type NopResolver struct {
	cc resolver.ClientConn
}

func NewNopResolver(c resolver.ClientConn) *NopResolver {
	return &NopResolver{
		cc: c,
	}
}

func (r *NopResolver) Close() {
}

func (r *NopResolver) ResolveNow(options resolver.ResolveNowOptions) {
}
