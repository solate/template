package virhis

import (
	"fmt"
	"github.com/solate/generation"
	"github.com/solate/generation/model"
	"github.com/solate/generation/utils"
)

func GetModuleParams(mds []model.MarkDown) (params map[string][]interface{}, err error) {

	//module 例子
	example, err := utils.ReadFile("./source/module.go")
	if err != nil {
		return
	}

	packageName := utils.GetPackageName(example)
	importBody := utils.GetImportBody(example)
	structName := utils.Ucfirst(packageName)

	params = make(map[string][]interface{})
	for _, md := range mds {

		url := md.URL
		method := md.Method
		note := md.Note
		fileName := utils.GetUriMethodName(url, "") //无前缀
		methodName := utils.GetUriMethodName(url, method)
		methodReq := methodName + "Req"
		methodReply := methodName + "Reply"

		var param []interface{}
		if md.ResponseType {
			param = []interface{}{
				packageName,
				importBody,
				methodReq,
				generation.GetTemplateReqBody(md.Request),
				methodReply,
				GetListBody(fileName, generation.GetTemplateReplyBody(md.Response)),
				note,
				structName,
				methodReq,
				methodReply,
				GetCodeBody(fileName, methodReq, methodReply),
			}

		} else {
			param = []interface{}{
				packageName,
				importBody,
				methodReq,
				generation.GetTemplateReqBody(md.Request),
				methodReply,
				generation.GetTemplateReplyBody(md.Response),
				note,
				structName,
				methodReq,
				methodReply,
				GetCodeBody(fileName, methodReq, methodReply),
			}
		}

		params[fileName] = param

	}

	return
}

const TemplateList = `
	List []%s
}

type %s struct{
    %s
`

func GetListBody(structName string, body string) string {
	return fmt.Sprintf(TemplateList, structName, structName, body)
}

func GetCodeBody(fileName, methodReq, methodReply string) string {

	file, err := utils.ReadFile("./template/module_body.tpl")
	if err != nil {
		fmt.Println(err)
	}

	return fmt.Sprintf(file,
		utils.Lcfirst(methodReq),
		methodReq,
		utils.Lcfirst(methodReply),
		methodReply,
		fileName,
		utils.Lcfirst(methodReq),
		utils.Lcfirst(methodReply),
		fileName,
		utils.Lcfirst(methodReply),
		utils.Lcfirst(methodReply),
		utils.Lcfirst(methodReply),
		fileName,
	)
}
