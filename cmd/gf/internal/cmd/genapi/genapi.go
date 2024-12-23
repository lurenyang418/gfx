package genapi

import (
	"log"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gtag"
	"github.com/lurenyang418/gfx/cmd/gfx/internal/consts"
)

const (
	CGenApiConfig         = `gfcli.gen.api`
	CGenApiUsage          = `gfx gen api [OPTION]`
	CGenApiBrief          = `gen api template for startup`
	CGenApiEg             = `gfx gen api`
	CGenApiBriefDstFolder = `destination folder path storing automatically generated go files. default: api`
	CGenApiName           = `destination api name. `
	CGenApiVersion        = `destination api version. default: v1`
)

func init() {
	gtag.Sets(g.MapStrStr{
		`CGenApiConfig`:         CGenApiConfig,
		`CGenApiUsage`:          CGenApiUsage,
		`CGenApiBrief`:          CGenApiBrief,
		`CGenApiEg`:             CGenApiEg,
		`CGenApiBriefDstFolder`: CGenApiBriefDstFolder,
		`CGenApiName`:           CGenApiName,
		`CGenApiVersion`:        CGenApiVersion,
	})
}

type (
	CGenApi      struct{}
	CGenApiInput struct {
		g.Meta     `name:"api" config:"{CGenApiConfig}" usage:"{CGenApiUsage}" brief:"{CGenApiBrief}" eg:"{CGenApiEg}"`
		DstFolder  string `short:"d" name:"dstFolder" brief:"{CGenApiBriefDstFolder}" d:"api"`
		ApiName    string `short:"s" name:"apiName" brief:"{CGenApiName}"`
		ApiVersion string `short:"v" name:"apiVersion" brief:"{CGenApiVersion}" d:"v1"`
	}
	CGenApiOutput struct{}
)

func (c *CGenApi) Api(ctx g.Ctx, in CGenApiInput) (out *CGenApiOutput, err error) {
	dstApiFolderPath := gfile.Join(in.DstFolder, in.ApiName, in.ApiVersion)

	apiFilePath := gfile.Join(dstApiFolderPath, in.ApiName+".go")
	cruds := []map[string]string{
		{"Add": "post"},
		{"Get": "get"},
		{"List": "get"},
		{"Update": "put"},
		{"Delete": "delete"},
	}

	if !gfile.Exists(apiFilePath) {
		content := gstr.ReplaceByMap(consts.TemplateGenApiNew, g.MapStrStr{
			"{ApiVersion}": in.ApiVersion,
		})
		if err = gfile.PutContentsAppend(apiFilePath, gstr.TrimLeft(content)); err != nil {
			return nil, err
		}
	}
	for _, vMap := range cruds {
		for k, v := range vMap {
			path := in.ApiName
			if k == "List" {
				path = path + "s"
			}
			content := gstr.ReplaceByMap(consts.TemplateGenApiNewPlaceholder, g.MapStrStr{
				"{Service}": gstr.CaseCamel(in.ApiName),
				"{Path}":    path,
				"{Method}":  k,
				"{method}":  v,
			})
			if err = gfile.PutContentsAppend(apiFilePath, gstr.TrimLeft(content)); err != nil {
				return nil, err
			}
		}
	}

	log.Println(`done!`)

	return nil, nil

}
