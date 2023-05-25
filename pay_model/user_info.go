package pay_model

type UserInfo struct {
	// 支付宝独有
	UserId             string `json:"user_id,omitempty"  dc:"支付宝唯一标识user_id"`
	Avatar             string `json:"avatar,omitempty"  dc:"头像"`
	Province           string `json:"province,omitempty"  dc:"省份"`
	City               string `json:"city,omitempty"  dc:"城市"`
	NickName           string `json:"nick_name,omitempty"  dc:"昵称"`
	IsStudentCertified string `json:"is_student_certified,omitempty"  dc:"是否学生"`
	UserType           string `json:"user_type,omitempty"  dc:"用户类型"`
	UserStatus         string `json:"user_status,omitempty"  dc:"用户状态"`
	IsCertified        string `json:"is_certified,omitempty"  dc:"是否实名认证"`
	Gender             string `json:"gender,omitempty"  dc:"性别"`

	// 微信独有

	OpenID     string `json:"openid" dc:"用户openId"`
	UnionID    string `json:"unionid" dc:"用户unionId"`
	Sex        int    `json:"sex" dc:"性别"`
	Country    string `json:"country" dc:"城镇"`
	HeadImgURL string `json:"headimgurl" dc:"头像"`
	ErrCode    int    `json:"errcode,omitempty" dc:""`
	ErrMsg     string `json:"errmsg,omitempty" dc:""`

	// 筷满客
	SysUserId int64 `json:"sys_user_id" dc:"筷满客平台用户id"`
}
