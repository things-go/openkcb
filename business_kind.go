package openkcb

const (
	//  账户类型
	CustomerType_PERSONAL   = "60" // 个人
	CustomerType_ENTERPRISE = "80" // 企业
	// 客户证件类型
	CustomerIdType_IDENTITY_CARD    = "1" // 身份证
	CustomerIdType_BUSINESS_LICENSE = "2" // 营业执照/统一社会信用代码
	// 货币代码
	CurrencyCode_CNY = "CNY"
	// 会员类型
	MemberType_PURCHASER = 1 // 买家
	MemberType_SELLER    = 2 // 卖家

	// 银行卡类型
	BankCardType_DEBIT_CARD      = "1" // 借记卡
	BankCardType_LOAD_CARD       = "2" // 贷记卡
	BankCardType_QUASI_LOAD_CARD = "3" // 准贷记卡
	BankCardType_CASH_CARD       = "4" // 储蓄卡
)
