package virhis

//科室信息
type DepartmentListReq struct{}
type DepartmentListReply struct {
	List []Department
}
type Department struct {
	Code        string `json:"code"`        //科室编码
	Name        string `json:"name"`        //名称
	ParentId    string `json:"parent_id"`   //父类科室ID
	Pinyin      string `json:"pinyin"`      //科室拼音码
	Position    string `json:"position"`    //科室地址
	Ichnography string `json:"ichnography"` //平面图
	Description string `json:"description"` //科室描述
	Remark      string `json:"remark"`      //备注
	//BranchCode  string //分院编码
	//BranchName  string //分院名称
}

//医生信息
type DoctorListReq struct{}
type DoctorListReply struct {
	List []Doctor
}
type Doctor struct {
	Code           string //医生编码
	Name           string //姓名
	Title          string //职称
	Pinyin         string //拼音码
	Introduction   string //介绍
	Specialty      string //擅长
	Avatar         string //头像,照片
	Sex            string //医生性别
	Phone          string //医生电话
	DepartmentCode string //科室编码
	DepartmentName string //科室名称
}

// 获得全院排班信息
type ScheduleListReq struct {
	BeginDate string //排班开始时间  yyyy-MM-dd hh:mm:ss
	EndDate   string //排班结束时间  yyyy-MM-dd hh:mm:ss
}
type ScheduleListReply struct {
	List []Schedule
}
type Schedule struct {
	Code               string //排班编码
	Date               string //排班日期 yyyy-MM-dd hh:mm:ss
	MaxNum             int    //最大挂号数
	RegistrationCode   string //挂号种类编码
	RegistrationType   string //挂号种类 1:普通号 2:专家号....
	UseNum             string //已挂号数
	RemainNum          string //剩余号数
	DepartmentCode     string //科室编码
	DepartmentName     string //科室名称
	DoctorCode         string //医生编码
	DoctorName         string //医生姓名
	DoctorTitle        string //医生职称
	DoctorIntroduction string //医生简介
	RegistrationFee    string //挂号金额 精确到分
	BeginDate          string //开始接诊时间  yyyy-MM-dd hh:mm:ss
	EndDate            string //停止接诊时间  yyyy-MM-dd hh:mm:ss
}

//挂号种类
type RegistrationTypeListReq struct {
}
type RegistrationTypeListReply struct {
	List []RegistrationType
}
type RegistrationType struct {
	Code string //挂号种类编码
	Name string //挂号种类名称
	Fee  string //挂号费
}

// 办理就诊卡
type MedicalCardReq struct {
	Name              string //姓名
	Phone             string //电话
	IdCard            string //身份证
	Sex               int    //性别 1：女； 2：男
	SocialSecurityNum string //社保号
	Emergency         string //联系人
	EmergencyPhone    string //联系人电话
	EmergencyAddr     string //联系人地址
}
type MedicalCardReply struct {
	MedicalCard string //就诊卡号
}

//就诊卡校验
type MedicalCardVerifyReq struct {
	Name        string //患者姓名
	Phone       string //患者手机号
	IdCard      string //患者身份证
	MedicalCard string //就诊卡号
}
type MedicalCardVerifyReply struct {
	Status string //状态:  1. 成功 2. 身份证号不存在 3. 身份证和姓名或就诊卡号不匹配
}

//挂号
type RegistrationReq struct {
	ReservationDate      string //预约时间 yyyy-MM-dd hh:mm:ss
	ScheduleCode         string //排班编号
	MedicalCard          string //医疗卡/就诊卡
	IdCard               string //身份证
	PatientName          string //患者姓名
	PatientSex           string //患者性别
	PatientPhone         string //患者电话
	DoctorCode           string //医生编码
	DepartmentCode       string //科室编码
	RegistrationTypeCode string //挂号类型编码

}
type RegistrationReply struct {
	ReservationNum   string //预约号
	ReservationDate  string //预约时间 yyyy-MM-dd hh:mm:ss
	RegistrationDate string //登记时间 yyyy-MM-dd hh:mm:ss
	Fee              string //挂号费，支付金额
	SchedulingCode   string //排班编号
	DepartmentName   string //科室名称
	DoctorName       string //医生名称
	PatientName      string //患者姓名
}

