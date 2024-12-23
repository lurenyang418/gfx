package consts

const TemplateGenApiNew = `
package {ApiVersion}

import (

	"github.com/gogf/gf/v2/frame/g"
)


`

const TemplateGenApiNewPlaceholder = `
type {Service}{Method}Req struct {
  	g.Meta ` + "`path:\"{Path}\" tags:\"{Service}\" method:\"{method}\" sm:\"TBD\"`" + `
  	// 数据 
}
type {Service}{Method}Res struct {
	// 数据
}

`
