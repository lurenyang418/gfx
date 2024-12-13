package genlogic

import (
	"log"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/text/gregex"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gtag"
	"github.com/lurenyang418/gfx/cmd/gfx/internal/consts"
)

const (
	CGenLogicConfig         = `gfcli.gen.logic`
	CGenLogicUsage          = `gfx gen logic [OPTION]`
	CGenLogicBrief          = `gen logic template for startup`
	CGenLogicEg             = `gfx gen logic`
	CGenLogicBriefDstFolder = `destination folder path storing automatically generated go files. default: internal/logic`
	CGenLogicServiceName    = `destination service name`
)

func init() {
	gtag.Sets(g.MapStrStr{
		`CGenLogicConfig`:         CGenLogicConfig,
		`CGenLogicUsage`:          CGenLogicUsage,
		`CGenLogicBrief`:          CGenLogicBrief,
		`CGenLogicEg`:             CGenLogicEg,
		`CGenLogicBriefDstFolder`: CGenLogicBriefDstFolder,
		`CGenLogicServiceName`:    CGenLogicServiceName,
	})
}

type (
	CGenLogic      struct{}
	CGenLogicInput struct {
		g.Meta      `name:"logic" config:"{CGenLogicConfig}" usage:"{CGenLogicUsage}" brief:"{CGenLogicBrief}" eg:"{CGenLogicEg}"`
		DstFolder   string `short:"d" name:"dstFolder" brief:"{CGenLogicBriefDstFolder}" d:"internal/logic"`
		ServiceName string `short:"s" name:"serviceName" brief:"{CGenLogicServiceName}"`
	}
	CGenLogicOutput struct{}
)

func (c *CGenLogic) Logic(ctx g.Ctx, in CGenLogicInput) (out *CGenLogicOutput, err error) {
	pwd := gfile.Pwd()
	goModPath := gfile.Join(pwd, "go.mod")
	importPath := ""
	if gfile.Exists(goModPath) {
		match, _ := gregex.MatchString(`^module\s+(.+)\s*`, gfile.GetContents(goModPath))
		importPath = gstr.Trim(match[1]) + "/internal/service"
	} else {
		return nil, gerror.New("go.mod file not found")
	}

	dstLogicFolderPath := gfile.Join(in.DstFolder, in.ServiceName)

	serviceFilePath := gfile.Join(dstLogicFolderPath, in.ServiceName+".go")
	servicePlaceholderPath := gfile.Join(dstLogicFolderPath, in.ServiceName+"_get.go")

	if !gfile.Exists(serviceFilePath) {
		content := gstr.ReplaceByMap(consts.TemplateGenLogicNew, g.MapStrStr{
			"{PackageName}": in.ServiceName,
			"{Service}":     gstr.CaseCamel(in.ServiceName),

			"{ImportPath}": importPath,
		})
		if err = gfile.PutContents(serviceFilePath, gstr.TrimLeft(content)); err != nil {
			return nil, err
		}
	}

	if !gfile.Exists(servicePlaceholderPath) {
		content := gstr.ReplaceByMap(consts.TemplateGenLogicNewPlaceholder, g.MapStrStr{
			"{PackageName}": in.ServiceName,
			"{Service}":     gstr.CaseCamel(in.ServiceName),
		})
		if err = gfile.PutContents(servicePlaceholderPath, gstr.TrimLeft(content)); err != nil {
			return nil, err
		}
	}

	log.Println(`done!`)

	return nil, nil

}
