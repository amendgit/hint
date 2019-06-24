卡管家login
```json
{
	"bankId": 3,
	"password": "zXFAC9pIYOTs2NVcpmxsLAGUguBb2eUhhkhzhCwftG7J4ezL3qsmkgVHsitRA\/JJHwGkvpxcFEvNRb\/0wnPPv+87cbJKJI5gLpHoQqyH4RKV9UcwPpio4WbLR7IWsnnFXZzYWAtIWw75+GRKtLdzPk8pYHaa9OTCE4wkgKwNEJrFqNRL4pJw6UpTzfYaqio+0dVh7siYs\/15W0h4HN50PxrB1ASsZAPFCEpIFhifqh5RlmESf1jVbKMq6jd99bB9krEMdtpx42bWVklfKcVn\/p33Wg49VDX6tmpBys7SQq4Ms72myFQCtoZcCG0gzOuzFvpSTJqaq41Rb2GUp5j+ag==",
	"isCreditCard": false,
	"publicKeyId": "055327b90b4b87f4b8f0badf25ff1517",
	"identification": {
		"isBreak": false,
		"pushToken": "",
		"idfv": "C420D9E0-E18E-44D2-9CFA-B0E61C5F5D0E",
		"devType": 1,
		"macAddress": "",
		"version": "12.1",
		"routeMac": "",
		"networkType": 2,
		"idfa": "3E45B6A0-9496-4F58-A212-50F33F6096B3"
	},
	"loanProductCode": "",
	"isSavePwd": true,
	"accType": 1,
	"accessName": "CMBCC",
	"entryName": "6225768769098482"
}
```

卡管家loginByEntry
```json
{
	"entryId": 12954303,
	"publicKeyId": "055327b90b4b87f4b8f0badf25ff1517",
	"isCreditCard": false,
	"password": "",
	"isSavePwd": true,
	"identification": {
		"isBreak": false,
		"pushToken": "",
		"idfv": "C420D9E0-E18E-44D2-9CFA-B0E61C5F5D0E",
		"devType": 1,
		"macAddress": "",
		"version": "12.1",
		"routeMac": "",
		"networkType": 2,
		"idfa": "3E45B6A0-9496-4F58-A212-50F33F6096B3"
	}
}
```

点石网银loginByEntry
```json
{
	"bankId": 3,
	"password": "123456",
	"isCreditCard": true,
	"publicKeyId": "055327b90b4b87f4b8f0badf25ff1517",
	"isSavePwd": true,
	"accType": 4,
	"entryName": "18636130636"
}
```

```bash
curl -H 'Host: crawler.credit-manager.k2.wacaiyun.com' -H 'Accept: */*' -H 'Accept-Language: en-us' -H 'X-Mc: B0000985' -H 'Content-Type: application/json' -H 'User-Agent: EbankImport_Example/1.0 CFNetwork/975.0.3 Darwin/18.2.0' -H 'X-Deviceid: 95100472-B711-45B0-B322-260FF973F055' -H 'X-Appver: 1.0.0' -H 'X-Platform: 73' -H 'X-Access-Token: 5lm4NBKCVeVxRg7YxLfw5UOgLty8GK1p97Q' --data-binary '{"entryId":12954303,"password":"","isCreditCard":false,"publicKeyId":"055327b90b4b87f4b8f0badf25ff1517","isSavePwd":true,"entryName":"18636130636"}' --compressed 'http://crawler.credit-manager.k2.wacaiyun.com/crawler/ebank/1.2/loginByEntry'
```

成功的那个
```bash
curl -H 'Host: crawler.credit-manager.k2.wacaiyun.com' -H 'Accept: */*' -H 'Accept-Language: en-us' -H 'X-Mc: B0000985' -H 'Content-Type: application/json' -H 'User-Agent: EbankImport_Example/1.0 CFNetwork/975.0.3 Darwin/18.2.0' -H 'X-Deviceid: 95100472-B711-45B0-B322-260FF973F055' -H 'X-Appver: 1.0.0' -H 'X-Platform: 73' -H 'X-Access-Token: 5lm4NBKCVeVxRg7YxLfw5UOgLty8GK1p97Q' --data-binary '{"entryId":12954303,"password":"","isCreditCard":false,"publicKeyId":"055327b90b4b87f4b8f0badf25ff1517","isSavePwd":true}' --compressed 'http://crawler.credit-manager.k2.wacaiyun.com/crawler/ebank/1.2/loginByEntry'
```

