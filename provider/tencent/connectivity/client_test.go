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
	id := "xxx"
	key := "xxx"
	region := "xxx"
	client = NewTencentCloudClient(id, key, region)
	//if err != nil {
	//	panic(err)
	//}
}