// 取消挂号
type RegistrationCancelReq struct {
	ReservationNum string //预约号
	Fee            string //金额
	CancelReason   string //取消原因
	Date           string //取消时间
	TransactionNum string //交易流水号
	OrderNum       string //订单编号
}
type RegistrationCancelReply struct {
}

// 挂号缴费
type RegistrationPayReq struct {
	ReservationNum string //预约号
	TransactionNum string //交易流水号
	OrderNum       string //订单编号
	Fee            string //支付金额
	PayDate        string //支付时间
	PayType        string //支付方式   只支持微信支付 WX  1:WX
}
type RegistrationPayReply struct {
	QueueNum       string //排队序号
	ReservationNum string //预约号/挂号门诊号
	ConfirmDate    string //支付确认时间
	OutpatientNum  string //门诊号
	Position       string //就诊位置
}

//门诊记录
type OutpatientRecordListReq struct {
	CardNum   string //证件号
	CardType  string //证件类型  1:身份证, 2: 就诊卡, 3:门诊号
	BeginDate string //查询开始时间
	EndDate   string //查询结束时间
}
type OutpatientRecordListReply struct {
	List []OutpatientRecord
}
type OutpatientRecord struct {
	OutpatientNum  string //门诊号
	Name           string //姓名
	IdCard         string //身份证
	Sex            string //性别
	Age            string //年龄
	DepartmentCode string //科室编码
	DepartmentName string //科室名称
	DoctorCode     string //医生编码
	DoctorName     string //医生姓名
	TotalFee       string //总费用
	OutpatientDate string //就诊时间 yyyy-MM-dd hh:mm:ss
}

//门诊病历
type OutpatientMedicalRecordListReq struct {
	CardNum   string //证件号
	CardType  string //证件类型  1:门诊号
	BeginDate string //查询开始时间
	EndDate   string //查询结束时间
}
type OutpatientMedicalRecordListReply struct {
	OutpatientNum            string // 门诊号
	Date                     string // 就诊时间
	ChiefComplaint           string // 主诉
	HistoryOfPresentIllness  string // 现病史
	HistoryOfPastIllness     string // 既往史
	PersonalHistory          string // 个人史
	FamilyHistory            string // 家族史
	AllergicHistory          string // 过敏史
	MenstrualHistory         string // 月经史
	MarriagePregnancyHistory string // 婚孕史
	Diagnosis                string // 诊断，诊断内容
	PhysicalExamination      string // 体格检查,体温,脉搏,血压, 呼吸
	InspectionResult         string // 检查结果
	Suggestion               string // 处理意见
	BriefHistoryOfDisease    string // 简要病史
	MedicalCard              string // 就诊卡号
	PatientName              string // 患者姓名
	PatientSex               string // 患者性别
	PatientAge               string // 患者年龄
	PatientAddr              string // 患者地址
	ChineseMedicineDiagnosis string // 中医诊断
	DepartmentName           string // 科室名称
	DoctorName               string // 医生姓名
	ReferenceRange           string // 参考范围
}

//门诊检查检验报告列表
type OutpatientReportListReq struct {
	CardNum    string //证件号
	CardType   string //证件类型  1:身份证, 2: 就诊卡, 3:门诊号
	ReportType int    //报告单类型（0 所有1 检验 2 检查）
	BeginDate  string //查询开始时间
	EndDate    string //查询结束时间
}
type OutpatientReportListReply struct {
	List []OutpatientReport
}
type OutpatientReport struct {
	ReportNum                 string //报告单号
	ReportStatus              string //报告单状态 1:未收费 2:未执行 3:已执行 4:已报告
	ReportName                string //报告单项目名称
	RecordDate                string //记录时间
	InspectionDate            string //送检时间
	RecorderName              string //记录人名称
	RecorderCode              string //记录人编码
	VisType                   string //就诊类型
	ReportType                string //报告单类型
	PatientName               string //病人姓名
	PatientSex                string //病人性别
	PatientAge                string //病人年龄
	ApplicationDepartmentCode string //申请执行科室编码
	ApplicationDepartmentName string //申请执行科室名称
	ExecutionDepartmentCode   string //执行科室编码
	ExecutionDepartmentName   string //执行科室名称
}

