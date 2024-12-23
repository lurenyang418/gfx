package consts

const TemplateGenLogicNew = `
package {PackageName}

import (
	"{ImportPath}"
)

type s{Service} struct{}

func init() {
	service.Register{Service}(New())
}

func New() *s{Service} {
	return &s{Service}{}
}
`

const TemplateGenLogicNewPlaceholder = `
package {PackageName}

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

)

// this is a handy copy
func (s *s{Service}) {Service}{Method}(ctx context.Context, req *model.{Service}{Method}Req ) (any, error){
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
`
