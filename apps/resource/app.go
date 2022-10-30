package resource

import "strings"

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

func NewDefaultTag() *Tag {
	return &Tag{
		Type:   TagType_USER,
		Weight: 1,
	}
}
