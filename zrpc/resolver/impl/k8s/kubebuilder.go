package k8s

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/zrpc/resolver/common"
	"github.com/zeromicro/go-zero/zrpc/resolver/impl/k8s/kube"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/proc"
	"github.com/zeromicro/go-zero/core/threading"
	"google.golang.org/grpc/resolver"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

const (
	resyncInterval = 5 * time.Minute
	nameSelector   = "metadata.name="
)

type kubeBuilder struct{}

func init() {
	resolver.Register(&kubeBuilder{})
}

func (b *kubeBuilder) Build(target resolver.Target, cc resolver.ClientConn,
	_ resolver.BuildOptions) (resolver.Resolver, error) {
	svc, err := kube.ParseTarget(target)
	if err != nil {
		return nil, err
	}

	config, err := rest.InClusterConfig()
	if err != nil {
		return nil, err
	}

	cs, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	if svc.Port == 0 {
		endpoints, err := cs.CoreV1().Endpoints(svc.Namespace).Get(context.Background(), svc.Name, v1.GetOptions{})
		if err != nil {
			return nil, err
		}
		svc.Port = int(endpoints.Subsets[0].Ports[0].Port)
	}

	handler := kube.NewEventHandler(func(endpoints []string) {
		var addrs []resolver.Address
		for _, val := range common.Subset(endpoints, common.SubsetSize) {
			addrs = append(addrs, resolver.Address{
				Addr: fmt.Sprintf("%s:%d", val, svc.Port),
			})
		}

		if err := cc.UpdateState(resolver.State{
			Addresses: addrs,
		}); err != nil {
			logx.Error(err)
		}
	})
	inf := informers.NewSharedInformerFactoryWithOptions(cs, resyncInterval,
		informers.WithNamespace(svc.Namespace),
		informers.WithTweakListOptions(func(options *v1.ListOptions) {
			options.FieldSelector = nameSelector + svc.Name
		}))
	in := inf.Core().V1().Endpoints()
	in.Informer().AddEventHandler(handler)
	threading.GoSafe(func() {
		inf.Start(proc.Done())
	})
	endpoints, err := cs.CoreV1().Endpoints(svc.Namespace).Get(context.Background(), svc.Name, v1.GetOptions{})
	if err != nil {
		return nil, err
	}
	handler.Update(endpoints)

	return common.NewNopResolver(cc), nil
}

func (b *kubeBuilder) Scheme() string {
	return common.KubernetesScheme
}
