package resolver

import (
	"fmt"
	"github.com/zeromicro/go-zero/zrpc/resolver/common"
	"strings"
)

// BuildDirectTarget returns a string that represents the given endpoints with direct schema.
func BuildDirectTarget(endpoints []string) string {
	return fmt.Sprintf("%s:///%s", common.DirectScheme,
		strings.Join(endpoints, common.EndpointSep))
}

// BuildDiscovTarget returns a string that represents the given endpoints with discov schema.
func BuildDiscovTarget(endpoints []string, key string) string {
	return fmt.Sprintf("%s://%s/%s", common.EtcdScheme,
		strings.Join(endpoints, common.EndpointSep), key)
}
