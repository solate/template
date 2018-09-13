package virhis

import (
	"fmt"
	"github.com/solate/generation"
	"github.com/solate/generation/model"
	"github.com/solate/generation/utils"
	"strings"
)

//获取一个service 需要的所有参数, 以及文件取名字
func GetServiceParmas(mds []model.MarkDown) (structName string, params []interface{}, err error) {

	//service 例子
	example, err := utils.ReadFile("./source/service.go")
	if err != nil {
		return
	}

	packageName := utils.GetPackageName(example)
	importBody := utils.GetImportBody(example)
	structName = utils.Ucfirst(packageName)

	routerBody := generation.GetTemplateRouter(mds)
	methodBody := generation.GetTemplateMethodDeclaration(mds)

	funcList := make([]string, 0, len(mds))
	for _, v := range mds {
		funcList = append(funcList, GetServiceFunc(v))
	}

	funcString := strings.Join(funcList, "\n\n")

	params = []interface{}{
		packageName,
		importBody,
		//struct
		structName,
		routerBody,
		//name()
		structName,
		packageName,
		//build()
		structName,
		structName,
		structName,
		methodBody,
		funcString,
	}

	return

}

//生成一个方法
func GetServiceFunc(md model.MarkDown) (str string) {

	url := md.URL
	method := md.Method
	note := md.Note
	methodName := utils.GetUriMethodName(url, method)
	methodReq := methodName + "Req"
	methodReply := methodName + "Reply"

	params := []interface{}{
		note, //注释
		methodName,
		methodReq,
		methodReply,
		methodName,
	}

	str, err := generation.ExampleByFile("./template/service_func.tpl", params...)
	if err != nil {
		fmt.Println(err)
	}

	return

}
