package util

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/fatedier/frp/pkg/util/version"
	"github.com/fatih/color"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

func MyVersion() error {
	color.Red("# 正在检查新版本更新...")
	response, err := http.Get("https://chrelyonly.cn/divFrpVersion.js")
	if err != nil {
		fmt.Println("请求错误：", err)
		return errors.New("请求错误")
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(response.Body)
	//读取数据
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("读取错误")
		return errors.New("读取错误")
	}
	var res DivVersionObject
	_ = json.Unmarshal(body, &res)
	contains := strings.Contains(res.DivVersion, DivVersion)
	color.Red("# div版本: v" + DivVersion)
	if contains {
		color.Green("# 当前是最新版本,无需更新...")
	} else {
		color.Red("# 发现新版本: v" + res.DivVersion)
	}
	color.Red("# 源版本: v" + version.Full())
	color.Red("# 版本消息: " + res.Msg)
	return nil
}
