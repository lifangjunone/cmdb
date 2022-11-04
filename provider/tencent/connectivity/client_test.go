package connectivity

import (
	"testing"
)

func TestTencentCloudClient(t *testing.T) {
	conn := C()
	if err := conn.Check(); err != nil {
		t.Fatal(err)
	}
	t.Log(conn.AccountID())

}

func init() {
	//err := LoadClientFromEnv()
	id := "AKIDJXiojFcmkV4TggT4hHi3GIF5ri0HSdgY"
	key := "qY7UiquumV4Vs6B8uZ6Vkq49A4qV2Oc2"
	region := "ap-beijing"
	client = NewTencentCloudClient(id, key, region)
	//if err != nil {
	//	panic(err)
	//}
}
