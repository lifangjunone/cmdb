package host

import (
	"context"
	"github.com/golang/protobuf/proto"
	"github.com/infraboard/mcube/http/request"
	"github.com/lifangjunone/cmdb/apps/resource"
	"github.com/lifangjunone/cmdb/utils"
	"net/http"
	"strings"
	"time"
)

const (
	AppName = "host"
)

func NewDescribeHostRequestWithID(id string) *DescribeHostRequest {
	return &DescribeHostRequest{
		DescribeBy: DescribeBy_HOST_ID,
		Value:      id,
	}
}

func NewUpdateHostDataByIns(ins *Host) *UpdateHostData {
	return &UpdateHostData{
		Information: ins.Information,
		Describe:    ins.Describe,
	}
}

func (x *Host) GenHash() error {
	// hash resource
	x.Base.ResourceHash = x.Information.Hash()
	// hash describe
	x.Base.DescribeHash = utils.Hash(x.Describe)
	return nil
}

func (x *Describe) KeyPairNameToString() string {
	return strings.Join(x.KeyPairName, ",")
}

func (x *Describe) SecurityGroupsToString() string {
	return strings.Join(x.SecurityGroups, ",")
}

func (h *Host) Put(req *UpdateHostData) {
	oldRH, oldDH := h.Base.ResourceHash, h.Base.DescribeHash

	h.Information = req.Information
	h.Describe = req.Describe
	h.Information.UpdateAt = time.Now().UnixMilli()
	h.GenHash()

	if h.Base.ResourceHash != oldRH {
		h.Base.ResourceHashChanged = true
	}
	if h.Base.DescribeHash != oldDH {
		h.Base.DescribeHashChanged = true
	}
}

func NewHostSet() *HostSet {
	return &HostSet{
		Items: []*Host{},
	}
}

func (x *HostSet) Add(item *Host) {
	x.Items = append(x.Items, item)
	return
}

func (x *HostSet) Length() int64 {
	return int64(len(x.Items))
}

func NewDefaultHost() *Host {
	return &Host{
		Base: &resource.Base{
			ResourceType: resource.Type_HOST,
		},
		Information: &resource.Information{},
		Describe:    &Describe{},
	}
}

func (x *HostSet) Clone() *HostSet {
	return proto.Clone(x).(*HostSet)
}

func (x *Describe) LoadKeyPairNameString(str string) {
	if str != "" {
		x.KeyPairName = strings.Split(str, ",")
	}

}

func (x *Describe) LoadSecurityGroupsString(str string) {
	if str != "" {
		x.SecurityGroups = strings.Split(str, ",")
	}
}

func NewQueryHostRequestFromHttp(r *http.Request) *QueryHostRequest {
	qs := r.URL.Query()
	page := request.NewPageRequestFromHTTP(r)
	kw := qs.Get("keywords")

	return &QueryHostRequest{
		Page:     page,
		Keywords: kw,
	}
}

func (x *DescribeHostRequest) Where() (string, interface{}) {
	switch x.DescribeBy {
	default:
		return "r.id = ?", x.Value
	}
}

type Pager interface {
	Next() bool
	Scan(ctx context.Context, set *HostSet) error
	SetPageSize(ps int64)
}
