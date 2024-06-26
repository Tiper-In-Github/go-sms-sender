# go-sms-sender --For Quanming Network

## This is a special version that is independently maintained
- If you need to use it, please pay attention to compatibility (as of now, it is compatible with the casdoor latest version from April 2024)
- 如需使用请注意兼容性(截至目前，兼容2024年4月的最新版casdoor)
- In order to use this independent support, you need to modify two parts of the `casdoor/web/src/Setting.js` configuration file (namely, the SMS section under`OtherProviderInfo`and`getProviderTypeOptions`)There are only two areas that need to be modified.
- 为了使用该独立支持，你需要修改`casdoor/web/src/Setting.js`配置文件中的两处(分别是：`OtherProviderInfo`和`getProviderTypeOptions`下的SMS部分)
需要修改的只有两处，示例片段:
```js
export const OtherProviderInfo = {
  SMS: {
    "Aliyun SMS": {
      logo: `${StaticBaseUrl}/img/social_aliyun.png`,
      url: "https://aliyun.com/product/sms",
    },
    "Amazon SNS": {
      logo: `${StaticBaseUrl}/img/social_aws.png`,
      url: "https://aws.amazon.com/cn/sns/",
    },
    "Azure ACS": {
      logo: `${StaticBaseUrl}/img/social_azure.png`,
      url: "https://azure.microsoft.com/en-us/products/communication-services",
    },
    "Infobip SMS": {
      logo: `${StaticBaseUrl}/img/social_infobip.png`,
      url: "https://portal.infobip.com/homepage/",
    },
    "Tencent Cloud SMS": {
      logo: `${StaticBaseUrl}/img/social_tencent_cloud.jpg`,
      url: "https://cloud.tencent.com/product/sms",
    },
    "Baidu Cloud SMS": {
      logo: `${StaticBaseUrl}/img/social_baidu_cloud.png`,
      url: "https://cloud.baidu.com/product/sms.html",
    },
    "Volc Engine SMS": {
      logo: `${StaticBaseUrl}/img/social_volc_engine.jpg`,
      url: "https://www.volcengine.com/products/cloud-sms",
    },
    "Huawei Cloud SMS": {
      logo: `${StaticBaseUrl}/img/social_huawei.png`,
      url: "https://www.huaweicloud.com/product/msgsms.html",
    },
    "UCloud SMS": {
      logo: `${StaticBaseUrl}/img/social_ucloud.png`,
      url: "https://www.ucloud.cn/site/product/usms.html",
    },
    "Twilio SMS": {
      logo: `${StaticBaseUrl}/img/social_twilio.svg`,
      url: "https://www.twilio.com/messaging",
    },
    "SmsBao SMS": {
      logo: `${StaticBaseUrl}/img/social_smsbao.png`,
      url: "https://www.smsbao.com/",
    },
    "SUBMAIL SMS": {
      logo: `${StaticBaseUrl}/img/social_submail.svg`,
      url: "https://www.mysubmail.com",
    },
    "Msg91 SMS": {
      logo: `${StaticBaseUrl}/img/social_msg91.ico`,
      url: "https://control.msg91.com/app/",
    },
    "OSON SMS": {
      logo: "https://osonsms.com/images/osonsms-logo.svg",
      url: "https://osonsms.com/",
    },
    "Custom HTTP SMS": {
      logo: `${StaticBaseUrl}/img/social_default.png`,
      url: "https://casdoor.org/docs/provider/sms/overview",
    },
    "Mock SMS": {
      logo: `${StaticBaseUrl}/img/social_default.png`,
      url: "",
    },
    "Quanm SMS": {  // 添加这段  Add this
      logo: "https://static.mp.quanmwl.com/static/images/favicon_dev.ico",
      url: "https://dev.quanmwl.com/console",
    },
  },
// ......
// 不用改的地方省略
// ......
else if (category === "SMS") {
    return (
      [
        {id: "Aliyun SMS", name: "Alibaba Cloud SMS"},
        {id: "Amazon SNS", name: "Amazon SNS"},
        {id: "Azure ACS", name: "Azure ACS"},
        {id: "Custom HTTP SMS", name: "Custom HTTP SMS"},
        {id: "Mock SMS", name: "Mock SMS"},
        {id: "OSON SMS", name: "OSON SMS"},
        {id: "Infobip SMS", name: "Infobip SMS"},
        {id: "Tencent Cloud SMS", name: "Tencent Cloud SMS"},
        {id: "Baidu Cloud SMS", name: "Baidu Cloud SMS"},
        {id: "Volc Engine SMS", name: "Volc Engine SMS"},
        {id: "Huawei Cloud SMS", name: "Huawei Cloud SMS"},
        {id: "UCloud SMS", name: "UCloud SMS"},
        {id: "Twilio SMS", name: "Twilio SMS"},
        {id: "SmsBao SMS", name: "SmsBao SMS"},
        {id: "SUBMAIL SMS", name: "SUBMAIL SMS"},
        {id: "Msg91 SMS", name: "Msg91 SMS"},
        {id: "Quanm SMS", name: "Quanm SMS"},  // 添加这段  Add this
      ]
    );
  }
// ......
// ......
```

