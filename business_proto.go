package openkcb

type OpenAccountReq struct {
	CustomerType     string `json:"customerType,omitempty"`
	CustomerIdType   string `json:"customerIdType,omitempty"`
	CurrencyCode     string `json:"currencyCode,omitempty"`
	CustomerIdNo     string `json:"customerIdNo,omitempty"`
	CustomerName     string `json:"customerName,omitempty"`
	PlatformMemberNo string `json:"platformMemberNo,omitempty"`
	MemberType       int    `json:"memberType,omitempty"`
	ContactsPhone    string `json:"contactsPhone,omitempty"`
	ContactsEmail    string `json:"contactsEmail,omitempty"`
}

type OpenAccountRsp struct {
	TraceId      string `json:"traceId,omitempty"`
	OpenDate     string `json:"openDate,omitempty"`
	AccountName  string `json:"accountName,omitempty"`
	AccountNo    string `json:"accountNo,omitempty"`
	AccountType  string `json:"accountType,omitempty"`
	CurrencyCode string `json:"currencyCode,omitempty"`
}

func (OpenAccountReq) Validate() error {
	return nil
}

type QueryAccountReq struct {
	AccountNo        string `json:"accountNo,omitempty"`
	PlatformMemberNo string `json:"platformMemberNo,omitempty"`
}

func (QueryAccountReq) Validate() error {
	return nil
}

type AccountInfo struct {
	AccountName         string          `json:"accountName,omitempty"`
	AccountNo           string          `json:"accountNo,omitempty"`
	AccountType         string          `json:"accountType,omitempty"`
	AvailableBalance    float64         `json:"availableBalance,omitempty"`
	BindCardList        []*BindCardInfo `json:"bindCardList,omitempty"`
	CurrencyCode        string          `json:"currencyCode,omitempty"`
	FreezeBalance       float64         `json:"freezeBalance,omitempty"`
	MemberType          int             `json:"memberType,omitempty"`
	OpenDate            int64           `json:"openDate,omitempty"`
	Status              int             `json:"status,omitempty"`
	TransitBalance      float64         `json:"transitBalance,omitempty"`
	WithdrawableBalance float64         `json:"withdrawableBalance,omitempty"`
}

type QueryAccountRsp struct {
	AccountInfo
}

type QueryAccountByMemberNoReq struct {
	PlatformMemberNo string `json:"platformMemberNo,omitempty"`
}

func (QueryAccountByMemberNoReq) Validate() error {
	return nil
}

type QueryAccountByMemberNoRsp []*AccountInfo

//type QueryAccountByMemberNoRsp struct {
//	AccountInfo
//}

type BindCardInfo struct {
}

type BindSmallAmountReq struct {
	PlatformMemberNo string `json:"platformMemberNo,omitempty"`
	AccountNo        string `json:"accountNo,omitempty"`
	BankCardType     string `json:"bankCardType,omitempty"`
	BankCode         string `json:"bankCode,omitempty"`
	BankName         string `json:"bankName,omitempty"`
	BankAccountNo    string `json:"bankAccountNo,omitempty"`
	BankAccountName  string `json:"bankAccountName,omitempty"`
	IdType           string `json:"idType,omitempty"`
	IdNo             string `json:"idNo,omitempty"`
	Mobile           string `json:"mobile,omitempty"`
}

func (BindSmallAmountReq) Validate() error {
	return nil
}

type BindSmallAmountRsp struct {
	TradeId   string `json:"tradeId,omitempty"`
	AccountNo string `json:"accountNo,omitempty"`
	Remark    string `json:"remark,omitempty"`
}

type QuerySmallAmountReq struct {
	PlatformMemberNo string `json:"platformMemberNo,omitempty"`
	TradeId          string `json:"tradeId,omitempty"`
	AccountNo        string `json:"accountNo,omitempty"`
}

func (QuerySmallAmountReq) Validate() error {
	return nil
}

type QuerySmallAmountRsp struct {
	TradeStatus int    `json:"tradeStatus,omitempty"`
	AccountNo   string `json:"accountNo,omitempty"`
	Remark      string `json:"remark,omitempty"`
}

