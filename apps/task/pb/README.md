## generate proto file 
### command description

### code struct

```text
|-cmdb
| |-app                      
|   |-task
|     |-pb
|       |-task.proto
|   |-resource
|     |-pb
|       |-resource.proto   
| |-common
|   |-pb
|     |-page
|       |-page.proto 
           
```
### proto file import example
```protobuf
syntax = "proto3";

package cmdb.task;
option go_package = "github.com/lifangjunone/cmdb/apps/task";

import "apps/resource/pb/resource.proto";
import "common/pb/page/page.proto";
```

### proto file generate command
```shell
protoc -I=. -I=./common/pb   --go_out=. --go_opt=module="github.com/lifangjunone/cmdb"  --go-grpc_out=. --go-grpc_opt=module="github.com/lifangjunone/cmdb"  apps/*/pb/*proto  common/pb/*/*proto
```