```bash
curl -H 'Host: crawler.credit-manager.k2.wacaiyun.com' -H 'Content-Type: application/json' -H 'Accept: */*' -H 'X-Md: 7E526991B3044442859E154DA0218F9F' -H 'Accept-Language: en;q=1' -H 'X-Mc: 21000020' -H 'X-Access-Sign: 35655872236583d8816b1ffd21194f0c76c98f0f' -H 'X-Trace-Id: A05084C23AD458DE91CC4527FFB364F7' -H 'X-Deviceid: 7E526991B3044442859E154DA0218F9F' -H 'User-Agent: SdkEbankLoginExample/1.0.0 (iPhone; iOS 12.1; Scale/2.00)' -H 'X-Appver: 1.0.0.0' -H 'X-Platform: 27' -H 'X-Access-Token: 5lm4NBKCVeVxRg7YxLfw5UOgLty8GK1p97Q' -H 'Cookie: session_id=EF5232953F6E8EBCFB526F04EDABE574' --data-binary '{"entryId":12954303,"publicKeyId":"055327b90b4b87f4b8f0badf25ff1517","isCreditCard":false,"password":"","isSavePwd":true,"identification":{"isBreak":false,"pushToken":"","idfv":"C420D9E0-E18E-44D2-9CFA-B0E61C5F5D0E","devType":1,"macAddress":"","version":"12.1","routeMac":"","networkType":2,"idfa":"3E45B6A0-9496-4F58-A212-50F33F6096B3"}}' --compressed 'http://crawler.credit-manager.k2.wacaiyun.com/crawler/ebank/1.3/loginByEntry'
```

