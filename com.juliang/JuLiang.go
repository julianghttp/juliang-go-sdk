package main

import (
	"juliang_sdk/com.juliang/common"
	"juliang_sdk/com.juliang/enums"
	"juliang_sdk/com.juliang/ext"
	"net/url"
	"reflect"
)

//动态代理 -- 提取动态代理
func DynamicGetIps(dynamicgetips common.DynamicGetIps) string {
	t := reflect.TypeOf(dynamicgetips)
	v := reflect.ValueOf(dynamicgetips)
	params := make(map[string]string)
	var appKey string = dynamicgetips.Key()
	for k := 0; k < t.NumField(); k++ {
		key := t.Field(k).Name
		if key == "key" {
			continue
		}
		value := v.Field(k).String()
		if value == "" {
			continue
		}
		params[key] = value
	}
	var getipsParams = ext.GetParams(params, appKey)
	Url, err := url.Parse(enums.DYNAMIC_GETIPS)
	if err != nil {
		return ""
	}
	return ext.Post(Url, getipsParams)
}

//动态代理 -- 校验IP可用性
func DynamicCheck(check common.DynamicCheck) string {
	t := reflect.TypeOf(check)
	v := reflect.ValueOf(check)
	params := make(map[string]string)
	var appKey string = check.Key()
	for k := 0; k < t.NumField(); k++ {
		key := t.Field(k).Name
		if key == "key" {
			continue
		}
		value := v.Field(k).String()
		if value == "" {
			continue
		}
		params[key] = value
	}
	var getipsParams = ext.GetParams(params, appKey)
	Url, err := url.Parse(enums.DYNAMIC_CHECK)
	if err != nil {
		return ""
	}
	return ext.Post(Url, getipsParams)
}

//动态代理 -- 设置代理IP白名单
func DynamicSetWhiteIp(ip common.DynamicSetWhiteIp) string {
	t := reflect.TypeOf(ip)
	v := reflect.ValueOf(ip)
	params := make(map[string]string)
	var appKey string = ip.Key()
	for k := 0; k < t.NumField(); k++ {
		key := t.Field(k).Name
		if key == "key" {
			continue
		}
		value := v.Field(k).String()
		if value == "" {
			continue
		}
		params[key] = value
	}
	var getipsParams = ext.GetParams(params, appKey)
	Url, err := url.Parse(enums.DYNAMIC_SETWHITEIP)
	if err != nil {
		return ""
	}
	return ext.Post(Url, getipsParams)
}

//动态代理 -- 获取IP白名单
func DynamicGetWhiteIp(ip common.DynamicGetWhiteIp) string {
	t := reflect.TypeOf(ip)
	v := reflect.ValueOf(ip)
	params := make(map[string]string)
	var appKey string = ip.Key()
	for k := 0; k < t.NumField(); k++ {
		key := t.Field(k).Name
		if key == "key" {
			continue
		}
		value := v.Field(k).String()
		if value == "" {
			continue
		}
		params[key] = value
	}
	var getipsParams = ext.GetParams(params, appKey)
	Url, err := url.Parse(enums.DYNAMIC_GETWHITEIP)
	if err != nil {
		return ""
	}
	return ext.Post(Url, getipsParams)
}

//动态代理 -- 获取代理剩余可用时长
func DynamicRemain(remain common.DynamicRemain) string {
	t := reflect.TypeOf(remain)
	v := reflect.ValueOf(remain)
	params := make(map[string]string)
	var appKey string = remain.Key()
	for k := 0; k < t.NumField(); k++ {
		key := t.Field(k).Name
		if key == "key" {
			continue
		}
		value := v.Field(k).String()
		if value == "" {
			continue
		}
		params[key] = value
	}
	var getipsParams = ext.GetParams(params, appKey)
	Url, err := url.Parse(enums.DYNAMIC_REMAIN)
	if err != nil {
		return ""
	}
	return ext.Post(Url, getipsParams)
}

//动态代理 -- 获取剩余可用时长
func DynamicBalance(balance common.DynamicBalance) string {
	t := reflect.TypeOf(balance)
	v := reflect.ValueOf(balance)
	params := make(map[string]string)
	var appKey string = balance.Key()
	for k := 0; k < t.NumField(); k++ {
		key := t.Field(k).Name
		if key == "key" {
			continue
		}
		value := v.Field(k).String()
		if value == "" {
			continue
		}
		params[key] = value
	}
	var getipsParams = ext.GetParams(params, appKey)
	Url, err := url.Parse(enums.DYNAMIC_BALANCE)
	if err != nil {
		return ""
	}
	return ext.Post(Url, getipsParams)
}

//获取账户余额
func UsersGetBalance(balance common.UsersGetBalance) string {
	t := reflect.TypeOf(balance)
	v := reflect.ValueOf(balance)
	params := make(map[string]string)
	var appKey string = balance.Key()
	for k := 0; k < t.NumField(); k++ {
		key := t.Field(k).Name
		if key == "key" {
			continue
		}
		value := v.Field(k).String()
		if value == "" {
			continue
		}
		params[key] = value
	}
	var getipsParams = ext.GetParams(params, appKey)
	Url, err := url.Parse(enums.USERS_GETBALANCE)
	if err != nil {
		return ""
	}
	return ext.Post(Url, getipsParams)
}

//独享代理 -- 获取代理详情
func AloneGetIps(ips common.AloneGetIps) string {
	t := reflect.TypeOf(ips)
	v := reflect.ValueOf(ips)
	params := make(map[string]string)
	var appKey string = ips.Key()
	for k := 0; k < t.NumField(); k++ {
		key := t.Field(k).Name
		if key == "key" {
			continue
		}
		value := v.Field(k).String()
		if value == "" {
			continue
		}
		params[key] = value
	}
	var getipsParams = ext.GetParams(params, appKey)
	Url, err := url.Parse(enums.ALONE_GETIPS)
	if err != nil {
		return ""
	}
	return ext.Post(Url, getipsParams)
}

//独享代理 -- 设置代理IP白名单
func AloneSetWhiteIp(ip common.AloneSetWhiteIp) string {
	t := reflect.TypeOf(ip)
	v := reflect.ValueOf(ip)
	params := make(map[string]string)
	var appKey string = ip.Key()
	for k := 0; k < t.NumField(); k++ {
		key := t.Field(k).Name
		if key == "key" {
			continue
		}
		value := v.Field(k).String()
		if value == "" {
			continue
		}
		params[key] = value
	}
	var getipsParams = ext.GetParams(params, appKey)
	Url, err := url.Parse(enums.ALONE_SETWHITEIP)
	if err != nil {
		return ""
	}
	return ext.Post(Url, getipsParams)
}

//独享代理 -- 获取代理IP白名单
func AloneGetWhiteIp(ip common.AloneGetWhiteIp) string {
	t := reflect.TypeOf(ip)
	v := reflect.ValueOf(ip)
	params := make(map[string]string)
	var appKey string = ip.Key()
	for k := 0; k < t.NumField(); k++ {
		key := t.Field(k).Name
		if key == "key" {
			continue
		}
		value := v.Field(k).String()
		if value == "" {
			continue
		}
		params[key] = value
	}
	var getipsParams = ext.GetParams(params, appKey)
	Url, err := url.Parse(enums.ALONE_GETWHITEIP)
	if err != nil {
		return ""
	}
	return ext.Post(Url, getipsParams)
}
