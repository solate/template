package virhis

import (
	"github.com/solate/generation"
	"github.com/solate/generation/config"
	"github.com/solate/generation/utils"
)

func SetConfig() (err error) {

	list, err := generation.ParseMarkdownByFile("./api.md")
	if err != nil {
		return
	}

	//service 模板
	templateService, err := utils.ReadFile("./template/service.tpl")
	if err != nil {
		return
	}

	//module 模板
	templateModule, err := utils.ReadFile("./template/module.tpl")
	if err != nil {
		return
	}

	//获得service 参数
	structName, params, err := GetServiceParmas(list)
	if err != nil {
		return
	}
	//获得modules 参数
	modulesParams, err := GetModuleParams(list)
	if err != nil {
		return
	}

	cfg := &config.Setting{
		ServiceTemplate:       templateService,
		ServiceTemplateParams: params,
		ServiceExport:         "./export/service/",
		ServiceFileName:       structName,
		ModuleTemplate:        templateModule,
		ModuleExport:          "./export/module/",
		ModuleTemplateParams:  modulesParams,
	}

	config.SetConfig(cfg)

	return
}
