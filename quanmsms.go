// Copyright 2022 The Casdoor Authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package go_sms_sender

import (
	"fmt"
	"strconv"
	"strings"

	"gitee.com/chengdu-quanming-network/quanmsms-go"
)

// console:  https://dev.quanmwl.com/console
// The use of SDK is to ensure the use of the latest standards

type QuanmSMSClient struct {
	openid     string
	apikey     string
	templateid string
}

var areaCodes = []string{
	"+1",   // USA, Canada
	"+7",   // Russia, Kazakhstan
	"+20",  // Egypt
	"+27",  // South Africa
	"+30",  // Greece
	"+31",  // Netherlands
	"+32",  // Belgium
	"+33",  // France
	"+34",  // Spain
	"+36",  // Hungary
	"+39",  // Italy
	"+40",  // Romania
	"+41",  // Switzerland
	"+43",  // Austria
	"+44",  // United Kingdom
	"+45",  // Denmark
	"+46",  // Sweden
	"+47",  // Norway
	"+48",  // Poland
	"+49",  // Germany
	"+51",  // Peru
	"+52",  // Mexico
	"+53",  // Cuba
	"+54",  // Argentina
	"+55",  // Brazil
	"+56",  // Chile
	"+57",  // Colombia
	"+58",  // Venezuela
	"+60",  // Malaysia
	"+61",  // Australia
	"+62",  // Indonesia
	"+63",  // Philippines
	"+64",  // New Zealand
	"+65",  // Singapore
	"+66",  // Thailand
	"+81",  // Japan
	"+82",  // South Korea
	"+84",  // Vietnam
	"+86",  // China
	"+90",  // Turkey
	"+91",  // India
	"+92",  // Pakistan
	"+93",  // Afghanistan
	"+94",  // Sri Lanka
	"+95",  // Myanmar
	"+98",  // Iran
	"+211", // South Sudan
	"+212", // Morocco
	"+213", // Algeria
	"+216", // Tunisia
	"+218", // Libya
	"+220", // Gambia
	"+221", // Senegal
	"+222", // Mauritania
	"+223", // Mali
	"+224", // Guinea
	"+225", // Ivory Coast
	"+226", // Burkina Faso
	"+227", // Niger
	"+228", // Togo
	"+229", // Benin
	"+230", // Mauritius
	"+231", // Liberia
	"+232", // Sierra Leone
	"+233", // Ghana
	"+234", // Nigeria
	"+235", // Chad
	"+236", // Central African Republic
	"+237", // Cameroon
	"+238", // Cape Verde
	"+239", // Sao Tome and Principe
	"+240", // Equatorial Guinea
	"+241", // Gabon
	"+242", // Republic of the Congo
	"+243", // Democratic Republic of the Congo
	"+244", // Angola
	"+245", // Guinea-Bissau
	"+246", // British Indian Ocean Territory
	"+247", // Ascension Island
	"+248", // Seychelles
	"+249", // Sudan
	"+250", // Rwanda
	"+251", // Ethiopia
	"+252", // Somalia
	"+253", // Djibouti
	"+254", // Kenya
	"+255", // Tanzania
	"+256", // Uganda
	"+257", // Burundi
	"+258", // Mozambique
	"+260", // Zambia
	"+261", // Madagascar
	"+262", // Reunion
	"+263", // Zimbabwe
	"+264", // Namibia
	"+265", // Malawi
	"+266", // Lesotho
	"+267", // Botswana
	"+268", // Swaziland
	"+269", // Comoros
	"+290", // Saint Helena
	"+291", // Eritrea
	"+297", // Aruba
	"+298", // Faroe Islands
	"+299", // Greenland
	"+350", // Gibraltar
	"+351", // Portugal
	"+352", // Luxembourg
	"+353", // Ireland
	"+354", // Iceland
	"+355", // Albania
	"+356", // Malta
	"+357", // Cyprus
	"+358", // Finland
	"+359", // Bulgaria
	"+370", // Lithuania
	"+371", // Latvia
	"+372", // Estonia
	"+373", // Moldova
	"+374", // Armenia
	"+375", // Belarus
	"+376", // Andorra
	"+377", // Monaco
	"+378", // San Marino
	"+379", // Vatican
	"+380", // Ukraine
	"+381", // Serbia
	"+382", // Montenegro
	"+383", // Kosovo
	"+385", // Croatia
	"+386", // Slovenia
	"+387", // Bosnia and Herzegovina
	"+389", // Macedonia
	"+420", // Czech Republic
	"+421", // Slovakia
	"+423", // Liechtenstein
	"+500", // Falkland Islands
	"+501", // Belize
	"+502", // Guatemala
	"+503", // El Salvador
	"+504", // Honduras
	"+505", // Nicaragua
	"+506", // Costa Rica
	"+507", // Panama
	"+508", // Saint Pierre and Miquelon
	"+509", // Haiti
	"+590", // Guadeloupe
	"+591", // Bolivia
	"+592", // Guyana
	"+593", // Ecuador
	"+594", // French Guiana
	"+595", // Paraguay
	"+596", // Martinique
	"+597", // Suriname
	"+598", // Uruguay
	"+599", // Netherlands Antilles
	"+670", // East Timor
	"+672", // Australian External Territories
	"+673", // Brunei
	"+674", // Nauru
	"+675", // Papua New Guinea
	"+676", // Tonga
	"+677", // Solomon Islands
	"+678", // Vanuatu
	"+679", // Fiji
	"+680", // Palau
	"+681", // Wallis and Futuna
	"+682", // Cook Islands
	"+683", // Niue
	"+685", // Samoa
	"+686", // Kiribati
	"+687", // New Caledonia
	"+688", // Tuvalu
	"+689", // French Polynesia
	"+690", // Tokelau
	"+691", // Micronesia
	"+692", // Marshall Islands
	"+850", // North Korea
	"+852", // Hong Kong
	"+853", // Macao
	"+855", // Cambodia
	"+856", // Laos
	"+880", // Bangladesh
	"+886", // Taiwan
	"+960", // Maldives
	"+961", // Lebanon
	"+962", // Jordan
	"+963", // Syria
	"+964", // Iraq
	"+965", // Kuwait
	"+966", // Saudi Arabia
	"+967", // Yemen
	"+968", // Oman
	"+970", // Palestine
	"+971", // United Arab Emirates
	"+972", // Israel
	"+973", // Bahrain
	"+974", // Qatar
	"+975", // Bhutan
	"+976", // Mongolia
	"+977", // Nepal
	"+992", // Tajikistan
	"+993", // Turkmenistan
	"+994", // Azerbaijan
	"+995", // Georgia
	"+996", // Kyrgyzstan
	"+998", // Uzbekistan
}

