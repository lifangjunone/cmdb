package secret

import (
	"encoding/base64"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/infraboard/mcube/crypto/cbc"
	"github.com/infraboard/mcube/http/request"
	"github.com/lifangjunone/cmdb/conf"
	"github.com/rs/xid"
	"net/http"
	"strings"
	"time"
)

const (
	AppName = "secret"
)

var (
	validate = validator.New()
)

func (x *CreateSecretRequest) Validate() error {
	if len(x.AllowRegions) == 0 {
		return fmt.Errorf("required less one allow_regions")
	}
	return validate.Struct(x)
}

func NewSecret(req *CreateSecretRequest) (*Secret, error) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}
	return &Secret{
		Id:       xid.New().String(),
		CreateAt: time.Now().UnixMilli(),
		Data:     req,
	}, nil
}

func NewSecretSet() (req *SecretSet) {
	return &SecretSet{
		Items: []*Secret{},
	}
}

func NewDefaultSecret() *Secret {
	return &Secret{
		Data: &CreateSecretRequest{
			RequestRate: 5,
		},
	}
}

func NewDescribeSecretRequest(id string) *DescribeSecretRequest {
	return &DescribeSecretRequest{
		Id: id,
	}
}

func NewCreateSecretRequest() *CreateSecretRequest {
	return &CreateSecretRequest{
		RequestRate: 5,
	}
}

func NewQuerySecretRequestFromHTTP(r *http.Request) *QuerySecretRequest {
	qs := r.URL.Query()

	return &QuerySecretRequest{
		Page:     request.NewPageRequestFromHTTP(r),
		Keywords: qs.Get("keywords"),
	}
}

func NewDeleteSecretRequestWithID(id string) *DeleteSecretRequest {
	return &DeleteSecretRequest{
		Id: id,
	}
}

func NewQuerySecretRequest() *QuerySecretRequest {
	return &QuerySecretRequest{
		Page:     request.NewDefaultPageRequest(),
		Keywords: "",
	}
}

func (x *CreateSecretRequest) EncryptAPISecret(key string) error {
	// 判断文本是否已经加密
	if strings.HasPrefix(x.ApiSecret, conf.CIPHER_TEXT_PREFIX) {
		return fmt.Errorf("text has ciphered")
	}

	cipherText, err := cbc.Encrypt([]byte(x.ApiSecret), []byte(key))
	if err != nil {
		return err
	}

	base64Str := base64.StdEncoding.EncodeToString(cipherText)
	x.ApiSecret = fmt.Sprintf("%s%s", conf.CIPHER_TEXT_PREFIX, base64Str)
	return nil
}

func (x *CreateSecretRequest) AllowRegionString() string {
	return strings.Join(x.AllowRegions, ",")
}

func (x *CreateSecretRequest) LoadAllowRegionFromString(regions string) {
	if regions != "" {
		x.AllowRegions = strings.Split(regions, ",")
	}
}

func (x *CreateSecretRequest) Desense() {
	if x.ApiSecret != "" {
		x.ApiSecret = "******"
	}
}

func (x *SecretSet) Add(ins *Secret) {
	x.Items = append(x.Items, ins)
}