//检查报告详情
type OutpatientInspectionReportDetailReq struct {
	CardNum   string //证件号
	CardType  string //证件类型  1:报告单号
	BeginDate string //查询开始时间
	EndDate   string //查询结束时间
}
type OutpatientInspectionReportDetailReply struct {
	ReportNum        string // 报告单号
	ReportName       string // 报告单项目名称
	SubName          string // 子项目类型
	CheckDate        string // 检查日期
	DepartmentName   string // 科室名称
	InspectionResult string // 检查结果
	DoctorAdvice     string // 医生建议
	Description      string // 描述
	ReportDate       string // 报告时间
	ApplyDoctorName  string // 申请医生
	VerifyDoctorName string // 审核医生
	ReportDoctorName string // 报告医生
	ImgDiagnosis     string // 影像诊断
	ImgAdvice        string // 影像建议
	ImgUrl           string // 影像urls
	PatientName      string // 病人姓名
	PatientSex       string // 病人性别
	PatientAge       string // 病人年龄
}

//检验报告详情
type OutpatientLaboratoryReportDetailReq struct {
	CardNum   string //证件号
	CardType  string //证件类型  1:报告单号
	BeginDate string //查询开始时间
	EndDate   string //查询结束时间
}
type OutpatientLaboratoryReportDetailReply struct {
	ReportNum        string // 报告单号
	ReportName       string // 报告单项目名称
	SubName          string // 子项目类型
	CheckDate        string // 检查日期
	ApplyDoctorName  string // 申请医生
	VerifyDoctorName string // 审核医生
	ReportDoctorName string // 报告医生
	ReportDate       string // 报告日期
	ReportItemCode   string // 报告单子项编码
	ReportItemName   string // 报告单子项名称
	ReportItemEnName string // 报告单子项英文名
	Status           string // 高低箭头(↑ ↓)
	Indicator        string // 状态: -1：偏低  0：正常 1：偏高 2其他
	Value            string // 检查结果值
	Unit             string // 检查结果值单位
	ValueRange       string // 正常范围
	ValueMin         string // 参考最低值
	ValueMax         string // 参考最高值
	DepartmentName   string // 科室名称
	PatientName      string // 病人姓名
	PatientSex       string // 病人性别
	PatientAge       string // 病人年龄
}

//门诊费用明细
type OutpatientFeeListReq struct {
	CardNum   string //证件号
	CardType  string //证件类型  1:门诊号
	BeginDate string //查询开始时间
	EndDate   string //查询结束时间
}
type OutpatientFeeListReply struct {
	List []OutpatientFeeDetail
}
type OutpatientFeeDetail struct {
	OutpatientNum           string // 门诊号
	PayNum                  string // 收费单号
	InvoiceNum              string // 发票号
	DepartmentCode          string // 科室编码
	DepartmentName          string // 科室名称
	DoctorCode              string // 医生编码
	DoctorName              string // 医生姓名
	AccountDate             string // 记账时间
	ItemCode                string // 收费项目编码
	ItemName                string // 项目名称
	UnitPrice               string // 单价
	Quantity                string // 数量
	Unit                    string // 单位
	AmountReceivableFee     string // 应收金额
	AmountReceiptsFee       string // 实收金额
	RemarkAmountFee         string // 记账金额
	CombinedItemCode        string // 组合项目编码
	CombinedItemName        string // 组合项目名称
	ExecutionDepartmentCode string // 执行科室编码
	ExecutionDepartmentName string // 执行科室名称
	ChargeCode              string // 收费方式编码
	ChargeName              string // 收费方式名称
	BillType                string // 单据类型
	BillNum                 string // 单据号

}

