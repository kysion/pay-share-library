package pay_model

type AlipayTradeSource struct {
	AppID          string `json:"app_id" dc:"支付宝分配给开发者的应用ID"`
	AuthAppID      string `json:"auth_app_id" dc:"认证AppId，一般是商家AppID"`
	BuyerID        string `json:"buyer_id" dc:"买家在支付宝的userId"`
	BuyerLogonID   string `json:"buyer_logon_id" dc:"买家登录id"`
	BuyerPayAmount string `json:"buyer_pay_amount" dc:"买家支付金额"`
	Charset        string `json:"charset" dc:"编码"`
	FundBillList   string `json:"fund_bill_list" dc:"找到的账单集合"`
	GmtCreate      string `json:"gmt_create" dc:""`
	GmtPayment     string `json:"gmt_payment" dc:""`
	InvoiceAmount  string `json:"invoice_amount" dc:"发票金额"`
	NotifyID       string `json:"notify_id" dc:"通知id"`
	NotifyTime     string `json:"notify_time" dc:"通知时间"`
	NotifyType     string `json:"notify_type" dc:"通知类型"`
	OutTradeNo     string `json:"out_trade_no" dc:"交易订单号"`
	PassbackParams string `json:"passback_params" dc:"公用回传参数"`
	PointAmount    string `json:"point_amount" dc:""`
	ReceiptAmount  string `json:"receipt_amount" dc:"发票金额"`
	SellerEmail    string `json:"seller_email" dc:"卖方联系邮箱"`
	SellerID       string `json:"seller_id" dc:"卖方支付宝的userId"`
	Sign           string `json:"sign" dc:"签名"`
	SignType       string `json:"sign_type" dc:"签名类型"`
	Subject        string `json:"subject" dc:"产品名称"`
	TotalAmount    string `json:"total_amount" dc:"总金额"`
	TradeNo        string `json:"trade_no" dc:"交易编号"`
	TradeStatus    string `json:"trade_status" dc:"交易状态"`
	Version        string `json:"version" dc:"版本"`
}
