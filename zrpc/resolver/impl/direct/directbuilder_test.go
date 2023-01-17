package direct

import (
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/zrpc/resolver/common"
	"github.com/zeromicro/go-zero/zrpc/resolver/mocked"
	"net/url"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zeromicro/go-zero/core/lang"
	"github.com/zeromicro/go-zero/core/mathx"
	"google.golang.org/grpc/resolver"
)

func TestDirectBuilder_Build(t *testing.T) {
	tests := []int{
		0,
		1,
		2,
		common.SubsetSize / 2,
		common.SubsetSize,
		common.SubsetSize * 2,
	}

	for _, test := range tests {
		test := test
		t.Run(strconv.Itoa(test), func(t *testing.T) {
			var servers []string
			for i := 0; i < test; i++ {
				servers = append(servers, fmt.Sprintf("localhost:%d", i))
			}
			var b directBuilder
			cc := mocked.New()
			target := fmt.Sprintf("%s:///%s", common.DirectScheme, strings.Join(servers, ","))
			uri, err := url.Parse(target)
			assert.Nil(t, err)
			cc.Err = errors.New("foo")
			_, err = b.Build(resolver.Target{
				URL: *uri,
			}, cc, resolver.BuildOptions{})
			assert.NotNil(t, err)
			cc.Err = nil
			_, err = b.Build(resolver.Target{
				URL: *uri,
			}, cc, resolver.BuildOptions{})
			assert.NoError(t, err)

			size := mathx.MinInt(test, common.SubsetSize)
			assert.Equal(t, size, len(cc.State.Addresses))
			m := make(map[string]lang.PlaceholderType)
			for _, each := range cc.State.Addresses {
				m[each.Addr] = lang.Placeholder
			}
			assert.Equal(t, size, len(m))
		})
	}
}

func TestDirectBuilder_Scheme(t *testing.T) {
	var b directBuilder
	assert.Equal(t, common.DirectScheme, b.Scheme())
}
