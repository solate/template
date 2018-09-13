package pro

import (
	"gitlab.dev.okapp.cc/med-his-group/medhis-platform-service/components/moduleman"
	"gitlab.dev.okapp.cc/med-his-group/medhis-platform-service/std"
	"gitlab.dev.okapp.cc/med-his-group/medhis-platform-service/virhis"
	"time"
)

//pro 版使用结构
type Pro struct{}

type GetOutpatientReportListReq struct {
	MedicalCard string `form:"medicalCard" json:"medicalCard" binding:"required"` //就诊卡/身份证/门诊号
	CardType    int    `form:"cardType" json:"cardType" binding:"required"`       //证件号类型 1 身份证 2，就诊卡，3 门诊号
	DateType    int    `form:"dateType" json:"dateType" binding:"required"`       //时间类型， 1. 全部 2 三个月 3. 一个月

}

type GetOutpatientReportListReply struct {
	List []OutpatientReportList
}
type OutpatientReportList struct {
	ReportNum   string `json:"reportNum"`   //报告单号
	PatientNum  string `json:"patientNum"`  //病人号
	PatientName string `json:"patientName"` //病人姓名
	PatientSex  string `json:"patientSex"`  //病人性别
	PatientAge  string `json:"patientAge"`  //病人年龄
}

//获得门诊检查检验报告，并按时间排序
func (*Pro) GetOutpatientReportList(c *moduleman.Context, req *GetOutpatientReportListReq, reply *GetOutpatientReportListReply) *moduleman.Error {

	hospitalId := c.HospitalId()

	his, err := NewHis(hospitalId)
	if err != nil {
		logger.Error("init his err, please check hospitalId", err)
		return moduleman.NewError(std.AdminCodeServerError, err.Error())
	}

	//构造需要传递的开始时间和结束时间
	var beginTime, endTime string
	switch req.DateType {
	case 1:
		beginTime = TimeFormart(AllBeginTime) //1990年
	case 2:
		beginTime = TimeFormart(time.Now().AddDate(0, -3, 0)) //三个月
	case 3:
		beginTime = TimeFormart(time.Now().AddDate(0, -1, 0)) //一个月
	}
	endTime = TimeFormart(time.Now())

	outpatientReq := &virhis.OutpatientInspectionReportListReq{
		MedicalCard: req.MedicalCard,
		CardType:    req.CardType,
		BeginTime:   beginTime,
		EndTime:     endTime,
	}
	outpatientReply := new(virhis.OutpatientInspectionReportListReply)
	err = his.OutpatientInspectionReportList(outpatientReq, outpatientReply)
	if err != nil {
		logger.Error("call OutpatientInspectionReportList err, please remote", err)
		moduleman.NewError(std.AdminCodeServerError, err.Error())
	}

	list := make([]OutpatientReportList, 0, len(outpatientReply.List))
	for _, v := range outpatientReply.List {
		list = append(list, OutpatientReportList{
			ReportNum:   v.ReportNum,
			PatientNum:  v.PatientNum,
			PatientName: v.PatientName,
			PatientSex:  v.PatientSex,
			PatientAge:  v.PatientAge,
		})
	}

	reply.List = list

	return nil
}
