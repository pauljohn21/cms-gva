package esgin

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/pauljohn21/cms-gva/server/global"
	"github.com/pauljohn21/cms-gva/server/model/cms"
	"github.com/pauljohn21/cms-gva/server/utils"

	"github.com/xuri/excelize/v2"

	"github.com/Esword618/unioffice/document"
	"github.com/Esword618/unioffice/schema/soo/wml"
)

func CrateTemplate(meLetter *cms.MeLetter) (int, string, error) {
	fmt.Println("开始生成模板")
	// 获取当前工作目录
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory:", err)
		return 0, "", err
	}
	// 拼接文件路径
	dataPath := filepath.Join(currentDir, meLetter.Respondent)
	fmt.Println(dataPath)

	data := readexecleToPerson(dataPath)
	comName := meLetter.Applicant
	var applicant *cms.Applicant
	// var comId string
	err = global.GVA_DB.Where("company = ?", comName).First(&applicant).Error
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("信用代码%s\n", applicant.Code)
	var sid *cms.MeLetter

	global.GVA_DB.Last(&sid, "id")
	pids := utils.GenerateOrderID(int(sid.ID))
	fmt.Println(pids)
	meLetter.Code = pids

	startTime, _ := utils.FormatTime(meLetter.StartCreatedAt)
	endTime, _ := utils.FormatTime(meLetter.EndCreatedAt)

	var namelist string
	for _, item := range data {
		namelist += fmt.Sprintf("%s,", item.Name)
	}
	bszh := meLetter.CoverageAllNzh
	fmt.Println("大写", bszh)
	bszh_lower := meLetter.CoverageAll
	bczrdata := fmt.Sprintf("财产保全保障人: %s与%s,因金融不良债权追偿纠纷案向法院提出财产保全申请,申请冻结被申请人名下价值人民币 %s元整 (小写:RMB %s)的银行存款,微信支付余额和微信支付功能,支付宝支付余额和支付宝支付功能以及其他等值货币资金如因申请人财产保全申请错误致使被申请人遭受经济损失,依法应由申请人承担的损害赔偿责任保险人承担连带赔偿责任,赔偿限额以人民币 %s元整 (小写: RMB %s)为限", comName, namelist, bszh, bszh_lower, bszh, bszh_lower)
	tbrddata := `  	1.投保人/被保险人将诚实谨慎行使诉讼权利保证无恶意诉讼或虚假诉讼的故意且与被告、被申请人无恶意串通。包括但不限于:		
	1.1 在合同有效期内保险标的的危险程度显著增加的（包括但不限于被申请人提供的证据足以推翻或动摇投保人/被保险人诉请的主要或关键事实的、一审法院作出不利于投保人/被保险人的判决等)，投保人/被保险人应当按照合同约定及时通知保险人；					
	1.2 应当解除保全措施情况发生后投保人/被保险人应及时解除财产保全措施；												
	
	1.3 投保人/被保险人起诉所依据的主要合同、主要文书证据没有伪造、变造；												
	1.4 投保人/被保险人在起诉状中陈述的事实基于善意无故意虚构或隐瞒；													
	1.5 保全金额系经投保人/被保险人合理评估和测算确定的，投保人/被保险人不存在不合理扩大保全金额的故意；					
	1.6 投保人/被保险人在起诉前已经征求过专业意见，对证明投保人/被保险人诉请的主要或关键事实有初步证明能力；				
	1.7 保险事故发生时或保险事故难以避免将要发生时，被保险人应当尽量采取必要的措施防止或减少损失，包括但不限于根据保险人的要求及时解除保全措施；										
	1.8 投保人/被保险人未经保险人同意不会擅自变更投保单中前述申明的保全方式或保全标的。									
	1.9如申请诉前财产保全的,保全申请人必须在30天内提起诉讼/申请仲裁，而且诉求金额与申请保全金额应大致相当，否则相应担保责任自动终止，保险人自始不承担任何保险责任。					
	2. 本保险合同的保险期间为自投保人/被保险人/申请人向法院提出诉讼财产保全申请之日起至保全损害之债诉讼时效届满时终止		
	3. 尊敬的客户: 为保障您的利益,我司提供保单查询和理赔咨询服务,如有相关问题请拨打我司24小时服务热线: 0451-82110855`

	// 创建一个新的Word文档
	doc := document.New()

	// Construct our header
	hdr := doc.AddHeader()
	parahdr := hdr.AddParagraph()
	runhdr := parahdr.AddRun()
	parahdr.Properties().SetAlignment(wml.ST_JcRight)
	runhdr.Properties().SetSize(18)
	runhdr.AddText("黑龙江省信仕正融资担保有限公司")
	parahdr1 := hdr.AddParagraph()
	runhdr1 := parahdr1.AddRun()
	parahdr1.Properties().SetAlignment(wml.ST_JcRight)
	runhdr1.Properties().SetSize(10)
	runhdr1.AddText("Heilongjiang Xinshizheng Financing Guarantee Co., Ltd")

	// Construct our footer

	fbr := doc.AddFooter()
	parafbr := fbr.AddParagraph()
	runfbr := parafbr.AddRun()
	parafbr.Properties().SetAlignment(wml.ST_JcLeft)
	runfbr.Properties().SetSize(10)
	runfbr.AddText("地址: 哈尔滨市香坊区哈平路29号4楼   电话:  0451-82110855")

	header := doc.AddParagraph()
	header.Properties().SetAlignment(wml.ST_JcCenter)
	header.Properties().SetStyle("Title")
	headerrun := header.AddRun()
	headerrun.Properties().SetBold(true)
	headerrun.AddText("诉讼财产保全责任险保单")

	pid := doc.AddParagraph()
	pid.Properties().SetAlignment(wml.ST_JcRight)
	pid.Properties().SetStyle("Heading2")
	pidf := pid.AddRun()
	pidf.Properties().SetBold(true)
	pidf.AddText("保单号: ")
	pidi := pid.AddRun()
	pidi.Properties().SetSize(12)
	pidi.Properties().SetBold(false)
	pidi.AddText(pids)

	fyid := doc.AddParagraph()
	fyid.Properties().SetAlignment(wml.ST_JcLeft)
	fyid.Properties().SetStyle("Heading2")
	fyidf := fyid.AddRun()
	fyidf.Properties().SetBold(true)
	fyidf.AddText("致:	")
	fyidf.AddTab()
	fyidi := fyid.AddRun()
	fyidi.Properties().SetBold(false)
	fyidi.Properties().SetSize(12)
	fyidi.AddText(meLetter.Court)

	bcr := doc.AddParagraph()
	bcr.Properties().SetAlignment(wml.ST_JcBoth)
	bcr.Properties().SetStyle("Heading2")
	bcrf := bcr.AddRun()
	bcrf.Properties().SetBold(true)
	bcrf.AddText("财产保全申请人:	")
	bcrf.AddTab()
	bcri := bcr.AddRun()
	bcri.Properties().SetSize(12)
	bcri.Properties().SetBold(false)
	bcri.AddText(fmt.Sprintf("%s, 						证件类型: 统一社会信用代码, 证件号码: %s", comName, applicant.Code))

	bsqr := doc.AddParagraph()
	bsqr.Properties().SetAlignment(wml.ST_JcLeft)
	rsqrun := bsqr.AddRun()
	rsqrun.Properties().SetBold(true)
	rsqrun.AddText("财产保全被申请人: ")
	// 添加一个表格
	table := doc.AddTable()
	table.Properties().SetWidthAuto()
	table.Properties().SetAlignment(wml.ST_JcTableEnd)
	table.Properties().SetStyle("TableGrid")

	// 循环添加数据
	// 示例数据实际应该从数据库或其他数据源获取

	// fmt.Println(datalist)
	for _, item := range data {
		row := table.AddRow()
		row.AddCell().AddParagraph().AddRun().AddText(item.Name)
		row.AddCell().AddParagraph().AddRun().AddText(item.Gender)
		row.AddCell().AddParagraph().AddRun().AddText("证件类型: " + item.Idtype)
		row.AddCell().AddParagraph().AddRun().AddText("证件号码: " + item.IdNumber)
	}

	doc.AddParagraph()
	bcjr := doc.AddParagraph()
	bcjr.Properties().SetAlignment(wml.ST_JcLeft)
	bcjr.Properties().SetStyle("Heading2")
	bcjrf := bcjr.AddRun()
	bcjrf.Properties().SetBold(true)
	bcjrf.AddText("保险金额:	")
	bcjri := bcjr.AddRun()
	bcjri.Properties().SetBold(false)
	bcjri.Properties().SetSize(12)
	bcjri.AddText(fmt.Sprintf("%s 元整 (小写: RMB %s)", bszh, bszh_lower))

	doc.AddParagraph()
	bxzr := doc.AddParagraph()
	bxzr.Properties().SetAlignment(wml.ST_JcLeft)
	bxzr.Properties().SetStyle("Heading2")
	bxzrf := bxzr.AddRun()
	bxzrf.Properties().SetBold(true)
	bxzrf.AddText("保险责任: ")
	bxzrl := doc.AddParagraph()
	bxzrl.Properties().SetStartIndent(60)
	bxzrl.Properties().SetAlignment(wml.ST_JcBoth)
	bxzrl.AddRun().AddText(bczrdata)

	doc.AddParagraph()
	bxqj := doc.AddParagraph()
	bxqj.Properties().SetAlignment(wml.ST_JcLeft)
	bxqj.Properties().SetStyle("Heading2")
	bxqjf := bxqj.AddRun()
	bxqjf.Properties().SetBold(true)
	bxqjf.AddText("保险期间:	")
	bxqji := bxqj.AddRun()
	bxqji.Properties().SetBold(false)
	bxqji.Properties().SetSize(12)
	bxqji.AddText(fmt.Sprintf("自 %s 起至 %s 止", startTime, endTime))

	doc.AddParagraph()
	hsbf := doc.AddParagraph()
	hsbf.Properties().SetAlignment(wml.ST_JcLeft)
	hsbf.Properties().SetStyle("Heading2")
	hsbff := hsbf.AddRun()
	hsbff.Properties().SetBold(true)
	hsbff.AddText("含税总保费:	")
	hsbfi := hsbf.AddRun()
	hsbfi.Properties().SetBold(false)
	hsbfi.Properties().SetSize(12)
	hsbfi.AddText(fmt.Sprintf("%s 元整 (小写: RMB %s)", meLetter.CoverageNzh, meLetter.Coverage))

	doc.AddParagraph()
	shtk := doc.AddParagraph()
	shtk.Properties().SetAlignment(wml.ST_JcLeft)
	shtk.Properties().SetStyle("Heading2")
	shtkf := shtk.AddRun()
	shtkf.Properties().SetBold(true)
	shtkf.AddText("适用条款:	")
	shtki := shtk.AddRun()
	shtki.Properties().SetBold(false)
	shtki.Properties().SetSize(12)
	shtki.AddText("诉讼财产保全责任险条款")

	doc.AddParagraph()
	tbrd1 := doc.AddParagraph()
	tbrd1.Properties().SetAlignment(wml.ST_JcLeft)
	tbrd2 := tbrd1.AddRun()
	tbrd2.Properties().SetBold(true)
	tbrd2.AddText("特别约定: ")
	tbrd := doc.AddParagraph()
	tbrd.Properties().SetStartIndent(60)
	tbrd.Properties().SetAlignment(wml.ST_JcBoth)
	tbrd.AddRun().AddText(tbrddata)

	parahdr = doc.AddParagraph()
	sectionhdr := parahdr.Properties().AddSection(wml.ST_SectionMarkNextColumn)
	sectionhdr.SetHeader(hdr, wml.ST_HdrFtrDefault)
	sectionhdr.SetFooter(fbr, wml.ST_HdrFtrDefault)

	header1 := doc.AddParagraph()
	header1.Properties().SetAlignment(wml.ST_JcCenter)
	header1.Properties().SetStyle("Title")
	headerrun1 := header1.AddRun()
	headerrun1.AddText("诉讼财产保全责任保险条款")

	header2 := doc.AddParagraph()
	header2.Properties().SetAlignment(wml.ST_JcCenter)
	header2.Properties().SetStyle("Heading1")
	headerrun2 := header2.AddRun()
	headerrun2.Properties().SetSize(18)
	headerrun2.Properties().SetBold(true)
	headerrun2.AddText("总则")
	doc.AddParagraph()

	td1 := doc.AddParagraph()
	td1.Properties().SetAlignment(wml.ST_JcLeft)
	td1f := td1.AddRun()
	td1f.Properties().SetBold(true)
	td1f.AddText("第一条: ")
	td1i := td1.AddRun()
	td1i.Properties().SetBold(false)
	td1i.AddText("本保险合同由保险条款、投保单、保险单、保险凭证以及批单组成。凡涉及本保险合同的约定，均应采用书面形式。")
	doc.AddParagraph()

	td2 := doc.AddParagraph()
	td2.Properties().SetAlignment(wml.ST_JcLeft)
	td2f := td2.AddRun()
	td2f.Properties().SetBold(true)
	td2f.AddText("第二条: ")
	td2i := td2.AddRun()
	td2i.Properties().SetBold(false)
	td2i.AddText("本保险合同中的诉讼财产保全，是指人民法院根据民事诉讼利害关系人或当事人的申请，在起诉前或诉讼过程中针对被申请人的财产或争议标的物，采取查封、扣押、冻结等限制处分的强制措施，以防止转移、隐匿、变卖财产，使其处于人民法院的有效监控之下的司法行为，以保障未来生效判决得以执行。诉讼财产保全包括诉前保全和诉中保全。")
	doc.AddParagraph()

	td3 := doc.AddParagraph()
	td3.Properties().SetAlignment(wml.ST_JcLeft)
	td3f := td3.AddRun()
	td3f.Properties().SetBold(true)
	td3f.AddText("第三条: ")
	td3i := td3.AddRun()
	td3i.Properties().SetBold(false)
	td3i.AddText("本保险合同中的申请人，是指在起诉前或诉讼过程中向人民法院提起诉讼财产保全的利害关系人或当事人；被申请人，是指因上述申请而被人民法院依法采取诉讼财产保全的民事诉讼利害关系人或当事人，具体在保险单中载明。")
	doc.AddParagraph()

	td4 := doc.AddParagraph()
	td4.Properties().SetAlignment(wml.ST_JcLeft)
	td4f := td4.AddRun()
	td4f.Properties().SetBold(true)
	td4f.AddText("第四条: ")
	td4i := td4.AddRun()
	td4i.Properties().SetBold(false)
	td4i.AddText("本保险合同中被保险人是指因民事纠纷向法院申请财产保全的民事诉讼当事人或利害关系人。投保人是指与保险人订立本保险合同，并按照合同负有支付保险费义务的人。")
	doc.AddParagraph()

	bczr := doc.AddParagraph()
	bczr.Properties().SetAlignment(wml.ST_JcCenter)
	bczrf := bczr.AddRun()
	bczrf.Properties().SetSize(12)
	bczrf.Properties().SetBold(true)
	bczrf.AddText("保险责任")

	td5 := doc.AddParagraph()
	td5.Properties().SetAlignment(wml.ST_JcLeft)
	td5f := td5.AddRun()
	td5f.Properties().SetBold(true)
	td5f.AddText("第五条: ")
	td5i := td5.AddRun()
	td5i.Properties().SetBold(false)
	td5i.AddText(" 在保险期间内，被保险人向人民法院或仲裁机构提出诉讼财产保全申请并经人民法院裁定同意，如因申请错误致使被申请人遭受损失，由被申请人另行提起诉讼财产保全损害责任纠纷诉讼，经人民法院生效判决认定申请人应承担损害赔偿责任的，保险人根据本条款的规定不可抗辩地在赔偿限额内承担赔偿责任")

	bczr1 := doc.AddParagraph()
	bczr1.Properties().SetAlignment(wml.ST_JcCenter)
	bczrf1 := bczr1.AddRun()
	bczrf1.Properties().SetSize(12)
	bczrf1.Properties().SetBold(true)
	bczrf1.AddText("责任免除")

	td6 := doc.AddParagraph()
	td6.Properties().SetAlignment(wml.ST_JcLeft)
	td6f := td6.AddRun()
	td6f.Properties().SetBold(true)
	td6f.AddText("第六条: ")
	td6i := td6.AddRun()
	td6i.Properties().SetBold(false)
	td6i.AddText("依据法院判决被保险人无需承担的经济损失，保险人不负责赔偿。")

	bczr2 := doc.AddParagraph()
	bczr2.Properties().SetAlignment(wml.ST_JcCenter)
	bczrf2 := bczr2.AddRun()
	bczrf2.Properties().SetSize(12)
	bczrf2.Properties().SetBold(true)
	bczrf2.AddText("赔偿限额")

	td7 := doc.AddParagraph()
	td7.Properties().SetAlignment(wml.ST_JcLeft)
	td7f := td7.AddRun()
	td7f.Properties().SetBold(true)
	td7f.AddText("第七条: ")
	td7i := td7.AddRun()
	td7i.Properties().SetBold(false)
	td7i.AddText("本保险的赔偿限额为诉讼财产保全的申请保全金额，具体以保险单上载明的保险金额为准。")

	bczr3 := doc.AddParagraph()
	bczr3.Properties().SetAlignment(wml.ST_JcCenter)
	bczrf3 := bczr3.AddRun()
	bczrf3.Properties().SetSize(12)
	bczrf3.Properties().SetBold(true)
	bczrf3.AddText("保险期间")

	td8 := doc.AddParagraph()
	td8.Properties().SetAlignment(wml.ST_JcLeft)
	td8f := td8.AddRun()
	td8f.Properties().SetBold(true)
	td8f.AddText("第八条: ")
	td8i := td8.AddRun()
	td8i.Properties().SetBold(false)
	td8i.AddText("本保险的保险期间为自被保险人向法院提出诉讼财产保全申请之日起至保全损害之债诉讼时效届满时终止。")

	bczr4 := doc.AddParagraph()
	bczr4.Properties().SetAlignment(wml.ST_JcCenter)
	bczrf4 := bczr4.AddRun()
	bczrf4.Properties().SetSize(12)
	bczrf4.Properties().SetBold(true)
	bczrf4.AddText("保险人义务")

	td9 := doc.AddParagraph()
	td9.Properties().SetAlignment(wml.ST_JcLeft)
	td9f := td9.AddRun()
	td9f.Properties().SetBold(true)
	td9f.AddText("第九条: ")
	td9i := td9.AddRun()
	td9i.Properties().SetBold(false)
	td9i.AddText("订立保险合同时，采用保险人提供的格式条款的，保险人向投保人提供的投保单应附格式条款，保险人应当向投保人说明保险合同的内容。对保险合同中免除保险人责任的条款，保险人在订立合同时应当在投保单、保险单或者其他保险凭证上做出足以引起投保人注意的提示，并对该条款的内容以书面或者口头形式向投保人做出明确说明；未作提示或者明确说明的，该条款不产生效力。")
	doc.AddParagraph()

	td10 := doc.AddParagraph()
	td10.Properties().SetAlignment(wml.ST_JcLeft)
	td10f := td10.AddRun()
	td10f.Properties().SetBold(true)
	td10f.AddText("第十条: ")
	td10i := td10.AddRun()
	td10i.Properties().SetBold(false)
	td10i.AddText("本保险合同成立后，保险人应当及时向投保人签发保险单或者其他保险凭证。")
	doc.AddParagraph()

	td11 := doc.AddParagraph()
	td11.Properties().SetAlignment(wml.ST_JcLeft)
	td11f := td11.AddRun()
	td11f.Properties().SetBold(true)
	td11f.AddText("第十一条: ")
	td11i := td11.AddRun()
	td11i.Properties().SetBold(false)
	td11i.AddText("保险事故发生后，投保人、被保险人提供的有关索赔的证明和资料不完整的，保险人应当及时一次性通知投保人、被保险人补充提供。")
	doc.AddParagraph()

	td12 := doc.AddParagraph()
	td12.Properties().SetAlignment(wml.ST_JcLeft)
	td12f := td12.AddRun()
	td12f.Properties().SetBold(true)
	td12f.AddText("第十二条: ")
	td12i := td12.AddRun()
	td12i.Properties().SetBold(false)
	td12i.AddText(`保险人收到被保险人的赔偿保险金的请求后，应当及时做出是否属于保险责任的核定；情况复杂的，应当在三十日内做出核定，但因索赔人原因、客观原因、事故情况异常复杂、需要查证、需要等待第三方意见等原因需要更长处理时间的除外。																	保险人应当将核定的结果通知被保险人；对属于保险责任的，在与被保险人达成赔偿保险金的协议后十日内，履行赔偿保险金的义务。本保险合同对赔偿保险金的期限有约定的，保险人应当按照约定履行赔偿保险金的义务。保险人依照前款的规定做出核定后，对不属于保险责任的，应当自做出核定之日起三日内向被保险人发出拒绝赔偿保险金的通知书，并说明理由。`)

	bczr5 := doc.AddParagraph()
	bczr5.Properties().SetAlignment(wml.ST_JcCenter)
	bczrf5 := bczr5.AddRun()
	bczrf5.Properties().SetSize(12)
	bczrf5.Properties().SetBold(true)
	bczrf5.AddText("投保人、被保险人义务")

	td13 := doc.AddParagraph()
	td13.Properties().SetAlignment(wml.ST_JcLeft)
	td13f := td13.AddRun()
	td13f.Properties().SetBold(true)
	td13f.AddText("第十三条: ")
	td13i := td13.AddRun()
	td13i.Properties().SetBold(false)
	td13i.AddText("除另有约定外，投保人应当在保险合同成立时一次性交付保险费。")
	doc.AddParagraph()

	td14 := doc.AddParagraph()
	td14.Properties().SetAlignment(wml.ST_JcLeft)
	td14f := td14.AddRun()
	td14f.Properties().SetBold(true)
	td14f.AddText("第十四条: ")
	td14i := td14.AddRun()
	td14i.Properties().SetBold(false)
	td14i.AddText(" 被保险人应将所涉及的基础债权债务纠纷案件的任何重大进展情况自其知道或者应当知道之日起二十日内告知保险人，本条款所称重大进展情况包括但不限于案件中止、被驳回起诉、调解或判决等对案件进展有重要影响的情况。")
	doc.AddParagraph()

	td15 := doc.AddParagraph()
	td15.Properties().SetAlignment(wml.ST_JcLeft)
	td15f := td15.AddRun()
	td15f.Properties().SetBold(true)
	td15f.AddText("第十五条: ")
	td15i := td15.AddRun()
	td15i.Properties().SetBold(false)
	td15i.AddText("被保险人由于财产保全错误遭到被申请人起诉时，应当及时通知保险人，并应尊重并采纳保险人对诉讼的抗辩意见。未经保险人同意，被保险人不得与被申请人和解、调解结案")
	doc.AddParagraph()

	td16 := doc.AddParagraph()
	td16.Properties().SetAlignment(wml.ST_JcLeft)
	td16f := td16.AddRun()
	td16f.Properties().SetBold(true)
	td16f.AddText("第十六条: ")
	td16i := td16.AddRun()
	td16i.Properties().SetBold(false)
	td16i.AddText("被保险人应积极行使诉讼权利或履行诉讼义务，避免因怠于行使诉讼权利而承担不利的诉讼后果。")
	doc.AddParagraph()

	td17 := doc.AddParagraph()
	td17.Properties().SetAlignment(wml.ST_JcLeft)
	td17f := td17.AddRun()
	td17f.Properties().SetBold(true)
	td17f.AddText("第十七条: ")
	td17i := td17.AddRun()
	td17i.Properties().SetBold(false)
	td17i.AddText("被保险人违反上述第十四条、第十五条、第十六条或保险法规定的义务，而导致的损失，保险人应当向被申请人先行赔付，但保险人有权向被保险人追偿。")
	doc.AddParagraph()

	td18 := doc.AddParagraph()
	td18.Properties().SetAlignment(wml.ST_JcLeft)
	td18f := td18.AddRun()
	td18f.Properties().SetBold(true)
	td18f.AddText("第十八条: ")
	td18i := td18.AddRun()
	td18i.Properties().SetBold(false)
	td18i.AddText("发生保险责任范围内的损失，应由有关责任方负责赔偿的，保险人自向被保险人或被申请人赔偿保险金之日起，在赔偿金额范围内代位行使被保险人对有关责任方请求赔偿的权利，被保险人应当向保险人提供必要的文件和其所知道的有关情况。")
	doc.AddParagraph()

	td19 := doc.AddParagraph()
	td19.Properties().SetAlignment(wml.ST_JcLeft)
	td19f := td19.AddRun()
	td19f.Properties().SetBold(true)
	td19f.AddText("第十九条: ")
	td19i := td19.AddRun()
	td19i.Properties().SetBold(false)
	td19i.AddText("订立保险合同前后，投保人和被保险人应及时提供和补充提供涉及保险合同载明的民事诉讼案件的诉讼材料，包括但不限于：							（一）起诉书、立案通知书、答辩状；									（二）诉讼财产保全申请书（副本）；								（三）相关证据及鉴定文件；										（四）调解书、判决书或裁定书。")

	bczr6 := doc.AddParagraph()
	bczr6.Properties().SetAlignment(wml.ST_JcCenter)
	bczrf6 := bczr6.AddRun()
	bczrf6.Properties().SetSize(12)
	bczrf6.Properties().SetBold(true)
	bczrf6.AddText("赔偿处理")

	td20 := doc.AddParagraph()
	td20.Properties().SetAlignment(wml.ST_JcLeft)
	td20f := td20.AddRun()
	td20f.Properties().SetBold(true)
	td20f.AddText("第二十条: ")
	td20i := td20.AddRun()
	td20i.Properties().SetBold(false)
	td20i.AddText("被保险人请求赔偿时，应向保险人提供下列证明和资料：		（一）被保险人填具的索赔申请书；									（二）法院判决书、调解书或和解书及相关证据材料，如调解及和解方式结案的，需事先经保险人同意；法院判决书包括：保险合同载明的诉讼案件的法院判决书，以及因财产保全申请错误导致的诉讼案件的法院判决书;				（三）财产保全裁定书。")
	doc.AddParagraph()

	td21 := doc.AddParagraph()
	td21.Properties().SetAlignment(wml.ST_JcLeft)
	td21f := td21.AddRun()
	td21f.Properties().SetBold(true)
	td21f.AddText("第二十一条: ")
	td21i := td21.AddRun()
	td21i.Properties().SetBold(false)
	td21i.AddText("发生保险事故时，保险人应承担的赔偿金额以法院判决书或保险人认可的调解、和解书载明的赔偿金额为准，最高不超过保险单载明的保险金额。")
	doc.AddParagraph()

	td22 := doc.AddParagraph()
	td22.Properties().SetAlignment(wml.ST_JcLeft)
	td22f := td22.AddRun()
	td22f.Properties().SetBold(true)
	td22f.AddText("第二十二条: ")
	td22i := td22.AddRun()
	td22i.Properties().SetBold(false)
	td22i.AddText("保险人对被保险人因错误申请财产保全给被申请人造成的损失，可以依照法律的规定或者保险合同的约定，直接向财产保全被申请人赔偿保险金。		被保险人给财产保全被申请人造成损害，被保险人对财产保全被申请人应负的赔偿责任确定的，根据被保险人的请求，保险人应当直接向该财产保全被申请人赔偿保险金。被保险人怠于请求的，财产保全被申请人有权就其应获赔偿部分直接向保险人请求赔偿保险金。")
	doc.AddParagraph()

	td23 := doc.AddParagraph()
	td23.Properties().SetAlignment(wml.ST_JcLeft)
	td23f := td23.AddRun()
	td23f.Properties().SetBold(true)
	td23f.AddText("第二十三条: ")
	td23i := td23.AddRun()
	td23i.Properties().SetBold(false)
	td23i.AddText("未经保险人同意，被保险人不得在财产保全损害责任纠纷中自行与被申请人或利害关系人达成调解或者和解协议。")
	doc.AddParagraph()

	td24 := doc.AddParagraph()
	td24.Properties().SetAlignment(wml.ST_JcLeft)
	td24f := td24.AddRun()
	td24f.Properties().SetBold(true)
	td24f.AddText("第二十四条: ")
	td24i := td24.AddRun()
	td24i.Properties().SetBold(false)
	td24i.AddText("被保险人向保险人请求赔偿保险金的诉讼时效依法律规定，自其知道或者应当知道保险事故发生之日起计算。")

	bczr7 := doc.AddParagraph()
	bczr7.Properties().SetAlignment(wml.ST_JcCenter)
	bczrf7 := bczr7.AddRun()
	bczrf7.Properties().SetSize(12)
	bczrf7.Properties().SetBold(true)
	bczrf7.AddText("争议处理和法律适用")

	td25 := doc.AddParagraph()
	td25.Properties().SetAlignment(wml.ST_JcLeft)
	td25f := td25.AddRun()
	td25f.Properties().SetBold(true)
	td25f.AddText("第二十五条: ")
	td25i := td25.AddRun()
	td25i.Properties().SetBold(false)
	td25i.AddText(" 因履行本保险合同发生的争议，由当事人协商解决。协商不成的，提交保险单载明的仲裁机构仲裁；保险单未载明仲裁机构且争议发生后未达成仲裁协议的，依法向人民法院起诉。")
	doc.AddParagraph()

	td26 := doc.AddParagraph()
	td26.Properties().SetAlignment(wml.ST_JcLeft)
	td26f := td26.AddRun()
	td26f.Properties().SetBold(true)
	td26f.AddText("第二十六条: ")
	td26i := td26.AddRun()
	td26i.Properties().SetBold(false)
	td26i.AddText("本保险合同的争议处理适用中华人民共和国法律（不包括港澳台地区法律）。")

	bczr8 := doc.AddParagraph()
	bczr8.Properties().SetAlignment(wml.ST_JcCenter)
	bczrf8 := bczr8.AddRun()
	bczrf8.Properties().SetSize(12)
	bczrf8.Properties().SetBold(true)
	bczrf8.AddText("其他事项")

	td27 := doc.AddParagraph()
	td27.Properties().SetAlignment(wml.ST_JcLeft)
	td27f := td27.AddRun()
	td27f.Properties().SetBold(true)
	td27f.AddText("第二十七条: ")
	td27i := td27.AddRun()
	td27i.Properties().SetBold(false)
	td27i.AddText("除法院驳回申请人诉讼财产保全申请外，本保险合同投保人不得退保。")

	bczr9 := doc.AddParagraph()
	bczr9.Properties().SetAlignment(wml.ST_JcCenter)
	bczrf9 := bczr9.AddRun()
	bczrf9.Properties().SetSize(12)
	bczrf9.Properties().SetBold(true)
	bczrf9.AddText("释义")

	tde := doc.AddParagraph()
	tde.Properties().SetAlignment(wml.ST_JcCenter)
	tdef := tde.AddRun()
	tdef.AddText("【保全损害之债】是指在诉讼财产保全中被申请人针对申请人错误保全造成的损失而享有的损害请求权")

	pages := 60
	estimatedPages := len(doc.Paragraphs()) / pages
	fmt.Println(estimatedPages)
	totalParagraphs := len(doc.Paragraphs())
	endpages := totalParagraphs - (pages * estimatedPages)
	fmt.Println(endpages)
	savepath := filepath.Join(currentDir, "resource/doc/demo.docx")
	fmt.Println(savepath)

	err = doc.SaveToFile(savepath)
	if err != nil {
		return 0, "", err
	}
	return estimatedPages, pids, nil
}

type Person struct {
	Name     string
	Gender   string
	Idtype   string
	IdNumber string
}

func readexecleToPerson(filepath string) []Person {
	var persons []Person
	// 读取Excel文件
	execle, err := excelize.OpenFile(filepath)
	if err != nil {
		log.Printf("打开execl文件失败: %s", err)
	}
	defer execle.Close()
	sheet := execle.GetSheetName(0)
	rows, err := execle.GetRows(sheet)
	if err != nil {
		log.Println(err)
	}
	for rowIndex, row := range rows {
		if rowIndex == 0 {
			continue
		}
		if len(row) < 4 {
			log.Printf("第%d行数据不足，跳过此行", rowIndex+1)
			continue
		}
		name := row[0]
		gender := row[1]
		idtype := row[2]
		idnub := row[3]
		persons = append(persons, Person{
			Name:     name,
			Gender:   gender,
			Idtype:   idtype,
			IdNumber: idnub,
		})
	}
	return persons
}