func GetQuanmSMSClient(openid string, apikey string, templateid string) (*QuanmSMSClient, error) {
	return &QuanmSMSClient{
		openid:     openid,
		apikey:     apikey,
		templateid: templateid,
	}, nil
}

func (c *QuanmSMSClient) SendMessage(param map[string]string, targetPhoneNumber ...string) error {
	code, ok := param["code"]
	if !ok {
		return fmt.Errorf("missing parameter: code")
	}

	if len(targetPhoneNumber) == 0 {
		return fmt.Errorf("missing parameter: targetPhoneNumber")
	}

	smsSDK := quanmsms.NewSmsSDK(c.openid, c.apikey, "https")
	template_args := map[string]interface{}{"code": code}
	for _, tel := range targetPhoneNumber {
		templateidInt64, intErr := strconv.Atoi(c.templateid)
		if intErr != nil {
			return fmt.Errorf("template id error")
		}
		// casdoos会给手机号加上区域代码，需要去除
		var phone string = tel
		for _, code := range areaCodes {
			if strings.HasPrefix(tel, code) {
				phone = strings.TrimPrefix(tel, code)
				break
			}
		}

		resp, err := smsSDK.Send(phone, int8(templateidInt64), template_args)
		if err != nil {
			return fmt.Errorf("other error:" + err.Error())
		}
		switch string(resp.State) {
		case "201":
			// 表单信息有误或触发限发机制
			return fmt.Errorf("refusal to send due to security reasons")
		case "202":
			// 账户重复
			return fmt.Errorf("account repeat")
		case "203":
			// 服务器异常或限流
			return fmt.Errorf("server error,Please try again later")
		case "205":
			// 请求不安全
			return fmt.Errorf("illegal request")
		case "207":
			// 配额不足
			return fmt.Errorf("insufficient balance")
		case "208":
			// 验签失败
			return fmt.Errorf("verification failed")
		case "209":
			// 账户被禁用接口
			return fmt.Errorf("insufficient permissions")
		case "210":
			// 账户被冻结
			return fmt.Errorf("account frozen")
		case "211":
			// 请求参数超过上限
			return fmt.Errorf("parameter too long")
		case "212":
			// 权限不足或使用了他人模板
			return fmt.Errorf("insufficient permissions or using someone else's template")
		case "213":
			// 调用状态异常
			return fmt.Errorf("status error")
		case "215":
			// 内容受限
			return fmt.Errorf("content restricted")
		case "216":
			// 内容违法
			return fmt.Errorf("content violation")
		}
	}

	return nil
}
