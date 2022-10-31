package impl

import (
	"context"
	"github.com/infraboard/mcube/exception"
	"github.com/lifangjunone/cmdb/apps/host"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *service) SyncHost(ctx context.Context, req *host.Host) (*host.Host, error) {
	exist, err := s.DescribeHost(ctx, host.NewDescribeHostRequestWithID(req.Base.Id))
	if err != nil {
		// 如果不是Not Found则直接返回
		if !exception.IsNotFoundError(err) {
			return nil, err
		}
	}

	// 检查ins已经存在 我们则需要更新ins
	if exist != nil {
		s.log.Debugf("update host: %s", req.Base.Id)
		exist.Put(host.NewUpdateHostDataByIns(req))
		if err := s.update(ctx, exist); err != nil {
			return nil, err
		}
		return req, nil
	}

	// 如果没有我们则直接保存
	s.log.Debugf("insert host: %s", req.Base.Id)
	if err := s.save(ctx, req); err != nil {
		return nil, err
	}

	return req, nil
}

func (s *service) QueryHost(ctx context.Context, req *host.QueryHostRequest) (*host.HostSet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryHost not implemented")
}
func (s *service) DescribeHost(ctx context.Context, req *host.DescribeHostRequest) (*host.Host, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeHost not implemented")
}
func (s *service) UpdateHost(ctx context.Context, req *host.UpdateHostRequest) (*host.Host, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateHost not implemented")
}
func (s *service) ReleaseHost(ctx context.Context, req *host.ReleaseHostRequest) (*host.Host, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReleaseHost not implemented")
}
