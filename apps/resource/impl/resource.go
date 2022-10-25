package impl

import (
	"context"
	"github.com/lifangjunone/cmdb/apps/resource"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *service) Search(ctx context.Context, req *resource.SearchRequest) (*resource.ResourceSet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Search not implemented")
}
func (s *service) QueryTag(ctx context.Context, req *resource.QueryTagRequest) (*resource.TagSet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryTag not implemented")
}
func (s *service) UpdateTag(ctx context.Context, req *resource.UpdateTagRequest) (*resource.Resource, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTag not implemented")
}