[![Go Report Card](https://goreportcard.com/badge/github.com/casdoor/go-sms-sender)](https://goreportcard.com/report/github.com/casdoor/go-sms-sender)
[![Go](https://github.com/casdoor/go-sms-sender/actions/workflows/ci.yml/badge.svg)](https://github.com/casdoor/go-sms-sender/actions/workflows/ci.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/casdoor/go-sms-sender.svg)](https://pkg.go.dev/github.com/casdoor/go-sms-sender)
![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/casdoor/go-sms-sender)

This is a powerful open-source library for sending SMS message, which will help you to easily integrate with the popular SMS providers. And it has been applied to [Casdoor](https://github.com/casdoor/casdoor), if you still don’t know how to use it after reading README.md, you can refer to it.

We support the following SMS providers, welcome to contribute.

- [Twilio](https://www.twilio.com)
- [Amazon SNS](https://aws.amazon.com/sns/)
- [Azure ACS](https://azure.microsoft.com/en-us/products/communication-services)
- [GCCPAY](https://gccpay.com/)
- [Infobip](https://www.infobip.com/)
- [SUBMAIL](https://en.mysubmail.com/)
- [SmsBao](https://www.smsbao.com/)
- [Alibaba Cloud](https://www.aliyun.com/product/sms)
- [Tencent Cloud](https://cloud.tencent.com/document/product/382)
- [Baidu Cloud](https://cloud.baidu.com/product/sms.html)
- [VolcEngine](https://www.volcengine.com/product/cloud-sms)
- [Huawei Cloud](https://www.huaweicloud.com/product/msgsms.html)
- [UCloud](https://www.ucloud.cn/site/product/usms.html)
- [Huyi](https://www.ihuyi.com/)
- [Netgsm](https://www.netgsm.com.tr/)
- [Oson Sms](https://osonsms.com/)
- [Quanming Sms](https://dev.quanmwl.com/)

## Installation

Use `go get` to install：

```
go get github.com/casdoor/go-sms-sender
```

## How to use

### Create Client

Different SMS providers need to provide different configuration, but we support a unit API as below to init and get the SMS client.

```go
func NewSmsClient(provider string, accessId string, accessKey string, sign string, template string, other ...string) (SmsClient, error)
```

- `provider` the name of SMS provider, such as `Aliyun SMS`
- `accessId`
- `accessKey`
- `sign` the sign name
- `template` the template code
- `other` other configuration

### Send Message

After initializing the SMS client, we can use the following API to send message.

```go
SendMessage(param map[string]string, targetPhoneNumber ...string) error
```

- `param` the parameters in the SMS template, such as 6 random numbers
- `targetPhoneNumber` the receivers, such as `+8612345678910`

## Example

### Twilio

Please get necessary information from Twilio [console](https://console.twilio.com/)

```go
package main

import "github.com/casdoor/go-sms-sender"

func main() {
	client, err := go_sms_sender.NewSmsClient(go_sms_sender.Twilio, "ACCOUNT_SID", "AUTH_TOKEN", "", "TEMPLATE_CODE")
	if err != nil {
		panic(err)
	}

	params := map[string]string{}
	params["code"] = "123456"
	phoneNumer := "+8612345678910"
	err = client.SendMessage(params, phoneNumer)
	if err != nil {
		panic(err)
	}
}
```

### Aliyun

Before you begin, you need to sign up for an Aliyun account and retrieve your [Credentials](https://usercenter.console.aliyun.com/#/manage/ak).

```go
package main

import "github.com/casdoor/go-sms-sender"

func main() {
	client, err := go_sms_sender.NewSmsClient(go_sms_sender.Aliyun, "ACCESS_KEY_ID", "ACCESS_KEY_SECRET", "SIGN_NAME", "TEMPLATE_CODE")
	if err != nil {
		panic(err)
	}

	params := map[string]string{}
	params["code"] = "473956"
	phoneNumer := "+8612345678910"
	err = client.SendMessage(params, phoneNumer)
	if err != nil {
		panic(err)
	}
}
```

### Tencent Cloud

```go
package main

import "github.com/casdoor/go-sms-sender"

func main() {
	client, err := go_sms_sender.NewSmsClient(go_sms_sender.TencentCloud, "secretId", "secretKey", "SIGN_NAME", "TEMPLATE_CODE", "APP_ID")
	if err != nil {
		panic(err)
	}

	params := map[string]string{}
	params["0"] = "473956"
	phoneNumer := "+8612345678910"
	err = client.SendMessage(params, phoneNumer)
	if err != nil {
		panic(err)
	}
}
```

### Netgsm

- yourAccessId: is KullaniciAdi
- yourAccessKey: is Sifre
- yourSign: is Baslik

```go
package main

import "github.com/casdoor/go-sms-sender"

func main() {
	client, err := go_sms_sender.NewSmsClient(go_sms_sender.Netgsm, "yourAccessId", "yourAccessKey", "yourSign", "yourTemplate")
	if err != nil {
		panic(err)
	}

	params := map[string]string{}
	params["param1"] = "value1"
	params["param2"] = "value2"
	phoneNumer := "+8612345678910"
	err = client.SendMessage(params, phoneNumer)
	if err != nil {
		panic(err)
	}
}
```

### Oson Sms

- senderId: is `login`
- secretAccessKey: is `hash`
- signName: is `from`
- templateCode: is `message`

```go
package main

func main() {
	client, err := go_sms_sender.NewSmsClient(go_sms_senderOsonSms, "senderId", "secretAccessKey", "signName", "templateCode")
	if err != nil {
		panic(err)
	}

	params := map[string]string{}
	params["code"] = "123456"
	phoneNumer := "+992123456789"
	err = client.SendMessage(params, phoneNumer)
	if err != nil {
		panic(err)
	}
}
```


### Running Tests

To run tests for the `go-sms-sender` library, navigate to the root folder of the project in your terminal and execute the following command:

```sh
go test -v ./...
```

you can modify mock_test.go file to mock an other tests

## License

This project is licensed under the [Apache 2.0 license](LICENSE).