//患者门诊处方信息列表
type OutpatientPrescriptionListReq struct {
	CardNum   string //证件号
	CardType  string //证件类型   1:门诊号
	BeginDate string //查询开始时间
	EndDate   string //查询结束时间
}
type OutpatientPrescriptionListReply struct {
	List []OutpatientPrescription
}
type OutpatientPrescription struct {
	OutpatientNum      string // 门诊号
	MedicalCard        string // 医疗卡号
	PatientName        string // 患者姓名
	PrescriptionNum    string // 处方号
	BillDepartmentCode string // 开单科室编码
	BillDepartmentName string // 开单科室名称
	DoctorCode         string // 医生编码
	DoctorName         string // 医生姓名
	DispensingNum      string // 发药人工号
	ChineseMedicineNum string // 付数（中药）
	TotalFee           string // 总金额
	BillDate           string // 开单日期
	DispensingDate     string // 发药日期
}

//处方明细
type OutpatientPrescriptionDetailReq struct {
	CardNum         string // 证件号
	CardType        string // 证件类型 1:门诊号
	PrescriptionNum string // 处方号
}
type OutpatientPrescriptionDetailReply struct {
	PrescriptionNum    string // 门诊号
	PrescriptionStatus string // 处方状态
	PharmacyUnit       string // 药房单位
	Fee                string // 金额
	Standard           string // 药品规格
	Quantity           string // 药品数量
	UnitPrice          string // 单价
	ItemName           string // 项目名称
	MedicineCode       string // 药品编码
	MedicineName       string // 药品名称
	Usage              string // 药品用法
}

//门诊缴费查询
type OutpatientPayQueryReq struct {
	CardNum  string // 证件号
	CardType string // 证件类型 1:身份证, 2: 就诊卡, 3:门诊号
}
type OutpatientPayQueryReply struct {
	List []OutpatientPayQuery
}
type OutpatientPayQuery struct {
	OutpatientNum             string // 门诊号
	PayNum                    string // 缴费单据号
	PayType                   string // 费用类型
	DoctorName                string // 医生姓名
	Fee                       string // 金额
	ItemName                  string // 项目名称
	BillDate                  string // 开单日期
	ApplicationDepartmentCode string // 申请执行科室编码
	ApplicationDepartmentName string // 申请执行科室名称
	Status                    string // 状态 1. 已缴费 2. 未交费。 3.已退费
}

// 门诊缴费
type OutpatientPayReq struct {
	CardNum        string // 证件号
	CardType       string // 证件类型  1:门诊号
	TotalFee       string // 总金额
	PayType        string // 支付方式
	PayNum         string // 缴费单据号
	OrderNum       string // 订单号
	TransactionNum string // 交易流水号

}
type OutpatientPayReply struct {
	OutpatientNum string // 门诊号
	Name          string // 患者姓名
	ChargeNum     string // 收费单号
	Fee           string // 金额
	AccountDate   string // 记账时间
	AccountFee    string // 记账金额
	Remark        string // 备注
}

//预约住院
type HospitalizationAppointmentReq struct {
	CardNum        string // 证件号
	CardType       string // 证件类型 1:身份证, 2: 就诊卡
	AdmissionDate  string // 入院时间
	Name           string // 患者姓名
	DepartmentName string // 入院科室
	FeeType        string // 费用类别
	Emergency      string // 联系人
	EmergencyPhone string // 联系人电话
	EmergencyAddr  string // 联系人地址
}
type HospitalizationAppointmentReply struct {
	Status             string // 状态
	HospitalizationNum string // 住院号
}

