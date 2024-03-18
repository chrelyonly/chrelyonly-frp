package util

var DivVersion = "3.1萝卜定制版"

// 全局参数方便调用
var GlobalDivConfigObject DivConfigObject

// 全局代理参数
var GlobalListMap ListMap

type DivVersionObject struct {
	DivVersion string `json:"version"`
	Msg        string `json:"msg"`
}

// div-config 配置文件
type DivConfigObject struct {
	// ServerName 服务器地址(默认上海阿里云服务器)
	ServerName string
	//验证token
	Token string
	//服务地址
	ServerAddr string
	//服务端口
	ServerPort int
}

type ListMap struct {
	ExeType                 string `json:"exe_type"`
	LocalIp                 string `json:"local_ip"`
	LocalPort               string `json:"local_port"`
	RemotePort              string `json:"remote_port"`
	CustomDomains           string `json:"custom_domains"`
	Comment                 string `json:"comment"`
	Type                    string `json:"type"`
	IsDom                   bool   `json:"false"`
	Plugin                  string `json:"plugin"`
	PluginLocalPath         string `json:"plugin_local_path"`
	PluginStripPrefix       string `json:"plugin_strip_prefix"`
	PluginHttpUser          string `json:"plugin_http_user"`
	PluginHttpPasswd        string `json:"plugin_http_passwd"`
	PluginLocalAddr         string `json:"plugin_local_addr"`
	PluginHostHeaderRewrite string `json:"plugin_host_header_rewrite"`
	PluginHeaderX           string `json:"plugin_header_X"`
	UseEncryption           bool   `json:"use_encryption"`
	UseCompression          bool   `json:"use_compression"`
}
