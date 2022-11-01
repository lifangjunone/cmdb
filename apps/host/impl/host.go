package impl

import (
	"context"
	"database/sql"
	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/sqlbuilder"
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
	query := sqlbuilder.NewQuery(queryHostSQL)

	if req.Keywords != "" {
		query.Where("r.name LIKE ? OR r.id = ? OR r.instance_id = ? OR r.private_ip LIKE ? OR r.public_ip LIKE ?",
			"%"+req.Keywords+"%",
			req.Keywords,
			req.Keywords,
			req.Keywords+"%",
			req.Keywords+"%",
		)
	}

	querySQL, args := query.Order("sync_at").Desc().Limit(req.Page.ComputeOffset(), uint(req.Page.PageSize)).BuildQuery()
	s.log.Debugf("sql: %s", querySQL)

	queryStmt, err := s.db.Prepare(querySQL)
	if err != nil {
		return nil, exception.NewInternalServerError("prepare query host error, %s", err.Error())
	}
	defer queryStmt.Close()

	rows, err := queryStmt.Query(args...)
	if err != nil {
		return nil, exception.NewInternalServerError(err.Error())
	}
	defer rows.Close()

	set := host.NewHostSet()
	var (
		publicIPList, privateIPList, keyPairNameList, securityGroupsList string
	)
	for rows.Next() {
		ins := host.NewDefaultHost()
		base := ins.Base
		info := ins.Information
		desc := ins.Describe
		err := rows.Scan(
			&base.Id, &base.Vendor, &base.Region, &base.Zone, &base.CreateAt, &info.ExpireAt,
			&info.Category, &info.Type, &info.Name, &info.Description,
			&info.Status, &info.UpdateAt, &base.SyncAt, &info.SyncAccount,
			&publicIPList, &privateIPList, &info.PayType, &base.DescribeHash, &base.ResourceHash,
			&base.SecretId, &base.Id,
			&desc.Cpu, &desc.Memory, &desc.GpuAmount, &desc.GpuSpec, &desc.OsType, &desc.OsName,
			&desc.SerialNumber, &desc.ImageId, &desc.InternetMaxBandwidthOut, &desc.InternetMaxBandwidthIn,
			&keyPairNameList, &securityGroupsList,
		)
		if err != nil {
			return nil, exception.NewInternalServerError("query host error, %s", err.Error())
		}
		info.LoadPrivateIPString(privateIPList)
		info.LoadPublicIPString(publicIPList)
		desc.LoadKeyPairNameString(keyPairNameList)
		desc.LoadSecurityGroupsString(securityGroupsList)
		set.Add(ins)
	}

	// 获取total SELECT COUNT(*) FROMT t Where ....
	countSQL, args := query.BuildCount()
	countStmt, err := s.db.Prepare(countSQL)
	if err != nil {
		return nil, exception.NewInternalServerError(err.Error())
	}

	defer countStmt.Close()
	err = countStmt.QueryRow(args...).Scan(&set.Total)
	if err != nil {
		return nil, exception.NewInternalServerError(err.Error())
	}

	return set, nil
}

func (s *service) DescribeHost(ctx context.Context, req *host.DescribeHostRequest) (*host.Host, error) {
	query := sqlbuilder.NewQuery(queryHostSQL).GroupBy("r.id")
	cond, val := req.Where()
	querySQL, args := query.Where(cond, val).BuildQuery()
	s.log.Debugf("sql: %s", querySQL)

	queryStmt, err := s.db.PrepareContext(ctx, querySQL)
	if err != nil {
		return nil, exception.NewInternalServerError("prepare describe host error, %s", err.Error())
	}
	defer queryStmt.Close()

	ins := host.NewDefaultHost()
	var (
		publicIPList, privateIPList, keyPairNameList, securityGroupsList string
	)
	base := ins.Base
	info := ins.Information
	desc := ins.Describe
	err = queryStmt.QueryRowContext(ctx, args...).Scan(
		&base.Id, &base.ResourceType, &base.Vendor, &base.Region, &base.Zone, &base.CreateAt, &info.ExpireAt,
		&info.Category, &info.Type, &info.Name, &info.Description,
		&info.Status, &info.UpdateAt, &base.SyncAt, &info.SyncAccount,
		&publicIPList, &privateIPList, &info.PayType, &base.DescribeHash, &base.ResourceHash,
		&base.SecretId, &base.Domain, &base.Namespace, &base.Env, &base.UsageMode, &base.Id,
		&desc.Cpu, &desc.Memory, &desc.GpuAmount, &desc.GpuSpec, &desc.OsType, &desc.OsName,
		&desc.SerialNumber, &desc.ImageId, &desc.InternetMaxBandwidthOut, &desc.InternetMaxBandwidthIn,
		&keyPairNameList, &securityGroupsList,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, exception.NewNotFound("%#v not found", req)
		}
		return nil, exception.NewInternalServerError("describe host error, %s", err.Error())
	}

	info.LoadPrivateIPString(privateIPList)
	info.LoadPublicIPString(publicIPList)
	desc.LoadKeyPairNameString(keyPairNameList)
	desc.LoadSecurityGroupsString(securityGroupsList)

	return ins, nil
}

func (s *service) UpdateHost(ctx context.Context, req *host.UpdateHostRequest) (*host.Host, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateHost not implemented")
}

func (s *service) ReleaseHost(ctx context.Context, req *host.ReleaseHostRequest) (*host.Host, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReleaseHost not implemented")
}
