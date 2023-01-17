package common

import "fmt"

const (
	// DirectScheme stands for direct scheme.
	DirectScheme = "direct"
	// DiscovScheme stands for discov scheme.
	DiscovScheme = "discov"
	// EtcdScheme stands for etcd scheme.
	EtcdScheme = "etcd"
	// KubernetesScheme stands for k8s scheme.
	KubernetesScheme = "k8s"
	// EndpointSepChar is the separator cha in endpoints.
	EndpointSepChar = ','
)

var EndpointSep = fmt.Sprintf("%c", EndpointSepChar)
