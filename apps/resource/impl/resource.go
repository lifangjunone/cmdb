package impl

import (
	"context"
	"fmt"
	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/sqlbuilder"
	"github.com/lifangjunone/cmdb/apps/resource"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strings"
)

func (s *service) Search(ctx context.Context, req *resource.SearchRequest) (*resource.ResourceSet, error) {
	// SQL 是一个模板,  到底应该使用左连接还是右连接，取决于是否需要关联Tag表
	// LEFT JOIN 是先扫描左表, RIGHT JOIN先扫描右表, 当有Tag过滤，需要关联右表, 可以以右表为准
	// 如果 扫描Tag表的成本比扫描Resource成本低，我们就使用RIGHT JOIN
	join := "LEFT"
	if req.HasTag() {
		join = "RIGHT"
	}

	// 构建过滤条件
	builder := sqlbuilder.NewQuery(fmt.Sprintf(sqlQueryResource, join))
	s.buildQuery(builder, req)

	// =========
	// 计数统计: COUNT语句
	// =========
	set := resource.NewResourceSet()

	// 获取total SELECT COUNT(*) FROM t Where ....
	countSQL, args := builder.BuildFromNewBase(fmt.Sprintf(sqlCountResource, join))
	countStmt, err := s.db.Prepare(countSQL)
	if err != nil {
		s.log.Debugf("count sql, %s, %v", countSQL, args)
		return nil, exception.NewInternalServerError("prepare count sql error, %s", err)
	}

	defer countStmt.Close()
	err = countStmt.QueryRow(args...).Scan(&set.Total)
	if err != nil {
		return nil, exception.NewInternalServerError("scan count value error, %s", err)
	}

	// =========
	// 查询分页数据: SELECT 语句
	// =========

	// tag查询时，以tag时间排序, 如果没有Tag就以资源的创建时间进行排序
	// 比如你添加资源, 最后添加的资源，最先被看到, 就是一个堆, Stack
	if req.HasTag() {
		builder.Order("t.create_at").Desc()
	} else {
		builder.Order("r.create_at").Desc()
	}

	// 获取分页数据
	querySQL, args := builder.
		GroupBy("r.id").
		Limit(req.Page.ComputeOffset(), uint(req.Page.PageSize)).
		BuildQuery()
	s.log.Debugf("sql: %s, args: %v", querySQL, args)

	queryStmt, err := s.db.PrepareContext(ctx, querySQL)
	if err != nil {
		return nil, exception.NewInternalServerError("prepare query resource error, %s", err.Error())
	}
	defer queryStmt.Close()

	rows, err := queryStmt.QueryContext(ctx, args...)
	if err != nil {
		return nil, exception.NewInternalServerError(err.Error())
	}
	defer rows.Close()

	var (
		publicIPList, privateIPList string
	)

	for rows.Next() {
		ins := resource.NewDefaultResource()
		base := ins.Base
		info := ins.Information
		err := rows.Scan(
			&base.Id, &base.ResourceType, &base.Vendor, &base.Region, &base.Zone, &base.CreateAt, &info.ExpireAt,
			&info.Category, &info.Type, &info.Name, &info.Description,
			&info.Status, &info.UpdateAt, &base.SyncAt, &info.SyncAccount,
			&publicIPList, &privateIPList, &info.PayType, &base.DescribeHash, &base.ResourceHash,
			&base.SecretId, &base.Domain, &base.Namespace, &base.Env, &base.UsageMode,
		)
		if err != nil {
			return nil, exception.NewInternalServerError("query resource error, %s", err.Error())
		}
		// 存入数据库的是一个列表, 格式: 10.10.1.1,10.10.2.2,.....,
		// 因此我们从数据库取出该数据, 对格式进行特殊处理
		info.LoadPrivateIPString(privateIPList)
		info.LoadPublicIPString(publicIPList)
		set.Add(ins)
	}

	// 补充资源的标签
	// 为什么 不在上个SQL，直接把Tag查出来喃?
	if req.WithTags {
		tags, err := QueryTag(ctx, s.db, set.ResourceIds())
		if err != nil {
			return nil, err
		}
		set.UpdateTag(tags)
	}
	return set, nil
}

