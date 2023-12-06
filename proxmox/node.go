package proxmox

import (
	"context"

	"github.com/k8s-proxmox/proxmox-go/api"
)

func (s *Service) Nodes(ctx context.Context) ([]*api.Node, error) {
	return s.restclient.GetNodes(ctx)
}

func (s *Service) Node(ctx context.Context, name string) (*api.Node, error) {
	return s.restclient.GetNode(ctx, name)
}