type BindUnionReq struct {
	PlatformMemberNo string `json:"platformMemberNo,omitempty"`
	AccountNo        string `json:"accountNo,omitempty"`
	BankCardType     string `json:"bankCardType,omitempty"`
	BankCode         string `json:"bankCode,omitempty"`
	BankName         string `json:"bankName,omitempty"`
	BankAccountNo    string `json:"bankAccountNo,omitempty"`
	BankAccountName  string `json:"bankAccountName,omitempty"`
	IdType           string `json:"idType,omitempty"`
	IdNo             string `json:"idNo,omitempty"`
	Mobile           string `json:"mobile,omitempty"`
}

func (BindUnionReq) Validate() error {
	return nil
}

type BindUnionRsp struct {
	TradeId   string `json:"tradeId,omitempty"`
	AccountNo string `json:"accountNo,omitempty"`
	Remark    string `json:"remark,omitempty"`
}

type ValidateAmountReq struct {
	PlatformMemberNo string `json:"platformMemberNo,omitempty"`
	AccountNo        string `json:"accountNo,omitempty"`
	TradeId          string `json:"tradeId,omitempty"`
	Amount           string `json:"amount,omitempty"`
}

func (ValidateAmountReq) Validate() error {
	return nil
}

type ValidateAmountRsp struct {
	TradeId         string `json:"tradeId,omitempty"`
	Status          int    `json:"status,omitempty"`
	BindDate        string `json:"bindDate,omitempty"`
	Remark          string `json:"remark,omitempty"`
	BankCardType    string `json:"bankCardType,omitempty"`
	BankCode        string `json:"bankCode,omitempty"`
	BankName        string `json:"bankName,omitempty"`
	BankAccountNo   string `json:"bankAccountNo,omitempty"`
	BankAccountName string `json:"bankAccountName,omitempty"`
}

type ValidateCodeReq struct {
	PlatformMemberNo string `json:"platformMemberNo,omitempty"`
	AccountNo        string `json:"accountNo,omitempty"`
	SmsCode          string `json:"smsCode,omitempty"`
	TradeId          string `json:"tradeId,omitempty"`
	Mobile           string `json:"mobile,omitempty"`
}

func (ValidateCodeReq) Validate() error {
	return nil
}

type ValidateCodeRsp struct {
	TradeId         string `json:"tradeId,omitempty"`
	Status          int    `json:"status,omitempty"`
	BindDate        string `json:"bindDate,omitempty"`
	Remark          string `json:"remark,omitempty"`
	BankCardType    string `json:"bankCardType,omitempty"`
	BankCode        string `json:"bankCode,omitempty"`
	BankName        string `json:"bankName,omitempty"`
	BankAccountNo   string `json:"bankAccountNo,omitempty"`
	BankAccountName string `json:"bankAccountName,omitempty"`
}

type RegisterBillReq struct {
	PlatformMemberNo string `json:"platformMemberNo,omitempty"`
	AccountNo        string `json:"accountNo,omitempty"`
	TradeType        int    `json:"tradeType,omitempty"`
	TradeAmount      string `json:"tradeAmount,omitempty"`
	Fee              string `json:"fee,omitempty"`
	Remark           string `json:"remark,omitempty"`
}

func (RegisterBillReq) Validate() error {
	return nil
}

type RegisterBillRsp struct {
	TradeId         string `json:"tradeId,omitempty"`
	PlatformCode    string `json:"platformCode,omitempty"`
	PlatformName    string `json:"platformName,omitempty"`
	PlatformTradeId string `json:"platformTradeId,omitempty"`
}

type RevokeRegisterBillReq struct {
	PlatformMemberNo string `json:"platformMemberNo,omitempty"`
	AccountNo        string `json:"accountNo,omitempty"`
	RefundType       int    `json:"refundType,omitempty"`
	OrigPlatTradeId  string `json:"origPlatTradeId,omitempty"`
	RefundAmount     string `json:"refundAmount,omitempty"`
	RefundFee        string `json:"refundFee,omitempty"`
	RefundDesc       string `json:"refundDesc,omitempty"`
	Remark           string `json:"remark,omitempty"`
	RefundPolicy     string `json:"refundPolicy,omitempty"`
}

func (RevokeRegisterBillReq) Validate() error {
	return nil
}

type RevokeRegisterBillRsp struct {
	RefundId         string `json:"refundId,omitempty"`
	PlatformCode     string `json:"platformCode,omitempty"`
	PlatformName     string `json:"platformName,omitempty"`
	PlatformRefundId string `json:"platformRefundId,omitempty"`
}

