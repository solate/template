package spro

import (
	"github.com/gin-gonic/gin"
	"gitlab.dev.okapp.cc/med-his-group/medhis-platform-service/components/log"
	"gitlab.dev.okapp.cc/med-his-group/medhis-platform-service/components/moduleman"
	"gitlab.dev.okapp.cc/med-his-group/medhis-platform-service/components/server"
	"gitlab.dev.okapp.cc/med-his-group/medhis-platform-service/modules/pro"
	"gitlab.dev.okapp.cc/med-his-group/medhis-platform-service/std"
	"net/http"
)

type Pro struct {
	GetOutpatientReportList func(*gin.Context) `path:"/outpatient/reports"` //获得门诊检查检验报告，并按时间排序
}

func (Pro) Name() string {
	return "pro"
}

func BuildPro() Pro {
	return Pro{
		GetOutpatientReportList: GetOutpatientReportList,
	}
}

//获得门诊检查检验报告，并按时间排序
func GetOutpatientReportList(c *gin.Context) {

	reply := new(pro.GetOutpatientReportListReply)

	req := new(pro.GetOutpatientReportListReq)
	if err := server.BindRequest(c, req); err != nil {
		c.JSON(http.StatusOK, std.BuildAdminResponse(false, std.AdminCodeParamError))
		return
	}
	if err := moduleman.Call(c, moduleman.Request{
		Module: "pro.Pro",
		Func:   "GetOutpatientReportList",
		Params: req,
	}, reply); err != nil {
		log.Error(err)
		c.JSON(http.StatusOK, std.BuildApiResponse(false, err.Code, err.Msg))
		return
	}
	c.JSON(http.StatusOK, std.BuildApiResponse(reply))

}
