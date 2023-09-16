package discovery

import (
	"context"
	"github.com/go-kratos/kratos/v2/selector"
)

// FilterHang is hang state filter.
func FilterHang() selector.NodeFilter {
	return func(_ context.Context, nodes []selector.Node) []selector.Node {
		newNodes := make([]selector.Node, 0)
		for _, n := range nodes {
			if v, ok := n.Metadata()["hang"]; ok && v == "true" {
				continue
			}
			newNodes = append(newNodes, n)
		}
		return newNodes
	}
}
