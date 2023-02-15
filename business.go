package openkcb

import (
	"context"
)

func (c *Client) OpenAccount(ctx context.Context, req *OpenAccountReq) (*OpenAccountRsp, error) {
	return Invoke[OpenAccountReq, OpenAccountRsp](c, ctx, "AC010101", req, "")
}

func (c *Client) QueryAccount(ctx context.Context, req *QueryAccountReq) (*QueryAccountRsp, error) {
	return Invoke[QueryAccountReq, QueryAccountRsp](c, ctx, "QY170101", req, "")
}

func (c *Client) QueryAccountByMemberNo(ctx context.Context, req *QueryAccountByMemberNoReq) (*QueryAccountByMemberNoRsp, error) {
	return Invoke[QueryAccountByMemberNoReq, QueryAccountByMemberNoRsp](c, ctx, "QY170201", req, "")
}

func (c *Client) BindSmallAmount(ctx context.Context, req *BindSmallAmountReq) (*BindSmallAmountRsp, error) {
	return Invoke[BindSmallAmountReq, BindSmallAmountRsp](c, ctx, "AC150301", req, "")
}

func (c *Client) ValidateAmount(ctx context.Context, req *ValidateAmountReq) (*ValidateAmountRsp, error) {
	return Invoke[ValidateAmountReq, ValidateAmountRsp](c, ctx, "AC150401", req, "")
}

func (c *Client) QuerySmallAmount(ctx context.Context, req *QuerySmallAmountReq) (*QuerySmallAmountRsp, error) {
	return Invoke[QuerySmallAmountReq, QuerySmallAmountRsp](c, ctx, "AC150501", req, "")
}

func (c *Client) BindUnion(ctx context.Context, req *BindUnionReq) (*BindUnionRsp, error) {
	return Invoke[BindUnionReq, BindUnionRsp](c, ctx, "AC150101", req, "")
}

func (c *Client) ValidateCode(ctx context.Context, req *ValidateCodeReq) (*ValidateCodeRsp, error) {
	return Invoke[ValidateCodeReq, ValidateCodeRsp](c, ctx, "AC150201", req, "")
}

func (c *Client) RegisterBill(ctx context.Context, req *RegisterBillReq, bizNo string) (*RegisterBillRsp, error) {
	return Invoke[RegisterBillReq, RegisterBillRsp](c, ctx, "TD110101", req, bizNo)
}

func (c *Client) RevokeRegisterBill(ctx context.Context, req *RevokeRegisterBillReq) (*RevokeRegisterBillRsp, error) {
	return Invoke[RevokeRegisterBillReq, RevokeRegisterBillRsp](c, ctx, "TD120101", req, "")
}

func (c *Client) QueryTradeList(ctx context.Context, req *QueryTradeListReq) (*QueryTradeListRsp, error) {
	return Invoke[QueryTradeListReq, QueryTradeListRsp](c, ctx, "QY170301", req, "")
}

func (c *Client) QueryTrade(ctx context.Context, req *QueryTradeReq) (*QueryTradeRsp, error) {
	return Invoke[QueryTradeReq, QueryTradeRsp](c, ctx, "QY170401", req, "")
}

func (c *Client) SecurePrepay(ctx context.Context, req *SecurePrepayReq, bizNo string) (*SecurePrepayRsp, error) {
	return Invoke[SecurePrepayReq, SecurePrepayRsp](c, ctx, "TD050101", req, bizNo)
}

func (c *Client) SecureConfirmPay(ctx context.Context, req *SecureConfirmPayReq) (*SecureConfirmPayRsp, error) {
	return Invoke[SecureConfirmPayReq, SecureConfirmPayRsp](c, ctx, "TD090101", req, "")
}

func (c *Client) SecurePayRefund(ctx context.Context, req *SecurePayRefundReq) (*SecurePayRefundRsp, error) {
	return Invoke[SecurePayRefundReq, SecurePayRefundRsp](c, ctx, "TD060101", req, "")
}

func (c *Client) Pay(ctx context.Context, req *PayReq, bizNo string) (*PayRsp, error) {
	return Invoke[PayReq, PayRsp](c, ctx, "TD050101", req, bizNo)
}

func (c *Client) Refund(ctx context.Context, req *RefundReq) (*RefundRsp, error) {
	return Invoke[RefundReq, RefundRsp](c, ctx, "TD060101", req, "")
}

func (c *Client) Freeze(ctx context.Context, req *FreezeReq) (*FreezeRsp, error) {
	return Invoke[FreezeReq, FreezeRsp](c, ctx, "TD030101", req, "")
}

func (c *Client) UnFreeze(ctx context.Context, req *UnFreezeReq) (*UnFreezeRsp, error) {
	return Invoke[UnFreezeReq, UnFreezeRsp](c, ctx, "TD040101", req, "")
}

func (c *Client) Withdraw(ctx context.Context, req *WithdrawReq) (*WithdrawRsp, error) {
	return Invoke[WithdrawReq, WithdrawRsp](c, ctx, "TD170101", req, "")
}

func (c *Client) BatchPay(ctx context.Context, req *BatchPayReq) (*BatchPayRsp, error) {
	return Invoke[BatchPayReq, BatchPayRsp](c, ctx, "TD180101", req, "")
}

// MemberRecharge 该接口暂时不需要开发
func (c *Client) MemberRecharge(ctx context.Context, req *MemberRechargeReq, bizNo string) (*MemberRechargeRsp, error) {
	return Invoke[MemberRechargeReq, MemberRechargeRsp](c, ctx, "TD020101", req, bizNo)
}

// RevokeMemberRecharge 该接口暂时不需要开发
func (c *Client) RevokeMemberRecharge(ctx context.Context, req *RevokeMemberRechargeReq) (*RevokeMemberRechargeRsp, error) {
	return Invoke[RevokeMemberRechargeReq, RevokeMemberRechargeRsp](c, ctx, "TD030102", req, "")
}