entryList
```json
{
	"status": {
		"responseCode": 200,
		"responseInfo": null
	},
	"entries": [{
		"entry": {
			"mId": 12953231,
			"nbkInputType": 1,
			"entryName": "6225768769098482",
			"importedStatus": 2,
			"importedTime": 1557027194,
			"importingTid": "bf281d2ac50e6ac30e32af3ce34fa22f",
			"errMsg": null,
			"hasPwd": true,
			"isDeleted": false,
			"nbkBank": {
				"bankId": 3,
				"bankName": "招商银行",
				"phone": "95555",
				"orderNo": 3,
				"isAvailable": true,
				"hasSmsCaptcha": false,
				"nbkInputTypeInfos": null,
				"nbkBankAccesses": [{
					"bankId": 3,
					"accessName": "CMBCC",
					"accessTitle": "信用卡",
					"isAvailable": true,
					"supportCreditCard": true,
					"openUrl": "http://bbs.creditcard.com.cn/thread-1093783-1-1.html?ccm_feedback=1&sfrom=sp",
					"forgetUrl": "https://mobile.cmbchina.com/MobileHtml/CreditCard/M_CustomerService/ResetPwd/RP_CheckCustomerInfo.aspx?ccm_feedback=1",
					"accessInputTypeInfos": [{
						"bankId": 3,
						"accessName": "CMBCC",
						"type": 4,
						"accTitle": "手机号",
						"pwdTitle": "登录密码",
						"accHint": "长度11位，银行预留手机号",
						"pwdHint": "一网通密码，6-8位字母、数字组合",
						"parseRule": "mobile",
						"isAvailable": true,
						"select": "",
						"pwdRegex": "^(?![0-9]+$)(?![a-zA-Z]+$)[0-9A-Za-z]{6,8}$",
						"pwdHint1": "登录密码",
						"pwdHint2": "6-8位字母、数字组合",
						"accHint1": "手机号",
						"accHint2": "长度11位",
						"accRegex": "^\\d{11}$"
					}, {
						"bankId": 3,
						"accessName": "CMBCC",
						"type": 2,
						"accTitle": "身份证",
						"pwdTitle": "密码",
						"accHint": "证件如有字母，请注意区分大小写",
						"pwdHint": "查询密码，6位数字",
						"parseRule": "identity",
						"isAvailable": true,
						"select": "",
						"pwdRegex": "",
						"pwdHint1": "查询密码",
						"pwdHint2": "6位数字",
						"accHint1": "身份证号码",
						"accHint2": "证件如有字母，请注意区分大小写",
						"accRegex": ""
					}, {
						"bankId": 3,
						"accessName": "CMBCC",
						"type": 1,
						"accTitle": "卡号",
						"pwdTitle": "密码",
						"accHint": "完整卡号",
						"pwdHint": "查询密码，6位数字",
						"parseRule": "card",
						"isAvailable": true,
						"select": "",
						"pwdRegex": "^\\d{6}$",
						"pwdHint1": "查询密码",
						"pwdHint2": "6位数字",
						"accHint1": "完整卡号",
						"accHint2": "须填写信用卡正面的完整卡号",
						"accRegex": "^\\d{15,19}$"
					}],
					"isNumVerifyCode": false
				}]
			},
			"realNameAuthentication": null,
			"access": "CMBCC"
		},
		"briefAccounts": null
	}, {
		"entry": {
			"mId": 12954303,
			"nbkInputType": 4,
			"entryName": "18636130636",
			"importedStatus": 2,
			"importedTime": 1556611913,
			"importingTid": "1d72d98bfe02b3c2a97a86eda98ab949",
			"errMsg": null,
			"hasPwd": true,
			"isDeleted": false,
			"nbkBank": {
				"bankId": 3,
				"bankName": "招商银行",
				"phone": "95555",
				"orderNo": 3,
				"isAvailable": true,
				"hasSmsCaptcha": false,
				"nbkInputTypeInfos": null,
				"nbkBankAccesses": [{
					"bankId": 3,
					"accessName": "CMBCC",
					"accessTitle": "信用卡",
					"isAvailable": true,
					"supportCreditCard": true,
					"openUrl": "http://bbs.creditcard.com.cn/thread-1093783-1-1.html?ccm_feedback=1&sfrom=sp",
					"forgetUrl": "https://mobile.cmbchina.com/MobileHtml/CreditCard/M_CustomerService/ResetPwd/RP_CheckCustomerInfo.aspx?ccm_feedback=1",
					"accessInputTypeInfos": [{
						"bankId": 3,
						"accessName": "CMBCC",
						"type": 4,
						"accTitle": "手机号",
						"pwdTitle": "登录密码",
						"accHint": "长度11位，银行预留手机号",
						"pwdHint": "一网通密码，6-8位字母、数字组合",
						"parseRule": "mobile",
						"isAvailable": true,
						"select": "",
						"pwdRegex": "^(?![0-9]+$)(?![a-zA-Z]+$)[0-9A-Za-z]{6,8}$",
						"pwdHint1": "登录密码",
						"pwdHint2": "6-8位字母、数字组合",
						"accHint1": "手机号",
						"accHint2": "长度11位",
						"accRegex": "^\\d{11}$"
					}, {
						"bankId": 3,
						"accessName": "CMBCC",
						"type": 2,
						"accTitle": "身份证",
						"pwdTitle": "密码",
						"accHint": "证件如有字母，请注意区分大小写",
						"pwdHint": "查询密码，6位数字",
						"parseRule": "identity",
						"isAvailable": true,
						"select": "",
						"pwdRegex": "",
						"pwdHint1": "查询密码",
						"pwdHint2": "6位数字",
						"accHint1": "身份证号码",
						"accHint2": "证件如有字母，请注意区分大小写",
						"accRegex": ""
					}, {
						"bankId": 3,
						"accessName": "CMBCC",
						"type": 1,
						"accTitle": "卡号",
						"pwdTitle": "密码",
						"accHint": "完整卡号",
						"pwdHint": "查询密码，6位数字",
						"parseRule": "card",
						"isAvailable": true,
						"select": "",
						"pwdRegex": "^\\d{6}$",
						"pwdHint1": "查询密码",
						"pwdHint2": "6位数字",
						"accHint1": "完整卡号",
						"accHint2": "须填写信用卡正面的完整卡号",
						"accRegex": "^\\d{15,19}$"
					}],
					"isNumVerifyCode": false
				}]
			},
			"realNameAuthentication": null,
			"access": "CMBCC"
		},
		"briefAccounts": null
	}, {
		"entry": {
			"mId": 12954720,
			"nbkInputType": 4,
			"entryName": "6225768769098482",
			"importedStatus": 2,
			"importedTime": 1536310170,
			"importingTid": "51e4d707eb7f1b52340cc217f1fd6a0d",
			"errMsg": null,
			"hasPwd": true,
			"isDeleted": false,
			"nbkBank": {
				"bankId": 3,
				"bankName": "招商银行",
				"phone": "95555",
				"orderNo": 3,
				"isAvailable": true,
				"hasSmsCaptcha": false,
				"nbkInputTypeInfos": null,
				"nbkBankAccesses": [{
					"bankId": 3,
					"accessName": "CMBCC",
					"accessTitle": "信用卡",
					"isAvailable": true,
					"supportCreditCard": true,
					"openUrl": "http://bbs.creditcard.com.cn/thread-1093783-1-1.html?ccm_feedback=1&sfrom=sp",
					"forgetUrl": "https://mobile.cmbchina.com/MobileHtml/CreditCard/M_CustomerService/ResetPwd/RP_CheckCustomerInfo.aspx?ccm_feedback=1",
					"accessInputTypeInfos": [{
						"bankId": 3,
						"accessName": "CMBCC",
						"type": 4,
						"accTitle": "手机号",
						"pwdTitle": "登录密码",
						"accHint": "长度11位，银行预留手机号",
						"pwdHint": "一网通密码，6-8位字母、数字组合",
						"parseRule": "mobile",
						"isAvailable": true,
						"select": "",
						"pwdRegex": "^(?![0-9]+$)(?![a-zA-Z]+$)[0-9A-Za-z]{6,8}$",
						"pwdHint1": "登录密码",
						"pwdHint2": "6-8位字母、数字组合",
						"accHint1": "手机号",
						"accHint2": "长度11位",
						"accRegex": "^\\d{11}$"
					}, {
						"bankId": 3,
						"accessName": "CMBCC",
						"type": 2,
						"accTitle": "身份证",
						"pwdTitle": "密码",
						"accHint": "证件如有字母，请注意区分大小写",
						"pwdHint": "查询密码，6位数字",
						"parseRule": "identity",
						"isAvailable": true,
						"select": "",
						"pwdRegex": "",
						"pwdHint1": "查询密码",
						"pwdHint2": "6位数字",
						"accHint1": "身份证号码",
						"accHint2": "证件如有字母，请注意区分大小写",
						"accRegex": ""
					}, {
						"bankId": 3,
						"accessName": "CMBCC",
						"type": 1,
						"accTitle": "卡号",
						"pwdTitle": "密码",
						"accHint": "完整卡号",
						"pwdHint": "查询密码，6位数字",
						"parseRule": "card",
						"isAvailable": true,
						"select": "",
						"pwdRegex": "^\\d{6}$",
						"pwdHint1": "查询密码",
						"pwdHint2": "6位数字",
						"accHint1": "完整卡号",
						"accHint2": "须填写信用卡正面的完整卡号",
						"accRegex": "^\\d{15,19}$"
					}],
					"isNumVerifyCode": false
				}]
			},
			"realNameAuthentication": null,
			"access": "CMBCC"
		},
		"briefAccounts": null
	}],
	"lastUptTime": 1557028149,
	"brokerDetailedEntries": null
}
```