type QueryTradeListReq struct {
	PlatformMemberNo string `json:"platformMemberNo,omitempty"`
	AccountNo        string `json:"accountNo,omitempty"`
	TradeType        int    `json:"tradeType,omitempty"`
	TradeStatus      int    `json:"tradeStatus,omitempty"`
	PlatformTradeId  string `json:"platformTradeId,omitempty"`
	TradeId          string `json:"tradeId,omitempty"`
	TradeTimeStart   string `json:"tradeTimeStart,omitempty"`
	TradeTimeEnd     string `json:"tradeTimeEnd,omitempty"`
	PageSize         string `json:"pageSize,omitempty"`
	PageNum          string `json:"pageNum,omitempty"`
}

func (QueryTradeListReq) Validate() error {
	return nil
}

type TradeInfo struct {
	AccountNo       string  `json:"accountNo,omitempty"`
	PlatformTradeId string  `json:"platformTradeId,omitempty"`
	TradeAmount     float64 `json:"tradeAmount,omitempty"`
	TradeFee        float64 `json:"tradeFee,omitempty"`
	TradeId         string  `json:"tradeId,omitempty"`
	TradeStatus     int     `json:"tradeStatus,omitempty"`
	TradeType       int     `json:"tradeType,omitempty"`
}

type QueryTradeListRsp struct {
	PageNum    int          `json:"pageNum,omitempty"`
	PageSize   int          `json:"pageSize,omitempty"`
	Size       int          `json:"size,omitempty"`
	TotalCount int          `json:"totalCount,omitempty"`
	TradeList  []*TradeInfo `json:"tradeList,omitempty"`
}

type QueryTradeReq struct {
	PlatformMemberNo string `json:"platformMemberNo,omitempty"`
	AccountNo        string `json:"accountNo,omitempty"`
	TradeType        int    `json:"tradeType,omitempty"`
	PlatformTradeId  string `json:"platformTradeId,omitempty"`
	TradeId          string `json:"tradeId,omitempty"`
}

func (QueryTradeReq) Validate() error {
	return nil
}

type QueryTradeRsp struct {
	AccountNo       string  `json:"accountNo,omitempty"`
	TradeType       int     `json:"tradeType,omitempty"`
	TradeId         string  `json:"tradeId,omitempty"`
	PlatformTradeId string  `json:"platformTradeId,omitempty"`
	Status          int     `json:"status,omitempty"`
	Amount          float64 `json:"amount,omitempty"`
	Fee             float64 `json:"fee,omitempty"`
	TradeTime       int64   `json:"tradeTime,omitempty"`
	RefundFlag      int     `json:"refundFlag,omitempty"`
}

type SecurePrepayReq struct {
	TradeType          int    `json:"tradeType,omitempty"`
	TradeAmount        string `json:"tradeAmount,omitempty"`
	SettleAmount       string `json:"settleAmount,omitempty"`
	Fee                string `json:"fee,omitempty"`
	ProductName        string `json:"productName,omitempty"`
	Quantity           int    `json:"quantity,omitempty"`
	PlatformMemberNo   string `json:"platformMemberNo,omitempty"`
	AccountNo          string `json:"accountNo,omitempty"`
	PlatformMerchantNo string `json:"platformMerchantNo,omitempty"`
	MAccountNo         string `json:"mAccountNo,omitempty"`
}

func (SecurePrepayReq) Validate() error {
	return nil
}

type SecurePrepayRsp struct {
	TradeId         string `json:"tradeId,omitempty"`
	PlatformCode    string `json:"platformCode,omitempty"`
	PlatformName    string `json:"platformName,omitempty"`
	PlatformTradeId string `json:"platformTradeId,omitempty"`
}

type SecureConfirmPayReq struct {
	PlatformMemberNo   string `json:"platformMemberNo,omitempty"`
	AccountNo          string `json:"accountNo,omitempty"`
	PlatformMerchantNo string `json:"platformMerchantNo,omitempty"`
	MAccountNo         string `json:"mAccountNo,omitempty"`
	TradeType          int    `json:"tradeType,omitempty"`
	PlatformTradeId    string `json:"platformTradeId,omitempty"`
	SettleAmount       string `json:"settleAmount,omitempty"`
	Fee                string `json:"fee,omitempty"`
}

func (SecureConfirmPayReq) Validate() error {
	return nil
}

type SecureConfirmPayRsp struct {
	TradeId         string `json:"tradeId,omitempty"`
	PlatformCode    string `json:"platformCode,omitempty"`
	PlatformName    string `json:"platformName,omitempty"`
	PlatformTradeId string `json:"platformTradeId,omitempty"`
}

