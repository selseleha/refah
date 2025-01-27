package bot_resources

import "gopkg.in/telebot.v3"

const (
	BtnBackToMainMenu = "بازگشت به منوی اصلی"
	BtnRefund         = "بازگشت"

	BtnWelfareClub        = "باشگاه رفاهی"
	BtnFunds              = "صندوق‌ها"
	BtnHealthAndTreatment = "بهداشت و درمان"
	BtnCompensation       = "جبران خدمت"
	BtnSports             = "ورزشی"
	BtnGeneralNews        = "اخبار اداره کل"
	BtnSurvey             = "نظرسنجی"

	BtnHowToReceiveCredit = "نحوه دریافت اعتبار رفاهی"
	BtnCreditBalance      = "مانده اعتبار"
	BtnContractCenters    = "مراکز طرف قرارداد"

	BtnStaffFutureFund    = "صندوق تامین آتیه کارکنان"
	BtnSupportFund        = "صندوق حمایتی"
	BtnParsFund           = "صندوق پارس"
	BtnSpecialSavingsFund = "صندوق پس انداز ویژه"

	BtnIntroductionFund = "معرفی"
	BtnSiteFund         = "آدرس سایت"
	BtnLastFileFund     = "دریافت فایل آخرین اطلاعیه"

	BtnHotels    = "هتل ها"
	BtnCompanies = "شرکت ها"

	BtnMedicalCenter = "مرکز پزشکی و دندانپزشکی بانک ملت"
	BtnJahanHotel    = "هتل جهان مشهد"
	BtnAtiyehHotel   = "هتل آتیه مشهد"

	BtnParsHotel       = "هتل پارس شیراز"
	BtnJahanAtyehHotel = "هتل جهان آتیه"

	BtnSafarMarket         = "سفرمارکت"
	BtnFedexTrip           = "فدکس تریپ"
	BtnIranHotelOnline     = "ایران هتل آنلاین"
	BtnShamsa              = "شمسآ"
	BtnReportCompanyIssues = "اعلام مشکلات شرکت ها"

	BtnIntroductionCompany = "معرفی"
	BtnConsumptionReport   = "گزارش مصرف"

	VeryLow  = "خیلی کم"
	Low      = "کم"
	Average  = "متوسط"
	Good     = "خوب"
	VeryGood = "خیلی خوب"
)

func GenerateMainMenuButtons() [][]telebot.ReplyButton {
	return [][]telebot.ReplyButton{{{Text: BtnFunds}, {Text: BtnWelfareClub}}, {{Text: BtnCompensation}, {Text: BtnHealthAndTreatment}}, {{Text: BtnGeneralNews}, {Text: BtnSports}}, {{Text: BtnSurvey}}}
}

func GenerateWelfareClubButtons() [][]telebot.ReplyButton {
	return [][]telebot.ReplyButton{
		{{Text: BtnCreditBalance}, {Text: BtnHowToReceiveCredit}},
		{{Text: BtnRefund}, {Text: BtnContractCenters}},
	}
}

func GenerateFundsButtons() [][]telebot.ReplyButton {
	return [][]telebot.ReplyButton{
		{{Text: BtnStaffFutureFund}, {Text: BtnSupportFund}},
		{{Text: BtnParsFund}, {Text: BtnSpecialSavingsFund}},
		{{Text: BtnBackToMainMenu}},
	}
}

func GenerateFundDescribeButtons() [][]telebot.ReplyButton {
	return [][]telebot.ReplyButton{
		{{Text: BtnSiteFund}, {Text: BtnIntroductionFund}},
		{{Text: BtnRefund}, {Text: BtnLastFileFund}},
		{{Text: BtnBackToMainMenu}},
	}
}

func GenerateContractCenterButtons() [][]telebot.ReplyButton {
	return [][]telebot.ReplyButton{
		{{Text: BtnCompanies}, {Text: BtnHotels}},
		{{Text: BtnRefund}},
		{{Text: BtnBackToMainMenu}},
	}
}

func GenerateSurveyButtons() [][]telebot.ReplyButton {
	return [][]telebot.ReplyButton{
		{{Text: BtnMedicalCenter}},
		{{Text: BtnJahanHotel}},
		{{Text: BtnAtiyehHotel}},
		{{Text: BtnBackToMainMenu}},
	}
}

func GenerateHotelsButtons() [][]telebot.ReplyButton {
	return [][]telebot.ReplyButton{
		{{Text: BtnJahanAtyehHotel}, {Text: BtnParsHotel}},
		{{Text: BtnRefund}},
		{{Text: BtnBackToMainMenu}},
	}
}

func GenerateSurveyInlineButtons() [][]telebot.InlineButton {
	return [][]telebot.InlineButton{
		{{Text: VeryLow, Data: "score_verylow"}},
		{{Text: Low, Data: "score_low"}},
		{{Text: Average, Data: "score_average"}},
		{{Text: Good, Data: "score_good"}},
		{{Text: VeryGood, Data: "score_verygood"}},
	}
}

func GetScoreButtons() []string {
	return []string{
		VeryLow,
		Low,
		Average,
		Good,
		VeryGood,
	}
}
func GenerateCompaniesButtons() [][]telebot.ReplyButton {
	return [][]telebot.ReplyButton{
		{{Text: BtnFedexTrip}, {Text: BtnSafarMarket}},
		{{Text: BtnShamsa}, {Text: BtnIranHotelOnline}},
		{{Text: BtnRefund}, {Text: BtnReportCompanyIssues}},
		{{Text: BtnBackToMainMenu}},
	}
}

func GenerateCompaniesDescribeButtons() [][]telebot.ReplyButton {
	return [][]telebot.ReplyButton{
		{{Text: BtnConsumptionReport}, {Text: BtnIntroductionCompany}},
		{{Text: BtnRefund}},
		{{Text: BtnBackToMainMenu}},
	}
}
