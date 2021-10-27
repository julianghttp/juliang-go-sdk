package ext

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"reflect"
	"sort"
	"strings"
)

//结构体转换为map
func StructToMap(t reflect.Type, v reflect.Value, params map[string]string) map[string]string {
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
	return params
}

//计算签名
func GetParams(object map[string]string, appKey string) map[string]string {
	cleanMap := cleanParams(object)
	var value string = joinStringsInASCII(cleanMap, "&", false, false)
	value = value + "&key=" + appKey
	//fmt.Println(value)
	h := md5.New()
	h.Write([]byte(value))
	var sign string = hex.EncodeToString(h.Sum(nil))
	//fmt.Println(sign)
	cleanMap["sign"] = sign
	return cleanMap
}

//格式化参数，去掉为空的参数
func cleanParams(kv map[string]string) map[string]string {
	params := make(map[string]string)
	for k, v := range kv {
		if k == "" {
			continue
		}
		params[k] = v
	}
	return params
}

//JoinStringsInASCII 按照规则，参数名ASCII码从小到大排序后拼接
//data 待拼接的数据
//sep 连接符
//onlyValues 是否只包含参数值，true则不包含参数名，否则参数名和参数值均有
//includeEmpty 是否包含空值，true则包含空值，否则不包含，注意此参数不影响参数名的存在
//exceptKeys 被排除的参数名，不参与排序及拼接
func joinStringsInASCII(data map[string]string, sep string, onlyValues, includeEmpty bool, exceptKeys ...string) string {
	var list []string
	var keyList []string
	m := make(map[string]int)
	if len(exceptKeys) > 0 {
		for _, except := range exceptKeys {
			m[except] = 1
		}
	}
	for k := range data {
		if _, ok := m[k]; ok {
			continue
		}
		value := data[k]
		if !includeEmpty && value == "" {
			continue
		}
		if onlyValues {
			keyList = append(keyList, k)
		} else {
			list = append(list, fmt.Sprintf("%s=%s", k, value))
		}
	}
	if onlyValues {
		sort.Strings(keyList)
		for _, v := range keyList {
			list = append(list, data[v])
		}
	} else {
		sort.Strings(list)
	}
	return strings.Join(list, sep)
}
