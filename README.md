## 用于生成 gf 的 logic

> 就懒那么一点点

```shell
go install github.com/lurenyang418/gfx/cmd/gfx

gfx gen logic -s user

# 目录如下
# internal/logic/user
# ├── user.go
# └── user_get.go

```

1. user.go

```go
package user

import "<mod>/internal/service"

type sUser struct{}

func init() {
	service.RegisterUser(New())
}

func New() *sUser {
	return &sUser{}
}

```

2. user_get.go

```go
package user

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

)

// this is a handy copy
func (s *sUser) GetByUserId(ctx context.Context, UserId int) (any, error){
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

```