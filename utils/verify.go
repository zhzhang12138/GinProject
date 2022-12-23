package utils

var (
	IdVerify        = Rules{"ID": []string{NotEmpty()}}
	LoginVerify     = Rules{"CaptchaId": {NotEmpty()}, "Captcha": {NotEmpty()}, "Username": {NotEmpty()}, "Password": {NotEmpty()}}
	AuthorityVerify = Rules{"AuthorityId": {NotEmpty()}, "AuthorityName": {NotEmpty()}}
	RegisterVerify  = Rules{"Username": {NotEmpty()}, "NickName": {NotEmpty()}, "Password": {NotEmpty()}, "AuthorityId": {NotEmpty()}}
	PageInfoVerify  = Rules{"Page": {NotEmpty()}, "PageSize": {NotEmpty()}}
)