type SecurePayRefundReq struct {
	PlatformMemberNo string `json:"platformMemberNo,omitempty"`
	AccountNo        string `json:"accountNo,omitempty"`
	RefundType       int    `json:"refundType,omitempty"`
	OrigPlatTradeId  string `json:"origPlatTradeId,omitempty"`
	RefundAmount     string `json:"refundAmount,omitempty"`
	RefundFee        string `json:"refundFee,omitempty"`
	RefundDesc       string `json:"refundDesc,omitempty"`
	Remark           string `json:"remark,omitempty"`
	RefundPolicy     string `json:"refundPolicy,omitempty"`
}

func (SecurePayRefundReq) Validate() error {
	return nil
}

type SecurePayRefundRsp struct {
	RefundId         string `json:"refundId,omitempty"`
	PlatformCode     string `json:"platformCode,omitempty"`
	PlatformName     string `json:"platformName,omitempty"`
	PlatformRefundId string `json:"platformRefundId,omitempty"`
}

type PayReq struct {
	TradeType          int    `json:"tradeType,omitempty"`
	TradeAmount        string `json:"tradeAmount,omitempty"`
	Fee                string `json:"fee,omitempty"`
	ProductName        string `json:"productName,omitempty"`
	Quantity           int    `json:"quantity,omitempty"`
	PlatformMemberNo   string `json:"platformMemberNo,omitempty"`
	AccountNo          string `json:"accountNo,omitempty"`
	PlatformMerchantNo string `json:"platformMerchantNo,omitempty"`
	MAccountNo         string `json:"mAccountNo,omitempty"`
}

func (PayReq) Validate() error {
	return nil
}

type PayRsp struct {
	TradeId         string `json:"tradeId,omitempty"`
	PlatformCode    string `json:"platformCode,omitempty"`
	PlatformName    string `json:"platformName,omitempty"`
	PlatformTradeId string `json:"platformTradeId,omitempty"`
}

type RefundReq struct {
	PlatformMemberNo   string `json:"platformMemberNo,omitempty"`
	AccountNo          string `json:"accountNo,omitempty"`
	RefundType         int    `json:"refundType,omitempty"`
	OrigPlatTradeId    string `json:"origPlatTradeId,omitempty"`
	RefundAmount       string `json:"refundAmount,omitempty"`
	ActualRefundAmount string `json:"actualRefundAmount,omitempty"`
	RefundFee          string `json:"refundFee,omitempty"`
	RefundDesc         string `json:"refundDesc,omitempty"`
	Remark             string `json:"remark,omitempty"`
	RefundPolicy       string `json:"refundPolicy,omitempty"`
}

func (RefundReq) Validate() error {
	return nil
}

type RefundRsp struct {
	RefundId         string `json:"refundId,omitempty"`
	PlatformCode     string `json:"platformCode,omitempty"`
	PlatformName     string `json:"platformName,omitempty"`
	PlatformRefundId string `json:"platformRefundId,omitempty"`
}

type FreezeReq struct {
	PlatformMemberNo string `json:"platformMemberNo,omitempty"`
	AccountNo        string `json:"accountNo,omitempty"`
	Amount           string `json:"amount,omitempty"`
	Remark           string `json:"remark,omitempty"`
}

func (FreezeReq) Validate() error {
	return nil
}

type FreezeRsp struct {
	TradeId   string `json:"tradeId,omitempty"`
	AccountNo string `json:"accountNo,omitempty"`
	Remark    string `json:"remark,omitempty"`
}

type UnFreezeReq struct {
	PlatformMemberNo string `json:"platformMemberNo,omitempty"`
	AccountNo        string `json:"accountNo,omitempty"`
	FreezeTradeId    string `json:"freezeTradeId,omitempty"`
	Amount           string `json:"amount,omitempty"`
	Remark           string `json:"remark,omitempty"`
}

func (UnFreezeReq) Validate() error {
	return nil
}

type UnFreezeRsp struct {
	TradeId   string `json:"tradeId,omitempty"`
	AccountNo string `json:"accountNo,omitempty"`
	Remark    string `json:"remark,omitempty"`
}

