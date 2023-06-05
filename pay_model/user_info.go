package pay_model

import "github.com/gogf/gf/v2/os/gtime"

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

	// 抖音独有
	SessionKey string `json:"session_key" dc:"抖音用户session_key"`

	AccessToken  string      `json:"accessToken"        description:"授权token"`
	RefreshToken string      `json:"refreshToken"       description:"抖音用户授权刷新Token"`
	ExpiresIn    *gtime.Time `json:"expiresIn"          description:"令牌过期时间"`

	Data struct {
		Avatar       string `json:"avatar"`
		AvatarLarger string `json:"avatar_larger"`
		ClientKey    string `json:"client_key"`
		EAccountRole string `json:"e_account_role"`
		ErrorCode    int    `json:"error_code"`
		LogId        string `json:"log_id"`
		Nickname     string `json:"nickname"`

		OpenId  string `json:"open_id"`
		UnionId string `json:"union_id"`
	} `json:"data"`

	// 筷满客
	SysUserId int64 `json:"sys_user_id" dc:"筷满客平台用户id"`
}