//住院记录
type HospitalizationRecordListReq struct {
	CardNum   string // 证件号
	CardType  string // 证件类型 1:身份证, 2: 就诊卡, 3:住院号
	BeginDate string // 查询开始时间 yyyy-MM-dd hh:mm:ss
	EndDate   string // 查询结束时间 yyyy-MM-dd hh:mm:ss
}
type HospitalizationRecordListReply struct {
	List []HospitalizationRecord
}
type HospitalizationRecord struct {
	HospitalizationNum    string // 住院号
	Name                  string // 患者姓名
	IdCard                string // 身份证
	Sex                   string // 性别
	Age                   string // 年龄
	MedicalCard           string // 医疗号
	ChargeType            string // 收费种类
	WardCode              string // 病区编码
	WardName              string // 病区名称
	DepartmentCode        string // 科室编码
	DepartmentName        string // 科室名称
	DoctorCode            string // 医生编码
	DoctorName            string // 医生姓名
	BedNum                string // 床号
	AdmissionDate         string // 入院时间
	DischargeDate         string // 出院时间
	PatientStatus         string // 病人状态
	HospitalizationStatus string // 住院状态
	Prepayment            string // 预交金
	ActualFee             string // 实际花费
	AccountBalance        string // 帐户余额
	MidwayCheckoutNum     int    // 中途结账次数
}

//住院病历
type HospitalizationMedicalRecordListReq struct {
	CardNum   string // 证件号
	CardType  string // 证件类型 1:住院号
	BeginDate string // 查询开始时间 yyyy-MM-dd hh:mm:ss
	EndDate   string // 查询结束时间 yyyy-MM-dd hh:mm:ss
}
type HospitalizationMedicalRecordListReply struct {
	HospitalizationNum       string // 住院号
	Date                     string // 就诊时间
	ChiefComplaint           string // 主诉
	HistoryOfPresentIllness  string // 现病史
	HistoryOfPastIllness     string // 既往史
	PersonalHistory          string // 个人史
	FamilyHistory            string // 家族史
	AllergicHistory          string // 过敏史
	MenstrualHistory         string // 月经史
	MarriagePregnancyHistory string // 婚孕史
	Diagnosis                string // 诊断，诊断内容
	PhysicalExamination      string // 体格检查,体温,脉搏,血压, 呼吸
	InspectionResult         string // 检查结果
	Suggestion               string // 处理意见
	BriefHistoryOfDisease    string // 简要病史
	CardNum                  string // 就诊卡号
	PatientName              string // 患者姓名
	PatientSex               string // 患者性别
	PatientAge               string // 患者年龄
	PatientAddr              string // 患者地址
	ChineseMedicineDiagnosis string // 中医诊断
	DepartmentName           string // 科室名称
	DoctorName               string // 医生姓名
	ReferenceRange           string // 参考范围
}

// 住院检查检验报告列表
type HospitalizationReportListReq struct {
	CardNum    string // 证件号
	CardType   string // 证件类型 1:身份证, 2: 就诊卡, 3:住院号
	ReportType string // 报告单类型 0 所有1 检验 2 检查
	BeginDate  string // 查询开始时间 yyyy-MM-dd hh:mm:ss
	EndDate    string // 查询结束时间 yyyy-MM-dd hh:mm:ss
}
type HospitalizationReportListReply struct {
	List []HospitalizationInspectionReportList
}
type HospitalizationInspectionReportList struct {
	ReportNum                 string // 报告单号
	ReportStatus              string // 报告单状态
	ReportName                string // 报告单项目名称
	RecordDate                string // 记录时间
	InspectionDate            string // 送检时间
	RecorderName              string // 记录人名称
	RecorderCode              string // 记录人编码
	VisType                   string // 就诊类型
	ReportType                string // 报告单类型
	PatientName               string // 病人姓名
	PatientSex                string // 病人性别
	PatientAge                string // 病人年龄
	ApplicationDepartmentCode string // 申请执行科室编码
	ApplicationDepartmentName string // 申请执行科室名称
	ExecutionDepartmentCode   string // 执行科室编码
	ExecutionDepartmentName   string // 执行科室名称
}

