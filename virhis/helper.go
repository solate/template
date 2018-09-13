package virhis

import (
	"fmt"
	"github.com/solate/generation"
	"github.com/solate/generation/utils"
	"strings"
)

func GenerateCode(str string, note string) string {

	str, err := utils.ReadFile("./template/virhis.txt")
	if err != nil {
		panic(err)
	}

	methodName := str
	methodReq := str + "Req"
	methodReply := str + "Reply"
	methodURL := "URL" + str
	reqStr := genReq(methodReq)

	return fmt.Sprintf(str, note, methodName, methodReq, methodReply, reqStr, methodURL, methodURL, methodReply)
}

func genReq(sub string) string {
	//读取文件
	str, err := utils.ReadFile("./source/virhis.go")
	if err != nil {
		panic(err)
	}

	list := generation.GenerationReq(str, sub, "data.Add(\"%s\", %s)")
	return strings.Join(list, "\n")
}
