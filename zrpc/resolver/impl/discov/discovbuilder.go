package discov

import (
	"github.com/zeromicro/go-zero/zrpc/resolver/common"
	"github.com/zeromicro/go-zero/zrpc/resolver/targets"
	"strings"

	"github.com/zeromicro/go-zero/core/discov"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/resolver"
)

type DiscovBuilder struct{}

func init() {
	resolver.Register(&DiscovBuilder{})
}

func (b *DiscovBuilder) Build(target resolver.Target, cc resolver.ClientConn, _ resolver.BuildOptions) (
	resolver.Resolver, error) {
	hosts := strings.FieldsFunc(targets.GetAuthority(target), func(r rune) bool {
		return r == common.EndpointSepChar
	})
	sub, err := discov.NewSubscriber(hosts, targets.GetEndpoints(target))
	if err != nil {
		return nil, err
	}

	update := func() {
		var addrs []resolver.Address
		for _, val := range common.Subset(sub.Values(), common.SubsetSize) {
			addrs = append(addrs, resolver.Address{
				Addr: val,
			})
		}
		if err := cc.UpdateState(resolver.State{
			Addresses: addrs,
		}); err != nil {
			logx.Error(err)
		}
	}
	sub.AddListener(update)
	update()

	return common.NewNopResolver(cc), nil
}

func (b *DiscovBuilder) Scheme() string {
	return common.DiscovScheme
}