//住院检查报告详情
type HospitalizationInspectionReportDetailReq struct {
	CardNum   string // 证件号
	CardType  string // 证件类型 1:报告单号
	BeginDate string // 查询开始时间 yyyy-MM-dd hh:mm:ss
	EndDate   string // 查询结束时间 yyyy-MM-dd hh:mm:ss
}
type HospitalizationInspectionReportDetailReply struct {
	ReportNum        string // 报告单号
	ReportName       string // 报告单项目名称
	SubName          string // 子项目类型
	CheckDate        string // 检查日期
	DepartmentName   string // 科室名称
	InspectionResult string // 检查结果
	DoctorAdvice     string // 医生建议
	Description      string // 描述
	ReportDate       string // 报告时间
	ApplyDoctorName  string // 申请医生
	VerifyDoctorName string // 审核医生
	ReportDoctorName string // 报告医生
	ImgDiagnosis     string // 影像诊断
	ImgAdvice        string // 影像建议
	Img              string // 影像图片
	PatientName      string // 病人姓名
	PatientSex       string // 病人性别
	PatientAge       string // 病人年龄
}

//住院检验报告详情
type HospitalizationLaboratoryReportDetailReq struct {
	CardNum   string // 证件号
	CardType  string // 证件类型 1:报告单号
	BeginDate string // 查询开始时间 yyyy-MM-dd hh:mm:ss
	EndDate   string // 查询结束时间 yyyy-MM-dd hh:mm:ss

}
type HospitalizationLaboratoryReportDetailReply struct {
	ReportNum        string // 报告单号
	ReportName       string // 报告单项目名称
	SubName          string // 子项目类型
	CheckDate        string // 检查日期
	ApplyDoctorName  string // 申请医生
	VerifyDoctorName string // 审核医生
	ReportDoctorName string // 报告医生
	ReportDate       string // 报告日期
	ReportItemCode   string // 报告单子项编码
	ReportItemName   string // 报告单子项名称
	ReportItemEnName string // 报告单子项英文名
	Status           string // 高低箭头(↑ ↓)
	Indicator        string // 状态: -1：偏低  0：正常 1：偏高 2其他
	Value            string // 检查结果值
	Unit             string // 检查结果值单位
	ValueRange       string // 正常范围
	ValueMin         string // 参考最低值
	ValueMax         string // 参考最高值
	DepartmentName   string // 科室名称
	PatientName      string // 病人姓名
	PatientSex       string // 病人性别
	PatientAge       string // 病人年龄
}

//住院预交金缴费
type HospitalizationPayReq struct {
	CardNum        string // 证件号
	CardType       string // 证件类型 1:住院号
	PrepaymentFee  string // 预交金
	PayType        string // 支付方式 只支持微信支付 WX  1:WX
	TransactionNum string // 交易流水号
	OrderNum       string // 订单号
}
type HospitalizationPayReply struct {
	HospitalizationNum string // 住院号
	Name               string // 姓名
	Fee                string // 金额
	ReceiptNum         string // 收据号
	Date               string // 时间
	Remark             string // 备注
}

//住院预交金余额查询
type HospitalizationBalanceReq struct {
	CardNum  string // 证件号
	CardType string // 证件类型 1:住院号
}
type HospitalizationBalanceReply struct {
	HospitalizationNum string // 住院号
	Name               string // 患者姓名
	BalanceFee         string // 余额 精确到分
}

//住院费用明细
type HospitalizationFeeListReq struct {
	CardNum   string // 证件号
	CardType  string // 证件类型 1:住院号
	BeginDate string // 查询开始时间 yyyy-MM-dd hh:mm:ss
	EndDate   string // 查询结束时间 yyyy-MM-dd hh:mm:ss
}

type HospitalizationFeeListReply struct {
	List []HospitalizationFeeRecord
}

type HospitalizationFeeRecord struct {
	HospitalizationNum      string // 住院号
	Bill                    string // 记账单
	Date                    string // 时间
	MedicalCard             string // 医疗号
	ChargeCode              string // 收费种类编码
	ChargeName              string // 收费种类编码
	MidwayCheckoutNum       int    // 中途结账次数
	ItemCode                string // 项目编码
	ItemName                string // 项目名称
	UnitPrice               string // 单价
	Quantity                string // 数量
	Unit                    string // 单位
	AmountReceivableFee     string // 应收金额
	AmountReceiptsFee       string // 实收金额
	DoctorCode              string // 医生编码
	DoctorName              string // 医生姓名
	ExecutionDepartmentCode string // 执行科室编码
	ExecutionDepartmentName string // 执行科室名称
	BillDepartmentCode      string // 开单科室编码
	BillDepartmentName      string // 开单科室名称
	CombinedItemCode        string // 组合项目编码
	CombinedItemName        string // 组合项目名称
	FinanceStatistics       string // 财务统计
	FeeStatistics           string // 费用统计

}

