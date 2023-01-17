package direct

import (
	"github.com/zeromicro/go-zero/zrpc/resolver/common"
	"github.com/zeromicro/go-zero/zrpc/resolver/targets"
	"strings"

	"google.golang.org/grpc/resolver"
)

type directBuilder struct{}

func init() {
	resolver.Register(&directBuilder{})
}

func (d *directBuilder) Build(target resolver.Target, cc resolver.ClientConn, _ resolver.BuildOptions) (
	resolver.Resolver, error) {
	endpoints := strings.FieldsFunc(targets.GetEndpoints(target), func(r rune) bool {
		return r == common.EndpointSepChar
	})
	endpoints = common.Subset(endpoints, common.SubsetSize)
	addrs := make([]resolver.Address, 0, len(endpoints))

	for _, val := range endpoints {
		addrs = append(addrs, resolver.Address{
			Addr: val,
		})
	}
	if err := cc.UpdateState(resolver.State{
		Addresses: addrs,
	}); err != nil {
		return nil, err
	}

	return common.NewNopResolver(cc), nil
}

func (d *directBuilder) Scheme() string {
	return common.DirectScheme
}
