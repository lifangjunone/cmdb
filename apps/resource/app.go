package resource

import (
	"fmt"
	"github.com/infraboard/mcube/http/request"
	"github.com/lifangjunone/cmdb/utils"
	"net/http"
	"strings"
)

const (
	AppName = "resource"
)

const (
	OperatorEqual        = "="
	OperatorNotEqual     = "!="
	OperatorLikeEqual    = "=~"
	OperatorNotLikeEqual = "!~"
)

// HasTag 判断查询条件是否包含tag标签
func (x *SearchRequest) HasTag() bool {
	return len(x.Tags) > 0
}

// RelationShip 多个值比较的关系说明:
func (x *TagSelector) RelationShip() string {
	switch x.Operator {
	case OperatorEqual, OperatorLikeEqual:
		return " OR "
	case OperatorNotEqual, OperatorNotLikeEqual:
		return " AND "
	default:
		return " OR "
	}
}

// NewResourceSet ResourceSet对象构建方法
func NewResourceSet() *ResourceSet {
	return &ResourceSet{
		Total: 0,
		Items: []*Resource{},
	}
}

// NewDefaultResource Resource对象构建方法
func NewDefaultResource() *Resource {
	return &Resource{
		Base:        &Base{},
		Information: &Information{},
	}
}

// LoadPrivateIPString 将逗号拼接的字符串转换为列表
func (x *Information) LoadPrivateIPString(s string) {
	if s != "" {
		x.PrivateIp = strings.Split(s, ",")
	}
}

// LoadPublicIPString 将逗号拼接的字符串转换为列表
func (x *Information) LoadPublicIPString(s string) {
	if s != "" {
		x.PublicIp = strings.Split(s, ",")
	}
}

func (x *ResourceSet) Add(item *Resource) {
	x.Items = append(x.Items, item)
}

func (x *ResourceSet) ResourceIds() (ids []string) {
	for i := range x.Items {
		ids = append(ids, x.Items[i].Base.Id)
	}

	return
}

func (x *ResourceSet) UpdateTag(tags []*Tag) {
	for i := range tags {
		for j := range x.Items {
			if x.Items[j].Base.Id == tags[i].ResourceId {
				x.Items[j].Information.AddTag(tags[i])
			}
		}
	}
}

func (x *Information) AddTag(tag *Tag) {
	x.Tags = append(x.Tags, tag)
}

func NewDefaultTag() *Tag {
	return &Tag{
		Type:   TagType_USER,
		Weight: 1,
	}
}

func ParExpr(str string) (*TagSelector, error) {
	var (
		op = ""
		kv = []string{}
	)

	// app=~v1,v2,v3
	if strings.Contains(str, OperatorLikeEqual) {
		op = "LIKE"
		kv = strings.Split(str, OperatorLikeEqual)
	} else if strings.Contains(str, OperatorNotLikeEqual) {
		op = "NOT LIKE"
		kv = strings.Split(str, OperatorNotLikeEqual)
	} else if strings.Contains(str, OperatorNotEqual) {
		op = "!="
		kv = strings.Split(str, OperatorNotEqual)
	} else if strings.Contains(str, OperatorEqual) {
		op = "="
		kv = strings.Split(str, OperatorEqual)
	} else {
		return nil, fmt.Errorf("no support operator [=, =~, !=, !~]")
	}

	if len(kv) != 2 {
		return nil, fmt.Errorf("key,value format error, requred key=value")
	}

	selector := &TagSelector{
		Key:      kv[0],
		Operator: op,
		Values:   []string{},
	}

	// 如果Value等于*表示只匹配key
	if kv[1] != "*" {
		selector.Values = strings.Split(kv[1], ",")
	}

	return selector, nil
}

func NewTagsFromString(tagStr string) (tags []*TagSelector, err error) {
	if tagStr == "" {
		return
	}

	items := strings.Split(tagStr, "&")
	for _, v := range items {
		// key1=v1,v2,v3 --> TagSelector
		t, err := ParExpr(v)
		if err != nil {
			return nil, err
		}
		tags = append(tags, t)
	}

	return
}

func (x *SearchRequest) AddTag(t ...*TagSelector) {
	x.Tags = append(x.Tags, t...)
}

func NewSearchRequestFromHTTP(r *http.Request) (*SearchRequest, error) {
	qs := r.URL.Query()
	req := &SearchRequest{
		Page:        request.NewPageRequestFromHTTP(r),
		Keywords:    qs.Get("keywords"),
		ExactMatch:  qs.Get("exact_match") == "true",
		Domain:      qs.Get("domain"),
		Namespace:   qs.Get("namespace"),
		Env:         qs.Get("env"),
		Status:      qs.Get("status"),
		SyncAccount: qs.Get("sync_account"),
		WithTags:    qs.Get("with_tags") == "true",
		Tags:        []*TagSelector{},
	}

	umStr := qs.Get("usage_mode")
	if umStr != "" {
		mode, err := ParseUsageModeFromString(umStr)
		if err != nil {
			return nil, err
		}
		req.UsageMode = &mode
	}

	rtStr := qs.Get("resource_type")
	if rtStr != "" {
		rt, err := ParseTypeFromString(rtStr)
		if err != nil {
			return nil, err
		}
		req.Type = &rt
	}

	// 单独处理Tag参数 app~=app1,app2,app3 --> TagSelector ---> req
	tgStr := qs.Get("tag")
	if tgStr != "" {
		tg, err := NewTagsFromString(tgStr)
		if err != nil {
			return nil, err
		}
		req.AddTag(tg...)
	}

	return req, nil
}

func (x *Information) Hash() string {
	return utils.Hash(x)
}

func (i *Information) PrivateIPToString() string {
	return strings.Join(i.PrivateIp, ",")
}

func (i *Information) PublicIPToString() string {
	return strings.Join(i.PublicIp, ",")
}

func NewThirdTag(key, value string) *Tag {
	return &Tag{
		Type:   TagType_THIRD,
		Key:    key,
		Value:  value,
		Weight: 1,
	}
}