type HIS interface {
	//基础
	DepartmentList(*DepartmentListReq, *DepartmentListReply) error                   //获取科室信息
	DoctorList(*DoctorListReq, *DoctorListReply) error                               //获取医生
	ScheduleList(*ScheduleListReq, *ScheduleListReply) error                         //医院全院排班
	RegistrationTypeList(*RegistrationTypeListReq, *RegistrationTypeListReply) error //挂号种类
	MedicalCard(*MedicalCardReq, *MedicalCardReply) error                            //办理就诊卡
	MedicalCardVerify(*MedicalCardVerifyReq, *MedicalCardVerifyReply) error          //就诊卡验证

	//挂号
	Registration(*RegistrationReq, *RegistrationReply) error                   //预约挂号
	RegistrationCancel(*RegistrationCancelReq, *RegistrationCancelReply) error //取消挂号
	RegistrationPay(*RegistrationPayReq, *RegistrationPayReply) error          //挂号缴费

	//门诊
	OutpatientRecordList(*OutpatientRecordListReq, *OutpatientRecordListReply) error                                     //门诊就诊记录
	OutpatientMedicalRecordList(*OutpatientMedicalRecordListReq, *OutpatientMedicalRecordListReply) error                //门诊病历
	OutpatientReportList(*OutpatientReportListReq, *OutpatientReportListReply) error                                     //门诊检查检验报告列表
	OutpatientInspectionReportDetail(*OutpatientInspectionReportDetailReq, *OutpatientInspectionReportDetailReply) error //门诊检查报告详情
	OutpatientLaboratoryReportDetail(*OutpatientLaboratoryReportDetailReq, *OutpatientLaboratoryReportDetailReply) error //门诊检验报告详情
	OutpatientFeeList(*OutpatientFeeListReq, *OutpatientFeeListReply) error                                              //门诊费用明细
	OutpatientPrescriptionList(*OutpatientPrescriptionListReq, *OutpatientPrescriptionListReply) error                   //患者门诊处方信息列表
	OutpatientPrescriptionDetail(*OutpatientPrescriptionDetailReq, *OutpatientPrescriptionDetailReply) error             //处方明细
	OutpatientPayQuery(*OutpatientPayQueryReq, *OutpatientPayQueryReply) error                                           //门诊待缴费查询, 未分单
	OutpatientPay(*OutpatientPayReq, *OutpatientPayReply) error                                                          //门诊缴费

	//住院
	HospitalizationAppointment(*HospitalizationAppointmentReq, *HospitalizationAppointmentReply) error                                  //预约住院
	HospitalizationRecordList(*HospitalizationRecordListReq, *HospitalizationRecordListReply) error                                     //住院记录
	HospitalizationMedicalRecordList(*HospitalizationMedicalRecordListReq, *HospitalizationMedicalRecordListReply) error                //住院病历
	HospitalizationReportList(*HospitalizationReportListReq, *HospitalizationReportListReply) error                                     //住院检查检验报告列表
	HospitalizationInspectionReportDetail(*HospitalizationInspectionReportDetailReq, *HospitalizationInspectionReportDetailReply) error //住院检查报告详情
	HospitalizationLaboratoryReportDetail(*HospitalizationLaboratoryReportDetailReq, *HospitalizationLaboratoryReportDetailReply) error //住院检验报告详情
	HospitalizationPay(*HospitalizationPayReq, *HospitalizationPayReply) error                                                          //住院预交金缴费
	HospitalizationBalance(*HospitalizationBalanceReq, *HospitalizationBalanceReply) error                                              //住院预交金余额
	HospitalizationFeeList(req *HospitalizationFeeListReq, reply *HospitalizationFeeListReply) error                                    //住院费用明细
}
