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
func (s *s{Service}) GetBy{Service}Id(ctx context.Context, {Service}Id int) (any, error){
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
`