func (s *service) QueryTag(ctx context.Context, req *resource.QueryTagRequest) (*resource.TagSet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryTag not implemented")
}

func (s *service) UpdateTag(ctx context.Context, req *resource.UpdateTagRequest) (*resource.Resource, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTag not implemented")
}

func (s *service) buildQuery(builder *sqlbuilder.Builder, req *resource.SearchRequest) {
	// 参数里面有模糊匹配与关键字匹配
	if req.Keywords != "" {
		if req.ExactMatch {
			// 精确匹配
			builder.Where("r.name = ? OR r.id = ? OR r.private_ip = ? OR r.public_ip = ?",
				req.Keywords,
				req.Keywords,
				req.Keywords,
				req.Keywords,
			)
		} else {
			// 模糊匹配
			builder.Where("r.name LIKE ? OR r.id = ? OR r.private_ip LIKE ? OR r.public_ip LIKE ?",
				"%"+req.Keywords+"%",
				"%"+req.Keywords+"%",
				req.Keywords+"%",
				req.Keywords+"%",
			)
		}
	}

	// 按照资源属性过滤
	if req.Domain != "" {
		builder.Where("r.domain = ?", req.Domain)
	}
	if req.Namespace != "" {
		builder.Where("r.namespace = ?", req.Namespace)
	}
	if req.Env != "" {
		builder.Where("r.env = ?", req.Env)
	}
	if req.UsageMode != nil {
		builder.Where("r.usage_mode = ?", req.UsageMode)
	}
	if req.Vendor != nil {
		builder.Where("r.vendor = ?", req.Vendor)
	}
	if req.SyncAccount != "" {
		builder.Where("r.sync_accout = ?", req.SyncAccount)
	}
	if req.Type != nil {
		builder.Where("r.resource_type = ?", req.Type)
	}
	if req.Status != "" {
		builder.Where("r.status = ?", req.Status)
	}

	// 如何通过Tag匹配资源, 通过tag key 和 tag value 进行联表查询 配上where条件
	// 我们允许输入多个Tag来对资源进行解索, 多个Tag之间的关系, 到底是AND OR  app=v1, product=p2
	// 我们实现的策略:  基于AND
	for i := range req.Tags {
		selector := req.Tags[i]

		// tag:   =v1, 做为Tag查询，Tag的key是必须的
		if selector.Key == "" {
			continue
		}

		// 添加Key过滤条件,  tag_key="xxxx" .*, 定制化 key如何统配
		builder.Where("t.t_key LIKE ?", strings.ReplaceAll(selector.Key, ".*", "%"))

		// 场景一: 定制Value如何统配, app=["app1", "app2", "app3"]
		// tag_value=? OR tag_value=?, 有几个Tag Value就需要构造结构Where OR条件
		// 创建二: app_count > 1

		// (tag_value LIKE ? OR tag_value LIKE ?)
		var condions []string
		var args []any
		for _, v := range selector.Values {
			// t.t_value [= != =~ !~] value
			condions = append(condions, fmt.Sprintf("t.t_value %s ?", selector.Operator))
			// 条件参数 args
			// args = append(args, v)

			// tag_value .* --> %, 做的特殊处理, 为了匹配正则里面的.*,
			// app=product1.*  --转换为--> app=prodcut1.%
			args = append(args, strings.ReplaceAll(v, ".*", "%"))
		}

		// tag的value是由多个条件组成的 app=~app1,app2, 根据表达式 [= != =~ !~], 来智能觉得value之间的关系
		if len(condions) > 0 {
			vwhere := fmt.Sprintf("( %s )", strings.Join(condions, selector.RelationShip()))
			builder.Where(vwhere, args...)
		}
	}
}
