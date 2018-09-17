package virhis

import (
	"fmt"
	"github.com/solate/generation"
	"github.com/solate/generation/model"
	"github.com/solate/generation/utils"
	"strings"
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
				methodName,
				methodReq,
				methodReply,
				//" //TODO 代码",
				GetCodeListBody(fileName, md.Request, md.Response),
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
				methodName,
				methodReq,
				methodReply,
				//" //TODO 代码",
				GetCodeBody(fileName, md.Request, md.Response),
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

func GetCodeListBody(fileName string, reqs []model.MarkDownReqTable, replys []model.MarkDownReplyTable) string {

	methodReq := fileName + "Req"
	methodReply := fileName + "Reply"

	file, err := utils.ReadFile("./template/module_body_list.tpl")
	if err != nil {
		fmt.Println(err)
	}

	var tmp []string
	for _, v := range reqs {
		var str = fmt.Sprintf("%s : req.%s,", v.Name, v.Name)
		tmp = append(tmp, str)
	}
	reqStr := strings.Join(tmp, "\n")

	var replyTmp []string
	for _, v := range replys {
		var str = fmt.Sprintf("%s: v.%s,", v.Name, v.Name)
		replyTmp = append(replyTmp, str)
	}
	replyStr := strings.Join(replyTmp, "\n")

	return fmt.Sprintf(file,
		utils.Lcfirst(methodReq),
		methodReq,
		reqStr,
		utils.Lcfirst(methodReply),
		methodReply,
		fileName,
		utils.Lcfirst(methodReq),
		utils.Lcfirst(methodReply),
		fileName,
		fileName,
		utils.Lcfirst(methodReply),
		utils.Lcfirst(methodReply),
		fileName,
		replyStr,
	)
}

func GetCodeBody(fileName string, reqs []model.MarkDownReqTable, replys []model.MarkDownReplyTable) string {
	methodReq := fileName + "Req"
	methodReply := fileName + "Reply"

	file, err := utils.ReadFile("./template/module_body.tpl")
	if err != nil {
		fmt.Println(err)
	}

	var reqTmp []string
	for _, v := range reqs {
		var str = fmt.Sprintf("%s : req.%s,", v.Name, v.Name)
		reqTmp = append(reqTmp, str)
	}
	reqStr := strings.Join(reqTmp, "\n")

	var tmp []string
	for _, v := range replys {
		var str = fmt.Sprintf("reply.%s = %s.%s", v.Name, utils.Lcfirst(methodReply), v.Name)
		tmp = append(tmp, str)
	}
	replyStr := strings.Join(tmp, "\n")

	return fmt.Sprintf(file,
		utils.Lcfirst(methodReq),
		methodReq,
		reqStr,
		utils.Lcfirst(methodReply),
		methodReply,
		fileName,
		utils.Lcfirst(methodReq),
		utils.Lcfirst(methodReply),
		fileName,
		replyStr,
	)
}
