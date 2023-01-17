package etcd

import (
	resolver2 "github.com/zeromicro/go-zero/zrpc/resolver/common"
	"github.com/zeromicro/go-zero/zrpc/resolver/impl/discov"
	"google.golang.org/grpc/resolver"
)

type etcdBuilder struct {
	discov.DiscovBuilder
}

func init() {
	resolver.Register(&etcdBuilder{})
}

func (b *etcdBuilder) Scheme() string {
	return resolver2.EtcdScheme
}
