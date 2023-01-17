package k8s

import (
	"github.com/zeromicro/go-zero/zrpc/resolver/common"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKubeBuilder_Scheme(t *testing.T) {
	var b kubeBuilder
	assert.Equal(t, common.KubernetesScheme, b.Scheme())
}