type WithdrawReq struct {
	PlatformMemberNo string `json:"platformMemberNo,omitempty"`
	AccountNo        string `json:"accountNo,omitempty"`
	TradeType        int    `json:"tradeType,omitempty"`
	TradeAmount      string `json:"tradeAmount,omitempty"`
	Fee              string `json:"fee,omitempty"`
	ReceAcctNo       string `json:"receAcctNo,omitempty"`
	ReceAcctType     string `json:"receAcctType,omitempty"`
	ReceAcctName     string `json:"receAcctName,omitempty"`
	Telephone        string `json:"telephone,omitempty"`
	AcctOrgan        string `json:"acctOrgan,omitempty"`
	InterBankNo      string `json:"interBankNo,omitempty"`
}

func (WithdrawReq) Validate() error {
	return nil
}

type WithdrawRsp struct {
	TradeId         string `json:"tradeId,omitempty"`
	PlatformCode    string `json:"platformCode,omitempty"`
	PlatformName    string `json:"platformName,omitempty"`
	PlatformTradeId string `json:"platformTradeId,omitempty"`
}

type BatchPayReq struct {
	PlatParentTradeId  string `json:"platParentTradeId,omitempty"`
	AccountNo          string `json:"accountNo,omitempty"`
	PlatformMemberNo   string `json:"platformMemberNo,omitempty"`
	PayModel           string `json:"payModel,omitempty"`
	TradeType          int    `json:"tradeType,omitempty"`
	TradeAmount        string `json:"tradeAmount,omitempty"`
	Fee                string `json:"fee,omitempty"`
	Remark             string `json:"remark,omitempty"`
	MerchantPayReqList []*BatchPayItem
}

type BatchPayItem struct {
	SettleAmount       string `json:"settleAmount,omitempty"`
	SettleFee          string `json:"settleFee,omitempty"`
	ProductName        string `json:"productName,omitempty"`
	Quantity           int    `json:"quantity,omitempty"`
	PlatformTradeId    string `json:"platformTradeId,omitempty"`
	PlatformMerchantNo string `json:"platformMerchantNo,omitempty"`
	MAccountNo         string `json:"mAccountNo,omitempty"`
}

func (BatchPayReq) Validate() error {
	return nil
}

type BatchPayRsp []BatchPayRspItem
type BatchPayRspItem struct {
	TradeId         string `json:"tradeId,omitempty"`
	PlatformCode    string `json:"platformCode,omitempty"`
	PlatformName    string `json:"platformName,omitempty"`
	PlatformTradeId string `json:"platformTradeId,omitempty"`
}

type MemberRechargeReq struct {
	PlatformMemberNo    string `json:"platformMemberNo,omitempty"`
	AccountNo           string `json:"accountNo,omitempty"`
	TradeType           int    `json:"tradeType,omitempty"`
	TradeAmount         string `json:"tradeAmount,omitempty"`
	Fee                 string `json:"fee,omitempty"`
	PayChannelTranSeqNo string `json:"payChannelTranSeqNo,omitempty"`
	PlatChargeOrderId   string `json:"platChargeOrderId,omitempty"`
	PlatformTradeId     string `json:"platformTradeId,omitempty"`
}

func (MemberRechargeReq) Validate() error {
	return nil
}

type MemberRechargeRsp struct {
	TradeId         string `json:"tradeId,omitempty"`
	PlatformCode    string `json:"platformCode,omitempty"`
	PlatformName    string `json:"platformName,omitempty"`
	PlatformTradeId string `json:"platformTradeId,omitempty"`
}

type RevokeMemberRechargeReq struct {
	PlatformMemberNo string `json:"platformMemberNo,omitempty"`
	AccountNo        string `json:"accountNo,omitempty"`
	RefundType       int    `json:"refundType,omitempty"`
	OrigPlatTradeId  string `json:"origPlatTradeId,omitempty"`
	RefundAmount     string `json:"refundAmount,omitempty"`
	RefundFee        string `json:"refundFee,omitempty"`
	RefundDesc       string `json:"refundDesc,omitempty"`
	Remark           string `json:"remark,omitempty"`
	RefundPolicy     string `json:"refundPolicy,omitempty"`
}

func (RevokeMemberRechargeReq) Validate() error {
	return nil
}

type RevokeMemberRechargeRsp struct {
	RefundId         string `json:"refundId,omitempty"`
	PlatformCode     string `json:"platformCode,omitempty"`
	PlatformName     string `json:"platformName,omitempty"`
	PlatformRefundId string `json:"platformRefundId,omitempty"`
}
